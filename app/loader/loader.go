package loader

import (
	"encoding/json"
	"fmt"

	"github.com/mrbboomm/go-execise/app/element"
)


type ILoader interface {
	LoadElements() []element.IElementCode
}


type MockLoader struct {}


func (l *MockLoader) LoadElements() []element.IElementCode {
	mockCodes := []string{
		`{"Name": "e1", "Code": "e1", "Type": "earn", "Param": { "Value": 100 }}`,
		`{"Name": "e2", "Code": "e2", "Type": "earn", "Param": { "Value": 10 }}`,
		`{"Name": "e3", "Code": "e3", "Type": "earn", "Param": { "Value": 1 }}`,
		`{"Name": "d1", "Code": "d1", "Type": "deduction", "Param": { "DeductValue": 3, "Value": 9 }}`,
		`{"Name": "d2", "Code": "d2", "Type": "deduction", "Param": { "Value": 9 }}`,
		`{"Name": "d2", "Code": "d2", "Type": "ded", "Param": { "Value": 9 }}`,
	}

	codes := []element.IElementCode{}
	for i := 0; i < len(mockCodes); i++ {
		_dataString := mockCodes[i]
		var elData element.InputElementData
		if err := json.Unmarshal([]byte(_dataString), &elData); err != nil {
			fmt.Printf("Fail to convert data with error  %s \n", err)
			continue
		}
		_code, _er := element.CreateElementCode(&elData)
		if _er != nil {
			fmt.Printf("Fail to create el %d with error  %s \n", i, _er.Error())
			continue
		}
		codes = append(codes, _code)
	}
	return codes
}