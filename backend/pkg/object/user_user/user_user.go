package user_user

import (
	"encoding/json"
	"matcha/backend/pkg/database/arangodb"
)

type RelationshipType string

const (
	RelationshipPass  RelationshipType = "pass"
	RelationshipSmash RelationshipType = "smash"
)

type UserUser struct {
	arangodb.EdgeDocument
	Relationship RelationshipType `json:"relationship"`
}

func (p UserUser) Name() string {
	return "user_users"
}

func (p UserUser) AsMap() (map[string]interface{}, error) {
	asBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	asMap := map[string]interface{}{}
	err = json.Unmarshal(asBytes, &asMap)
	return asMap, err
}
