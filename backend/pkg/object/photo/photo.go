package photo

import (
	"encoding/json"
	"matcha/backend/pkg/database/arangodb"
)

type Photo struct {
	arangodb.Document
	B64 string `json:"b64"`
}

func (p Photo) Name() string {
	return "photos"
}

func (p Photo) AsMap() (map[string]interface{}, error) {
	asBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	asMap := map[string]interface{}{}
	err = json.Unmarshal(asBytes, &asMap)
	return asMap, err
}
