package rival710

import (
	"fmt"
	"log"
	"os"

	"github.com/yi-zhang/rival-710-extreme-feedback/utils"

	// load environment variables
	_ "github.com/joho/godotenv/autoload"
)

const (
	game          = "CI_CD"
	keepAliveTime = float64(15000)
)

var (
	propsPath = os.Getenv("CORE_PROPS_PATH")
	bindAPI,
	triggerAPI,
	removeEventAPI,
	removeGameAPI,
	keepAliveAPI string
)

func init() {
	loadAPIs()
}

func loadAPIs() {
	parsed, ok := utils.ParseJSONFile(propsPath)
	if !ok {
		log.Fatalln("Cannot load SteelSeries Engine configurations.")
	}
	address, ok := parsed["address"]
	if !ok {
		log.Fatalln("Missing host address for SteelSeries Engine.")
	}
	host := fmt.Sprintf("http://%s", address)
	bindAPI = fmt.Sprintf("%s/bind_game_event", host)
	triggerAPI = fmt.Sprintf("%s/game_event", host)
	removeEventAPI = fmt.Sprintf("%s/remove_game_event", host)
	removeGameAPI = fmt.Sprintf("%s/remove_game", host)
	keepAliveAPI = fmt.Sprintf("%s/game_heartbeat", host)
}
