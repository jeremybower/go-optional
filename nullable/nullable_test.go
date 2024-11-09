package nullable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	t.Parallel()

	str := "test"

	tests := []struct {
		name  string
		opt   Value[string]
		valid bool
		value *string
	}{
		{"invalid", Invalid[string](), false, nil},
		{"valid nil", New[string](nil), true, nil},
		{"valid non-nil", New(&str), true, &str},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Value[string]{Valid: tt.valid, Value: tt.value}, tt.opt)
		})
	}
}

func TestValueUnmarshallJSON(t *testing.T) {
	t.Parallel()

	str1 := "test1"
	str2 := "test2"

	tests := []struct {
		name        string
		b           []byte
		initialVal  Value[string]
		expectedVal Value[string]
		expectedErr error
	}{
		{"invalid after invalid", []byte{}, Invalid[string](), Invalid[string](), nil},
		{"invalid after value", []byte{}, New(&str1), Invalid[string](), nil},
		{"null after invalid", []byte("null"), Invalid[string](), New[string](nil), nil},
		{"null after value", []byte("null"), New(&str1), New[string](nil), nil},
		{"value after invalid", []byte(`"test1"`), Invalid[string](), New(&str1), nil},
		{"value after value", []byte(`"test2"`), New[string](&str1), New(&str2), nil},
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

	str := "test"

	Value[string]{}.IfValid(func(v *string) {
		t.Error("should not be called")
	})

	New(&str).IfValid(func(v *string) {
		assert.Equal(t, &str, v)
	})
}

func TestValueOr(t *testing.T) {
	t.Parallel()

	str := "test"

	assert.Nil(t, Invalid[string]().Or(nil))
	assert.Equal(t, &str, Invalid[string]().Or(&str))
	assert.Equal(t, &str, New(&str).Or(nil))
}
