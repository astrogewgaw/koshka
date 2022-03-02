package data

import (
	"database/sql"
	"encoding/json"
)

type NullBool struct{ sql.NullBool }
type NullInt struct{ sql.NullInt64 }
type NullFloat struct{ sql.NullFloat64 }
type NullString struct{ sql.NullString }

func (NB NullBool) MarshalJSON() ([]byte, error) {
	if NB.Valid {
		return json.Marshal(NB.Bool)
	}
	return json.Marshal(nil)
}

func (NB *NullBool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		NB.Bool = *b
		NB.Valid = true
	} else {
		NB.Valid = false
	}
	return nil
}

func (NI NullInt) MarshalJSON() ([]byte, error) {
	if NI.Valid {
		return json.Marshal(NI.Int64)
	}
	return json.Marshal(nil)
}

func (NI *NullInt) UnmarshalJSON(data []byte) error {
	var i *int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		NI.Int64 = *i
		NI.Valid = true
	} else {
		NI.Valid = false
	}
	return nil
}

func (NF NullFloat) MarshalJSON() ([]byte, error) {
	if NF.Valid {
		return json.Marshal(NF.Float64)
	}
	return json.Marshal(nil)
}

func (NF *NullFloat) UnmarshalJSON(data []byte) error {
	var f *float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	if f != nil {
		NF.Float64 = *f
		NF.Valid = true
	} else {
		NF.Valid = false
	}
	return nil
}

func (NS NullString) MarshalJSON() ([]byte, error) {
	if NS.Valid {
		return json.Marshal(NS.String)
	}
	return json.Marshal(nil)
}

func (NS *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		NS.String = *s
		NS.Valid = true
	} else {
		NS.Valid = false
	}
	return nil
}
