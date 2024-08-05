package types

type User struct {
	Username string
	Password string
	Role     string
}

type Role struct {
	Name        string
	Permissions []string
}
