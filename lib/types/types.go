package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// RawNullJSON ...
type RawNullJSON struct {
	JSON  []byte
	Valid bool
}

// Scan ...
func (nt *RawNullJSON) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case []byte:
		nt.JSON = v
		nt.Valid = true
		return
	}

	nt.Valid = false
	return fmt.Errorf("%s", value)
}

// Value implements the driver Valuer interface.
func (nt RawNullJSON) Value() (driver.Value, error) {
	if nt.JSON == nil {
		return nil, nil
	}
	return nt.JSON, nil
}

// CheckNullRawJSON ...
func CheckNullRawJSON(j *json.RawMessage) RawNullJSON {
	var nj RawNullJSON
	if j != nil {
		nj = RawNullJSON{
			JSON:  []byte(*j),
			Valid: true,
		}
	} else {
		nj = RawNullJSON{
			JSON:  []byte(nil),
			Valid: false,
		}
	}
	return nj
}

// CheckNullString ...
func CheckNullString(s *string) sql.NullString {
	var c sql.NullString
	if s != nil {
		c = sql.NullString{
			String: string(*s),
			Valid:  true,
		}
	} else {
		c = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	return c
}

// CheckNullInt32 ...
func CheckNullInt32(s *int32) sql.NullInt32 {
	var c sql.NullInt32

	if s != nil {
		c = sql.NullInt32{
			Int32: int32(*s),
			Valid: true,
		}
	} else {
		c = sql.NullInt32{
			Int32: 0,
			Valid: false,
		}
	}

	return c
}

// CheckNullFloat64 ...
func CheckNullFloat64(s *float64) sql.NullFloat64 {
	var c sql.NullFloat64

	if s != nil {
		c = sql.NullFloat64{
			Float64: float64(*s),
			Valid:   true,
		}
	} else {
		c = sql.NullFloat64{
			Float64: 0,
			Valid:   false,
		}
	}

	return c
}

// CheckNullBool ...
func CheckNullBool(b *bool) sql.NullBool {
	var c sql.NullBool
	if b != nil {
		c = sql.NullBool{
			Bool:  *b,
			Valid: true,
		}
	} else {
		c = sql.NullBool{
			Bool:  false,
			Valid: false,
		}
	}

	return c
}

// CheckNullTime ...
func CheckNullTime(t *time.Time) sql.NullTime {
	var c sql.NullTime
	if t != nil {
		c = sql.NullTime{
			Time:  *t,
			Valid: true,
		}
	} else {
		c = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	}

	return c
}
