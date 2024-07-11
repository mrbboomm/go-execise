package element

import (
	"encoding/json"
)

type DeductionElementFactory struct {
}

func (f *DeductionElementFactory) Create(el BaseElement, param interface{}) (IElementCode, error) {
	element := &DeductionElement{}
	element.BaseElement = &el
	data, _ := json.Marshal(param)
	json.Unmarshal(data, &element.Param)
	println("json", element.Param)
	err := element.Param.Validate()
	if err != nil {
		return nil, err
	}
	return element, nil
}

type DeductionParam struct {
	DeductValue float32
	Value float32
}

func (p *DeductionParam) Validate() error {
	if p.DeductValue < 0 {
		return &ElementError{message: "Deduction Value must more than 0"}
	}
	return nil
}

func (p *DeductionParam) GetValue() interface{} {
	return p.DeductValue
}

type DeductionElement struct {
	*BaseElement
	Param *DeductionParam
}

func (e *DeductionElement) CalValue()  (float32, error)  {
	_p := e.Param.DeductValue * e.Param.Value
	return _p, nil
}
