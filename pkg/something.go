//go:generate mockgen -source=something.go -destination=mocks.go -package=example

package example

type User struct {
	ID   string
	Name string
}

type Database interface {
	GetUser(id string) (User, error)
}
