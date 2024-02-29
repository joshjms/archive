package user

type User struct {
	ID       int
	Username string
	Password string
}

type Signin struct {
	Username string
	Password string
}

func (u *User) Create() (err error) {
	return
}
