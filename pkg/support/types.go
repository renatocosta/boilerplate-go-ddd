package support

import (
	"database/sql"
	"encoding/json"
)

type String struct {
	sql.NullString
}

func NewString(value string) String {
	return String{
		NullString: sql.NullString{
			String: value,
			Valid:  true,
		},
	}
}

func (s *String) Scan(value interface{}) error {
	return s.NullString.Scan(value)
}

func (ns *String) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}
