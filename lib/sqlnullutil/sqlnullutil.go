package sqlnullutil

import (
	"database/sql"
	"strconv"
	"time"
)

// NewNullString creates a sql.NullString struct from string
func NewNullString(val string, notNull bool) sql.NullString {
	var ns sql.NullString
	if val != "" || notNull {
		ns.String, ns.Valid = val, true
	}

	return ns
}

// NewNullInt64 creates a sql.NullInt64 struct from int64
func NewNullInt64(val int64, notNull bool) sql.NullInt64 {
	var ni sql.NullInt64
	if val != 0 || notNull {
		ni.Int64, ni.Valid = val, true
	}

	return ni
}

// NewNullInt32 creates a sql.NullInt32 struct from int32
func NewNullInt32(val int32, notNull bool) sql.NullInt32 {
	var ni sql.NullInt32
	if val != 0 || notNull {
		ni.Int32, ni.Valid = val, true
	}

	return ni
}

// NewNullFloat64 creates a sql.NullFloat64 struct from float64
func NewNullFloat64(val float64, notNull bool) sql.NullFloat64 {
	var nf sql.NullFloat64
	if val != 0 || notNull {
		nf.Float64, nf.Valid = val, true
	}

	return nf
}

// NewNullBool creates a sql.NullBool struct from float64
func NewNullBool(val bool, notNull bool) sql.NullBool {
	var nb sql.NullBool
	if val == true || notNull {
		nb.Bool, nb.Valid = val, true
	}

	return nb
}

// NewNullTime creates a sql.NullTime struct from time.Time
func NewNullTime(val time.Time, notNull bool) sql.NullTime {
	var nt sql.NullTime
	if (val != time.Time{}) || notNull {
		nt.Time, nt.Valid = val, true
	}

	return nt
}

// GetString gets string value from sql.NullString
func GetString(nullval sql.NullString) string {
	if nullval.Valid {
		return nullval.String
	}

	return ""
}

// GetInt64 gets int64 value from sql.NullInt64
func GetInt64(nullval sql.NullInt64) int64 {
	if nullval.Valid {
		return nullval.Int64
	}

	return 0
}

// GetInt32 gets int32 value from sql.NullInt32
func GetInt32(nullval sql.NullInt32) int32 {
	if nullval.Valid {
		return nullval.Int32
	}

	return 0
}

// GetFloat64 gets float64 value from sql.NullFloat64
func GetFloat64(nullval sql.NullFloat64) float64 {
	if nullval.Valid {
		return nullval.Float64
	}

	return 0
}

// GetBool gets bool value from sql.NullBool
func GetBool(nullval sql.NullBool) bool {
	if nullval.Valid {
		return nullval.Bool
	}

	return false
}

// GetTime gets time.Time value from sql.NullTime
func GetTime(nullval sql.NullTime) time.Time {
	if nullval.Valid {
		return nullval.Time
	}

	return time.Time{}
}

// ConvertNullStringToFloat64 converts sql.NullString to float64
func ConvertNullStringToFloat64(nullval sql.NullString) float64 {
	if nullval.Valid {
		val, err := strconv.ParseFloat(nullval.String, 64)
		if err != nil {
			return 0
		}

		return val
	}

	return 0
}

// ConvertNullStringToFloat32 converts sql.NullString to float32
func ConvertNullStringToFloat32(nullval sql.NullString) float32 {
	if nullval.Valid {
		val, err := strconv.ParseFloat(nullval.String, 32)
		if err != nil {
			return 0
		}

		return float32(val)
	}

	return 0
}

// ConvertNullStringToInt64 converts sql.NullString to int64
func ConvertNullStringToInt64(nullval sql.NullString) int64 {
	if nullval.Valid {
		val, err := strconv.ParseInt(nullval.String, 10, 64)
		if err != nil {
			return 0
		}

		return int64(val)
	}

	return 0
}

// ConvertNullStringToInt32 converts sql.NullString to int32
func ConvertNullStringToInt32(nullval sql.NullString) int32 {
	if nullval.Valid {
		val, err := strconv.ParseInt(nullval.String, 10, 32)
		if err != nil {
			return 0
		}

		return int32(val)
	}

	return 0
}

// ConvertNullTimeToString converts sql.NullTime to string
func ConvertNullTimeToString(nullval sql.NullTime, layout string) string {
	if nullval.Valid {
		return nullval.Time.Format(layout)
	}

	return ""
}
