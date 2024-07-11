package element

type BaseElement struct {
	Name string
	Code string
	ElType string
}


func (e *BaseElement) GetType() string {
	return e.ElType
}

func (e *BaseElement) GetCode() string {
	return e.Code
}

func (e *BaseElement) GetName() string {
	return e.Code
}


type IElementCode interface {
	GetType() string
	GetCode() string
	GetName() string
	/**
	* Calculate value
	**/
	CalValue() (float32, error)
}
