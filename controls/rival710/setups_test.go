package rival710

import (
	"fmt"
	"testing"
)

var _ = func() int {
	propsPath = "../../testdata/testcorepath.json"
	return 1
}()

func TestLoadAPIs(t *testing.T) {
	var (
		host = "http://localhost:51874/"
		apis = []*string{
			&bindAPI,
			&triggerAPI,
			&removeEventAPI,
			&removeGameAPI,
			&keepAliveAPI,
		}
		urls = []string{
			"bind_game_event",
			"game_event",
			"remove_game_event",
			"remove_game",
			"game_heartbeat",
		}
	)
	for i, api := range apis {
		expected := fmt.Sprintf("%s%s", host, urls[i])
		if *api != expected {
			t.Errorf("Expected %s, got %s", expected, *api)
		}
	}
}
