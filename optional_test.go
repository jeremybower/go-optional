package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	t.Parallel()

	value := 42

	tests := []struct {
		name       string
		opt        Value[int]
		valid      bool
		value      int
		valueOrNil *int
	}{
		{"invalid", Invalid[int](), false, 0, nil},
		{"valid", New(value), true, value, &value},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Value[int]{Valid: tt.valid, Value: tt.value}, tt.opt)
			assert.Equal(t, tt.valueOrNil, tt.opt.ValueOrNil())
		})
	}
}

func TestValueUnmarshallJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		b           []byte
		initialVal  Value[string]
		expectedVal Value[string]
		expectedErr error
	}{
		{"invalid after invalid", []byte{}, Invalid[string](), Invalid[string](), nil},
		{"invalid after value", []byte{}, New("test"), Invalid[string](), nil},
		{"null after invalid", []byte("null"), Invalid[string](), Invalid[string](), ErrNullValue},
		{"null after value", []byte("null"), New("test"), New("test"), ErrNullValue},
		{"value after invalid", []byte(`"test"`), Invalid[string](), New("test"), nil},
		{"value after value", []byte(`"test2"`), New[string]("test1"), New("test2"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.initialVal
			err := opt.UnmarshalJSON(tt.b)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedVal, opt)
			}
		})
	}
}

func TestValueIfValid(t *testing.T) {
	t.Parallel()

	Value[int]{}.IfValid(func(v int) {
		t.Error("should not be called")
	})

	New(42).IfValid(func(v int) {
		assert.Equal(t, 42, v)
	})
}

func TestValueOr(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 42, Invalid[int]().Or(42))
	assert.Equal(t, 42, New(42).Or(43))
}
