package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEarningCalVal(t *testing.T) {
	tests := []struct {
		description string
		input *EarningParam 
		expect float32 
	} {
		{
			description: "success 1",
			input: &EarningParam{Value: 100},
			expect: 100,
		},
		{
			description: "success 2",
			input: &EarningParam{Value: -100},
			expect: -100,
		},
		{
			description: "success 3",
			input: &EarningParam{Value: -10},
			expect: -10,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			e := &EarningElement{} 
			e.Param = test.input
			actual, _ := e.CalValue()
			assert.Equal(t, test.expect, actual)
		
		})
	}
}