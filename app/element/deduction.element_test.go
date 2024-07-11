package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeductionCalVal(t *testing.T) {
	tests := []struct {
		description string
		input *DeductionParam 
		expect float32 
	} {
		{
			description: "",
			input: &DeductionParam{Value: 10, DeductValue: 10},
			expect: 100,
		},
		{
			description: "",
			input: &DeductionParam{Value: -10, DeductValue: 10},
			expect: -100,
		},
		{
			description: "",
			input: &DeductionParam{Value: -10, DeductValue: 0},
			expect: 0,
		},
	}
	for _, test := range tests {
		t.Run("success", func(t *testing.T) {
			e := &DeductionElement{} 
			e.Param = test.input
			actual, _ := e.CalValue()
			assert.Equal(t, test.expect, actual)
		
		})
	}
}