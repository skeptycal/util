package main

const (
	// arbitrary 1 ppt tolerance // TODO - move to config
	precisionConstant = onePPT
)

// tolerance constants describe allowable
// tolerances for estimates
type tolerance = float64

// func (t *tolerance) String() string {
// 	return fmt.Sprintf("%.3F", t)
// }

const (
	fivePercent = 5e-2
	twoPercent  = 2e-2
	onePercent  = 1e-2
	onePPT      = 1e-3
	onePPM      = 1e-6
	onePPB      = 1e-9
)
