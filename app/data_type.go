package app

import (
	"time"

	"github.com/google/uuid"
	"grest.dev/grest"
)

// NullBool is nullable Bool, it embeds a grest.NullBool, so you can add your own method or override grest.NullBool method
type NullBool struct {
	grest.NullBool
}

// NewNullBool return NullBool
func NewNullBool(val bool) NullBool {
	n := NullBool{}
	n.Valid = true
	n.Bool = val
	return n
}

// NullInt64 is nullable Int64, it embeds a grest.NullInt64, so you can add your own method or override grest.NullInt64 method
type NullInt64 struct {
	grest.NullInt64
}

// NewNullInt64 return NullInt64
func NewNullInt64(val int64) NullInt64 {
	n := NullInt64{}
	n.Valid = true
	n.Int64 = val
	return n
}

// NullFloat64 is nullable Float64, it embeds a grest.NullFloat64, so you can add your own method or override grest.NullFloat64 method
type NullFloat64 struct {
	grest.NullFloat64
}

// NewNullFloat64 return NullFloat64
func NewNullFloat64(val float64) NullFloat64 {
	n := NullFloat64{}
	n.Valid = true
	n.Float64 = val
	return n
}

// NullDate is nullable Date, it embeds a grest.NullDate, so you can add your own method or override grest.NullDate method
type NullDate struct {
	grest.NullDate
}

// NewNullDate return NullDate
func NewNullDate(val time.Time) NullDate {
	n := NullDate{}
	n.Valid = true
	n.Time = val
	return n
}

// NullTime is nullable Time, it embeds a grest.NullTime, so you can add your own method or override grest.NullTime method
type NullTime struct {
	grest.NullTime
}

// NewNullTime return NullTime
func NewNullTime(val time.Time) NullTime {
	n := NullTime{}
	n.Valid = true
	n.String = val.Format("15:04:05")
	return n
}

// NullDateTime is nullable DateTime, it embeds a grest.NullDateTime, so you can add your own method or override grest.NullDateTime method
type NullDateTime struct {
	grest.NullDateTime
}

// NewNullDateTime return NullDateTime
func NewNullDateTime(val time.Time) NullDateTime {
	n := NullDateTime{}
	n.Valid = true
	n.Time = val
	return n
}

// NullUnixTime is nullable Unix DateTime, it embeds a grest.NullUnixTime, so you can add your own method or override grest.NullUnixTime method
type NullUnixTime struct {
	grest.NullUnixTime
}

// NewNullUnixTime return NullUnixTime
func NewNullUnixTime(val time.Time) NullUnixTime {
	n := NullUnixTime{}
	n.Valid = true
	n.Time = val
	return n
}

// NullString is nullable String, it embeds a grest.NullString, so you can add your own method or override grest.NullString method
type NullString struct {
	grest.NullString
}

// NewNullString return NullString
func NewNullString(val string) NullString {
	n := NullString{}
	n.Valid = true
	n.String = val
	return n
}

// NullText is nullable Text, it embeds a grest.NullText, so you can add your own method or override grest.NullText method
type NullText struct {
	grest.NullText
}

// NewNullText return NullText
func NewNullText(val string) NullText {
	n := NullText{}
	n.Valid = true
	n.String = val
	return n
}

// NullUUID is nullable UUID, it embeds a grest.NullUUID, so you can add your own method or override grest.NullUUID method
type NullUUID struct {
	grest.NullUUID
}

// NewNullUUID return NullUUID
func NewNullUUID(val ...string) NullUUID {
	n := NullUUID{}
	n.Valid = true
	if len(val) > 0 {
		n.String = val[0]
	} else {
		n.String = uuid.NewString()
	}
	return n
}

// NullJSON is nullable JSON, it embeds a grest.NullJSON, so you can add your own method or override grest.NullJSON method
type NullJSON struct {
	grest.NullJSON
}
