package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type RoleEnum string

const (
	UserEnum  RoleEnum = "user"
	AdminEnum RoleEnum = "admin"
)

func (r *RoleEnum) Scan(value interface{}) error {
	*r = RoleEnum(value.([]byte))
	return nil
}

func (r RoleEnum) Value() (driver.Value, error) {
	return string(r), nil
}

type JSON json.RawMessage

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type Image struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
