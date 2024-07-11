package element

// custom error
type IFactory interface{
	Create(el BaseElement, param interface{}) (IElementCode, error)
}

type InputElementData struct {
	Name string
	Code string
	Type string
	Param any
}

// Register new type of element code here
func createFactory(elementType string) (IFactory, error) {
	switch (elementType){
		case "earn": 
			_f := EarningElementFactory{}
			return &_f, nil
		case "deduction":
			_f := DeductionElementFactory{}
			return &_f, nil
		}
	
	return nil, &ElementError{message: "Invalid Type --- " + elementType}
}

func CreateElementCode(elementData *InputElementData) (IElementCode, error) { 
	elementFactory, factoryError := createFactory(elementData.Type)
	if factoryError != nil {
		return nil, &ElementError{message: factoryError.Error()} 
	}
	return elementFactory.Create(BaseElement{Name: elementData.Name, ElType: elementData.Type}, elementData.Param)

}