package nullable

import (
	"bytes"
	"encoding/json"
)

type Value[T any] struct {
	Valid bool
	Value *T
}

func Invalid[T any]() Value[T] {
	return Value[T]{}
}

func New[T any](value *T) Value[T] {
	return Value[T]{Valid: true, Value: value}
}

func (v *Value[T]) Reset() {
	v.Value = nil
	v.Valid = false
}

func (v Value[T]) IfValid(fn func(v *T)) {
	if v.Valid {
		fn(v.Value)
	}
}

func (v *Value[T]) UnmarshalJSON(b []byte) error {
	if len(b) > 0 {
		v.Valid = true
		if bytes.Equal(b, []byte("null")) {
			v.Value = nil
			return nil
		}

		return json.Unmarshal(b, &v.Value)
	}

	v.Reset()
	return nil
}
