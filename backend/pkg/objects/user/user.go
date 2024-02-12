package user

import "matcha/backend/pkg/objects"

type User struct {
	objects.Object
	name      string
	Key       string `json:"_key"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (u User) Name() string {
	return "users"
}
