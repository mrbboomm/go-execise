package element


type DeductionElement struct {
	BaseElement
}

func (e *DeductionElement) CalValue()  (float32, error)  {
	return e.Param.Value, nil
}
