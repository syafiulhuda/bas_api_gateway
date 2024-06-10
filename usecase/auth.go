package usecase

type Login struct {
}

type LoginInterface interface {
	Authenticated(username, password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (l *Login) Authenticated(username, password string) bool {
	if username == "admin" && password == "admin123" {
		return true
	}
	return false

}
