package runner

import (
	"fmt"

	"github.com/mrbboomm/go-execise/app/element"
	"github.com/mrbboomm/go-execise/app/loader"
	"github.com/mrbboomm/go-execise/app/types"
)
type ElementOutput struct {
	Element element.IElementCode
	Output float32
}


type Runner struct {
	elements []element.IElementCode
	status types.CalStatus
	result float32
	elementOutput []ElementOutput
	loader loader.ILoader
}

func (r *Runner) Load() []element.IElementCode {
	return r.loader.LoadElements()
}

func NewRunner(loader loader.ILoader) *Runner {
	return &Runner{status: types.Initial, result: 0, elementOutput: []ElementOutput{}, loader: loader}
}
func (r *Runner) Result() []ElementOutput {
	return r.elementOutput
}

func (r *Runner) LogResult() string {
	var resultStr = ""
	resultStr += fmt.Sprintln("==== Running Result ====")
	resultStr += fmt.Sprintf("status: %s \n", r.status)
	resultStr += fmt.Sprintf("result: %f \n", r.result)
	resultStr += fmt.Sprintln("runningElements")
	for i := 0; i < len(r.elementOutput); i++ {
		_el := r.elementOutput[i]
		resultStr += fmt.Sprintf("  code: %s type: %s val: %f\n", _el.Element.GetCode(), _el.Element.GetType(), _el.Output)
	}
	return resultStr
}

func (r *Runner) updateStatus(status types.CalStatus) {
	r.status = status
	fmt.Printf("⏩⏩⏩ cal state: %s \n", r.status)
}

func handleCal(r *Runner, el element.IElementCode) {
		_calVal, _calErr := el.CalValue()
		finalVal := float32(0)
		if _calErr != nil{
			fmt.Printf("Has error on %s with error message : %s\n", el.GetCode(), _calErr.Error())
		}
		if el.GetType() == "earn" {
			finalVal = r.result + _calVal
		} else if el.GetType() == "deduction" {
			if _bal := r.result - _calVal; _bal > 0 {
				finalVal = _bal
			} else {
				finalVal = 0
			}
		}
		r.result = finalVal
		r.elementOutput = append(r.elementOutput, 
		ElementOutput{
			Element: el,
			Output: _calVal,
		})
		fmt.Printf("  calulating for code %s: type: %s, val: %f  balance: %f\n", el.GetCode(), el.GetType(), _calVal, r.result)
	}

func (r *Runner) Run() {
	r.updateStatus(types.Initial)
	r.updateStatus(types.Loading)
	codes := r.loader.LoadElements()
	r.updateStatus(types.Calulating)
	for i := 0; i < len(codes) ; i++ {
		_c := codes[i]
		handleCal(r, _c)
	}
	r.updateStatus(types.Calculated)
	r.updateStatus(types.Completed)
}