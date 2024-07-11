package element

import (
	"encoding/json"
	"fmt"
)


func Load() []IElementCode {
	mockCodes := []string{
		`{"Name": "e1", "Type": "earn", "Param": { "Value": 100 }}`,
		`{"Name": "e2", "Type": "earn", "Param": { "Value": 10 }}`,
		`{"Name": "e3", "Type": "earn", "Param": { "Value": 1 }}`,
		`{"Name": "d1", "Type": "deduction", "Param": { "DeductValue": 3, "Value": 9 }}`,
		`{"Name": "d2", "Type": "deduction", "Param": { "Value": 9 }}`,
		`{"Name": "d2", "Type": "ded", "Param": { "Value": 9 }}`,
	}

	codes := []IElementCode{}
	for i := 0; i < len(mockCodes); i++ {
		_dataString := mockCodes[i]
		var elData InputElementData
		if err := json.Unmarshal([]byte(_dataString), &elData); err != nil {
			fmt.Printf("Fail to convert data with error  %s \n", err)
			continue
		}
		_code, _er := CreateElementCode(&elData)
		if _er != nil {
			fmt.Printf("Fail to create el %d with error  %s \n", i, _er.Error())
			continue
		}
		codes = append(codes, _code)
	}
	return codes
}