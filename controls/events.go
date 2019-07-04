package controls

// ci events
const (
	passingColor = iota + 1
	passingScreen
	brokenColor
	brokenScreen
	buildingColor
	buildingScreen
	builtScreen
	builtTactile
	buildFailedScreen
	buildFailedTactile
)

// cd events
const (
	deployingColor = iota + 101
	deployingScreen
	pendingColor
	pendingScreen
	pendingTactile
	deployBrokenColor
	deployBrokenScreen
	deployedScreen
	deployedTactile
	deployFailedScreen
	deployFailedTactile
)
