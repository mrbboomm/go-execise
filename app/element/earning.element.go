package element

import (
	"encoding/json"
)

type EarningElementFactory struct {
}

func (f *EarningElementFactory) Create(el BaseElement, param interface{}) (IElementCode, error) {
	element := &EarningElement{}
	element.BaseElement = &el
	// https://www.golinuxcloud.com/go-map-to-struct/#google_vignette
	data, _ := json.Marshal(param)
	json.Unmarshal(data, &element.Param)
	err := element.Param.Validate()
	if err != nil {
		return nil, err
	}
	return element, nil
}

type EarningParam struct {
	Value float32
}
func (p *EarningParam) Validate() error {
	return nil
}
func (p *EarningParam) GetValue() interface{} {
	return nil
}

type EarningElement struct {
	*BaseElement
	Param *EarningParam
}

func (e *EarningElement) CalValue() (float32, error) {
	val := e.Param.Value
	return val, nil
}