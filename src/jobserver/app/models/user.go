package models

type User struct {
	Id        int
	UserName  string
	Password  string
	Salt      string
	Email     string
	LastLogin int64
	LastIp    string
	Status    int
}

func (u *User) Update(fields ...string) error {
	return nil
}

func UserAdd(user *User) (int64, error) {
	return 0, nil
}

func UserGetById(id int) (*User, error) {
	return nil, nil
}

func UserGetByName(userName string) (*User, error) {
	return nil, nil
}

func UserUpdate(user *User, fields ...string) error {
	return nil
}
