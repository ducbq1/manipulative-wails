package services

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/dustin/go-humanize"
)

// WriteCounter stores dl state of the cl binaries
type WriteCounter struct {
	Total    uint64
	LastEmit uint64
	Filename string
	a        *homeService
}

func (a *homeService) javaInstalled() bool {
	var javaInstalled bool
	if a.paths.Java[len(a.paths.Java)-9:] != "javaw.exe" {
		javaInstalled = false
	} else {
		javaInstalled = true
	}
	return javaInstalled
}

// normalizeAmounts takes amount/fee in int64 and normalizes it. Example: passing 821500000000 will return 8215
func normalizeAmounts(i int64) (string, error) {
	f := fmt.Sprintf("%.8f", float64(i)/1e8)
	return f, nil
}

// TempFileName creates temporary file names for the transaction files
func (a *homeService) TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(a.paths.TMPDir, prefix+hex.EncodeToString(randBytes)+suffix)
}

// Write emits the download progress of the CL binaries to the frontend
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)

	if (wc.Total - wc.LastEmit) > uint64(800) {
		wc.a.RT.Events.Emit("downloading", wc.Filename, humanize.Bytes(wc.Total))
		wc.LastEmit = wc.Total
	}

	return n, nil
}

func (a *homeService) directoryCreator(directories ...string) error {
	for _, d := range directories {
		err := os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *homeService) fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data []byte) error {

	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (a *homeService) newReleaseAvailable() {
	update := new(UpdateWallet)
	update.app = a
	currentRelease := a.Version

	a.log.Infoln("Checking for new releases...")

	go func() {
		for i := 200; i > 0; i-- {
			newRelease := update.GetLatestRelease()
			if currentRelease != newRelease {
				a.log.Infoln("There's a newer release available")
				a.RT.Events.Emit("new_release", newRelease)
			}
			time.Sleep(time.Duration(i) * time.Second)
		}
	}()

}

type UpdateWallet struct {
	currentVersion  string
	newVersion      string
	mollyBinaryPath string
	dagFolderPath   string
	app             *homeService
}

func (u *UpdateWallet) Run() {
	var err error

	u.currentVersion = u.app.Version
	u.newVersion = u.GetLatestRelease()
	u.mollyBinaryPath, err = os.Executable()
	if err != nil {
		u.app.log.Errorln("Unable to collect the path of the molly wallet binary. Reason: ", err)
	}
	u.dagFolderPath = u.app.paths.DAGDir

	err = u.TriggerUpdate()
	if err != nil {
		u.app.log.Errorln("Unable to Update Molly Wallet. Reason: ", err)
	}

}

func (u *UpdateWallet) TriggerUpdate() error {

	main := u.dagFolderPath + "/update" + ".exe"
	args := []string{"-init_dag_path=" + u.dagFolderPath, "-init_molly_path=" + u.mollyBinaryPath, "-new_version=" + u.newVersion, "-upgrade=" + "true"}

	cmd := exec.Command(main, args...)
	u.app.log.Infoln("Running command: ", cmd)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr // Captures STDERR

	err := cmd.Run()
	if err != nil {
		errFormatted := fmt.Sprint(err) + ": " + stderr.String()
		return errors.New(errFormatted)
	}
	return nil
}

func (u *UpdateWallet) GetLatestRelease() string {

	const (
		url = "https://api.github.com/repos/grvlle/constellation_wallet/releases/latest"
	)

	resp, err := http.Get(url)
	if err != nil {
		u.app.log.Errorln("Failed to send HTTP request. Reason: ", err)
		return ""
	}
	if resp == nil {
		u.app.log.Errorln("Killing pollTokenBalance after 10 failed attempts to get balance from mainnet, Reason: ", err)
		return ""
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		u.app.log.Warn("Unable to update token balance. Reason: ", err)
		return ""
	}

	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return ""
	}

	release := result["tag_name"]
	bytes := []byte(release.(string))
	version := string(bytes[1:6])
	return version

}
