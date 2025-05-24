package users

import "errors"

type Users struct {
	Username   string
	Password   string
	IsLoggedIn bool
}

var Uusers []Users

func RegisterUser(username string, password string) (*Users, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password cannot be blank")
	}
	for i := 0; i < len(Uusers); i++ {
		if Uusers[i].Username == username {
			return nil, errors.New("there is already a user with this username")
		}
	}
	Uusers = append(Uusers, Users{Username: username, Password: password, IsLoggedIn: true})
	return &Users{
		Username:   username,
		Password:   password,
		IsLoggedIn: true,
	}, nil
}

func LoginUser(username string, password string) (*Users, error) {
<<<<<<< HEAD
	for i := range Uusers {
		if Uusers[i].Username == username || Uusers[i].Password == password {
			if !Uusers[i].IsLoggedIn {
				Uusers[i].IsLoggedIn = true
				return &Uusers[i], nil
=======
	for i := range users {
		if users[i].Username == username || users[i].Password == password {
			if !users[i].isLoggedIn {
				users[i].isLoggedIn = true
				return &users[i], nil
>>>>>>> master
			} else {
				return nil, errors.New("this user is already loggedin")
			}
		}
	}
	return nil, errors.New("username or password is incorrect")
}

func (u *Users) LogOut() {
<<<<<<< HEAD
	for i := range Uusers {
		if Uusers[i].Username == u.Username || Uusers[i].Password == u.Password {
			Uusers[i].IsLoggedIn = false
=======
	for i := range users {
		if users[i].Username == u.Username || users[i].Password == u.Password {
			users[i].isLoggedIn = false
>>>>>>> master
		}
	}
}
