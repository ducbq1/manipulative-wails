package services

import (
	"context"
	"io"
	"os"
	"os/user"
	"sync"
	"tinyrdm/backend/domain"
	"tinyrdm/backend/types"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type homeService struct {
	ctx        context.Context
	ctxCancel  context.CancelFunc
	mutex      sync.Mutex
	db         *gorm.DB
	RT         *wails.Runtime
	log        *logrus.Logger
	wallet     types.Session
	Version    string
	killSignal chan struct{}
	paths      struct {
		HomeDir        string
		DAGDir         string
		TMPDir         string
		EncryptedDir   string
		EncPrivKeyFile string
		EmptyTXFile    string
		AddressFile    string
		ImageDir       string
		Java           string
	}
	KeyStoreAccess      bool
	TransactionFinished bool
	TransactionFailed   bool
	UserLoggedIn        bool
	NewUser             bool
	WalletImported      bool
	WalletCLI           struct {
		URL     string
		Version string
	}
}

var home *homeService
var onceHome sync.Once

func Home() *homeService {
	if home == nil {
		onceHome.Do(func() {
			home = &homeService{}
		})
	}
	return home
}

func (c *homeService) Start(ctx context.Context) {
	c.ctx, c.ctxCancel = context.WithCancel(ctx)
}

func (a *homeService) ShutDown() {
	a.wallet = types.Session{}
	close(a.killSignal) // Kills the Go Routines
}

func (a *homeService) init(runtime *wails.Runtime) error {
	var err error

	a.log = logrus.New()
	err = a.initDirectoryStructure()
	if err != nil {
		a.log.Errorln("Unable to set up directory structure. Reason: ", err)
	}

	a.initLogger()

	a.UserLoggedIn = false
	a.NewUser = false
	a.TransactionFinished = true
	a.RT = runtime
	a.killSignal = make(chan struct{}) // Used to kill go routines and hand back system resources
	a.WalletCLI.URL = "https://github.com/Constellation-Labs/constellation/releases/download"
	a.WalletCLI.Version = "2.6.0"
	a.db, err = gorm.Open(sqlite.Open(a.paths.DAGDir+"/store.db"), &gorm.Config{})
	if err != nil {
		a.log.Panicln("failed to connect database", err)
	}
	// Migrate the schema

	a.db.AutoMigrate(&domain.Wallet{})
	a.newReleaseAvailable()

	return nil
}

// initLogger writes logs to STDOUT and a.paths.DAGDir/wallet.log
func (a *homeService) initLogger() {
	logFile, err := os.OpenFile(a.paths.DAGDir+"/wallet.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		a.log.Fatal("Unable to create log file.")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	a.log.SetOutput(mw)
	a.log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

// Initializes the Directory Structure and stores the paths to the WalletApplication struct.
func (a *homeService) initDirectoryStructure() error {

	user, err := user.Current()
	if err != nil {
		return err
	}

	a.paths.HomeDir = user.HomeDir             // Home directory of the user
	a.paths.DAGDir = a.paths.HomeDir + "/.dag" // DAG directory for configuration files and wallet specific data
	a.paths.TMPDir = a.paths.DAGDir + "/tmp"
	a.paths.EncPrivKeyFile = a.paths.EncryptedDir
	a.paths.EmptyTXFile = a.paths.TMPDir + "/genesis_tx"
	a.paths.ImageDir = "./frontend/src/assets/img/" // Image Folder

	a.log.Info("DAG Directory: ", a.paths.DAGDir)

	err = a.directoryCreator(a.paths.DAGDir, a.paths.TMPDir)
	if err != nil {
		return err
	}

	return nil
}

func (c *homeService) StartDatabase() (resp types.JSResp) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	c.db = db
	resp.Success = true
	return
}
