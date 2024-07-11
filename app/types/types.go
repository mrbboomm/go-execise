package types

type CalStatus string

const (
	Initial CalStatus = "initial"
	Loading CalStatus = "loading"
	Calulating CalStatus = "calculating"
	Calculated CalStatus = "calculated"
	Completed CalStatus = "complete"
)