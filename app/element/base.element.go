package element


type Param struct {
	Value float32
}

func (p *Param) Validate() bool {
	return p != nil
}

type BaseElement struct {
	Name string
	Code string
	ElType string
	Param Param
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
	CalValue() (float32, error)
}
