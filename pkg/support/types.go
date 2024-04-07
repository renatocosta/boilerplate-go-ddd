package support

import "database/sql"

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
