package users

type Role int

const (
	Student Role = iota
	Teacher
)

type User struct {
	ID    int64
	Name  string
	Email string
	Role  Role
}
