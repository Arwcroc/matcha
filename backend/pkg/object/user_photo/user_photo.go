package user_photo

import (
	"encoding/json"
	"matcha/backend/pkg/database/arangodb"
)

type UserPhoto struct {
	arangodb.EdgeDocument
	Index int `json:"index"`
}

func (p UserPhoto) Name() string {
	return "user_photos"
}

func (p UserPhoto) AsMap() (map[string]interface{}, error) {
	asBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	asMap := map[string]interface{}{}
	err = json.Unmarshal(asBytes, &asMap)
	return asMap, err
}
