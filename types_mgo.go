//This file adds support for mongodb null values.
//If you are already using mgo.v2 in your code using null.String or any other
//null type will not serialize/deserialize well to/from the database.
//This file is a dropin to help you with null values.

package null

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// GetBSON provides BSON encoding for String
func (str String) GetBSON() (interface{}, error) {

	if str.Valid {
		return str.String, nil
	}
	return nil, nil
}

// SetBSON provides BSON decoding for String
func (str *String) SetBSON(raw bson.Raw) error {

	var tmp string
	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	str.String = tmp
	str.Valid = true

	return nil
}

// GetBSON provides BSON encoding for Int
func (i Int) GetBSON() (interface{}, error) {

	if i.Valid {
		return i.Int64, nil
	}
	return nil, nil
}

// SetBSON provides BSON decoding for Int
func (i *Int) SetBSON(raw bson.Raw) error {

	var tmp int64
	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	i.Int64 = tmp
	i.Valid = true

	return nil
}

// GetBSON provides BSON encoding for Float
func (f Float) GetBSON() (interface{}, error) {

	if f.Valid {
		return f.Float64, nil
	}
	return nil, nil
}

// SetBSON provides BSON decoding for Float
func (f *Float) SetBSON(raw bson.Raw) error {

	var tmp float64
	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	f.Float64 = tmp
	f.Valid = true

	return nil
}

// GetBSON provides BSON encoding for Bool
func (b Bool) GetBSON() (interface{}, error) {

	if b.Valid {
		return b.Bool, nil
	}
	return nil, nil
}

// SetBSON provides BSON decoding for Bool
func (b *Bool) SetBSON(raw bson.Raw) error {

	var tmp bool
	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	b.Bool = tmp
	b.Valid = true

	return nil
}

// GetBSON provides BSON encoding for Time
func (t Time) GetBSON() (interface{}, error) {

	if t.Valid {
		return t.Time, nil
	}
	return nil, nil
}

// SetBSON provides BSON decoding for Time
func (t *Time) SetBSON(raw bson.Raw) error {

	var tmp time.Time
	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	t.Time = tmp
	t.Valid = true

	return nil
}
