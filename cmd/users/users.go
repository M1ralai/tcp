package users

import "errors"

type Users struct {
	Username   string
	Password   string
	isLoggedIn bool
}

var users []Users

func RegisterUser(username string, password string) (*Users, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password cannot be blank")
	}
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			return nil, errors.New("there is already a user with this username")
		}
	}
	users = append(users, Users{Username: username, Password: password, isLoggedIn: true})
	return &Users{
		Username: username,
		Password: password,
	}, nil
}

func LoginUser(username string, password string) (*Users, error) {
	for i := 0; i < len(users); i++ {
		if users[i].Username == username || users[i].Password == password {
			users[i].isLoggedIn = true
			return &users[i], nil
		}
	}
	return nil, errors.New("username or password is incorrect")
}

func (u *Users) LogOut() {
	for i := 0; i < len(users); i++ {
		if users[i].Username == u.Username || users[i].Password == u.Password {
			users[i].isLoggedIn = false
		}
	}
}
