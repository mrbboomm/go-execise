package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFactory(t *testing.T) {
	validTypes := []string{"earn", "deduction"}
	
	for _, eType := range validTypes {
		t.Run("successfully created", func(t *testing.T) {
			_, err := createFactory(eType)
			assert.NoError(t, err)
		})	
	}

	t.Run("fail to create factory", func(t *testing.T) {
		_, err := createFactory("unknown")
		assert.EqualError(t, err, "\n\t=> Element error occured: Invalid Type --- unknown")
	})
}

func TestCreateElementCode(t *testing.T) {
	tests := []struct{
		input InputElementData
		hasError bool
		errMessage string
	}{
		{
			input: InputElementData{Name: "test", Code: "test", Type: "earn", Param: map[string]interface{}{"Value": 5}},
			hasError: false,
		},
		{
			input: InputElementData{Name: "test", Code: "test", Type: "earn2", Param: map[string]interface{}{"Value": 5}},
			hasError: true,
			errMessage: "\n\t=> Element error occured: \n\t=> Element error occured: Invalid Type --- earn2",
		},
		{
			input: InputElementData{Name: "test", Code: "test", Type: "deduction", Param: map[string]interface{}{"Value": 5}},
			hasError: true,
			errMessage: "\n\t=> Element error occured: Deduction Value must more than 0",
		},
	}

	for _, test := range tests {
		t.Run("create eleement code", func(t *testing.T) {
			_, err := CreateElementCode(&test.input)
			if test.hasError {
				assert.Error(t, err)
				assert.Equal(t, test.errMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}