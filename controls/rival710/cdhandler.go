package rival710

import (
	"log"

	"github.com/yi-zhang/rival-710-extreme-feedback/controls"
	"github.com/yi-zhang/rival-710-extreme-feedback/utils"
)

func (c Controller) executeCd(mode string, data map[string]interface{}) {
	switch mode {
	case "deploying":
		c.handleDeploying(data)
	case "pending":
		c.handlePending(data)
	case "deploy-broken":
		c.handleDeployBroken(data)
	case "deployed":
		c.handleDeployed(data)
	case "deploy-failed":
		c.handleDeployFailed(data)
	}
}

func (c Controller) handleDeploying(data map[string]interface{}) {
	if branch := tryGetBranch(data); branch != "" {
		screenMeta := newScreenMeta(utils.StringPrefixSum(branch), 0, "", false)
		c.Screen.applyShift(*screenMeta, 100, controls.DeployingScreen)
		colorMeta := newColorMeta(nil, 0)
		c.Color.applyBreath(*colorMeta, controls.DeployingColor)
	}
}

func (c Controller) handlePending(data map[string]interface{}) {
	if branch := tryGetBranch(data); branch != "" {
		screenMeta := newScreenMeta([]string{branch, ""}, 0, "", false)
		c.Screen.applyShift(*screenMeta, 650, controls.PendingScreen)
		c.Color.applyBlink(*newColorMeta(nil, 1), controls.PendingColor)
	}
}

func (c Controller) handleDeployBroken(data map[string]interface{}) {
	if branch := tryGetBranch(data); branch != "" {
		screenMeta := newScreenMeta([]string{branch, ""}, 0, "", false)
		c.Screen.applyShift(*screenMeta, 400, controls.DeployBrokenScreen)
		c.Color.applyBlink(*newColorMeta(nil, 2), controls.DeployBrokenColor)
	}
}

func tryGetBranch(data map[string]interface{}) string {
	branch, ok := utils.StringFromMap(data, "branch")
	if !ok {
		log.Print("Missing branch information.")
		return ""
	}
	return branch
}

func (c Controller) handleDeployed(data map[string]interface{}) {
	c.handleComplete(data, controls.DeployedScreen, controls.DeployedTactile)
}

func (c Controller) handleDeployFailed(data map[string]interface{}) {
	c.handleComplete(data, controls.DeployFailedScreen, controls.DeployFailedTactile)
}
