package controls

// ci events
const (
	PassingColor = iota + 1
	PassingScreen
	BrokenColor
	BrokenScreen
	BuildingColor
	BuildingScreen
	BuiltScreen
	BuiltTactile
	BuildFailedScreen
	BuildFailedTactile
)

// cd events
const (
	DeployingColor = iota + 101
	DeployingScreen
	PendingColor
	PendingScreen
	PendingTactile
	DeployBrokenColor
	DeployBrokenScreen
	DeployedScreen
	DeployedTactile
	DeployFailedScreen
	DeployFailedTactile
)
