package element

type EarningElement struct {
	BaseElement
}

func (e *EarningElement) CalValue() (float32, error) {
	return e.Param.Value, nil
}