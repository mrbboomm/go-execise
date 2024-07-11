package element

import (
	"encoding/json"
	"fmt"
)

// custom error
type IFactory interface{
	Create(el BaseElement, param interface{}) (IElementCode, error)
}

// Register new type of element code here
func createFactory(elType string) (IFactory, error) {
	switch (elType){
		case "earn": 
			_f := EarningElementFactory{}
			return &_f, nil
		case "deduction":
			_f := DeductionElementFactory{}
			return &_f, nil
		}
	return nil, &ElementError{message: "Invalid Type"}
}
type ElementError struct {
	message string
}

func (e *ElementError) Error() string {
	return "\n	=> Element error occured: " + e.message
}

type InputElementData struct {
	Name string
	Type string
	Param any
}

func CreateElementCode(elementData *InputElementData) (IElementCode, error) { 
	elementFactory, factoryError := createFactory(elementData.Type)
	if factoryError != nil {
		return nil, &ElementError{message: factoryError.Error()} 
	}
	return elementFactory.Create(BaseElement{Name: elementData.Name, ElType: elementData.Type}, elementData.Param)

}


func LoadElements() []IElementCode {
	x := map[string]interface{}{"value": 3}
	fmt.Println(x["value"])
	mockCodes := []string{
		`{"Name": "e1", "Type": "earn", "Param": { "Value": 100 }}`,
		`{"Name": "e2", "Type": "earn", "Param": { "Value": 10 }}`,
		`{"Name": "e3", "Type": "earn", "Param": { "Value": 1 }}`,
		`{"Name": "d1", "Type": "deduction", "Param": { "DeductValue": 3, "Value": 9 }}`,
	}

	codes := []IElementCode{}
	for i := 0; i < len(mockCodes); i++ {
		_dataString := mockCodes[i]
		var elData InputElementData
		if err := json.Unmarshal([]byte(_dataString), &elData); err != nil {
			fmt.Printf("Fail to convert data with error  %s \n", err)
			continue
		}
		fmt.Println(elData.Name)
		_code, _er := CreateElementCode(&elData)
		if _er != nil {
			fmt.Printf("Fail to create el %s with error  %s \n", elData, _er.Error())
			continue
		}
		codes = append(codes, _code)
	}
	return codes
}