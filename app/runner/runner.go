package runner

import (
	"fmt"

	"github.com/mrbboomm/go-execise/app/element"
	"github.com/mrbboomm/go-execise/app/types"
)
type ElementOutput struct {
	Element element.IElementCode
	Output float32
}



type Runner struct {
	status types.CalStatus
	result float32
	elementOutput []ElementOutput
}

type PayrollBalance float32

func New() *Runner {
	return &Runner{status: types.Initial, result: 0, elementOutput: []ElementOutput{}}
}

func (r *Runner) Result() {
	fmt.Println("==== Running Result ====")
	fmt.Printf("status: %s \n", r.status)
	fmt.Printf("result: %f \n", r.result)
	fmt.Println("runningElements")
	for i := 0; i < len(r.elementOutput); i++ {
		_el := r.elementOutput[i]
		fmt.Printf("  code: %s type: %s val: %f\n", _el.Element.GetCode(), _el.Element.GetType(), _el.Output)
	}
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
	codes := element.LoadElements()
	r.updateStatus(types.Calulating)
	for i := 0; i < len(codes) ; i++ {
		_c := codes[i]
		handleCal(r, _c)
	}
	r.updateStatus(types.Calculated)
	r.updateStatus(types.Completed)
}

