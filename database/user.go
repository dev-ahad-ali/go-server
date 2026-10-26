package database

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

var users []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}

	users = append(users, u)
	return u
}
