package element

import "fmt"

func CreateElementCode(name string, elType string, param Param) (IElementCode, error) { 
	if elType == "earn" {
		return &EarningElement{BaseElement {Name: name, Code: name, ElType: elType, Param: param}}, nil
	} 
	if elType == "deduction" {
		return &DeductionElement{BaseElement {Name: name, Code: name, ElType: elType, Param: param}}, nil
	}

	return nil, fmt.Errorf("Invalid Element Type")


}

type mockElementData struct {
	name string
	elt string
	param Param
}

func LoadElements() []IElementCode {
	mockCodes := []mockElementData{
		{name: "e1", elt: "earn", param: Param{Value: 1}},
		{name: "e2", elt: "earn", param: Param{Value: 100}},
		{name: "e3", elt: "earn", param: Param{Value: 3}},
		{name: "d3", elt: "deduction", param: Param{Value: 3}},
	}
	codes := []IElementCode{}
	for i := 0; i < len(mockCodes); i++ {
		_m := mockCodes[i]
		_code, _er := CreateElementCode(_m.name, _m.elt, _m.param)
		if _er == nil {
			codes = append(codes, _code)
		}
	}
	return codes
}