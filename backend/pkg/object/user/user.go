package user

import "encoding/json"

type User struct {
	name      string
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
}

func (u User) Name() string {
	return "users"
}

func (u User) AsMap() (map[string]interface{}, error) {
	asBytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	asMap := map[string]interface{}{}
	err = json.Unmarshal(asBytes, &asMap)
	return asMap, err
}
