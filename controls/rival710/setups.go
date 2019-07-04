package rival710

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yi-zhang/rival-710-extreme-feedback/utils"

	// load environment variables
	_ "github.com/joho/godotenv/autoload"
)

const (
	game          = "CI_CD"
	keepAliveTime = 15000
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
	keepAlive()
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

func keepAlive() {
	var (
		loop     func()
		interval = time.Duration(utils.MaxInt(5000, keepAliveTime-5000))
	)
	loop = func() {
		data := map[string]interface{}{"game": game}
		if utils.PostJSON(keepAliveAPI, data) {
			log.Println("Keep alive request sent.")
		}
		select {
		case <-time.After(interval * time.Millisecond):
			loop()
		}
	}
	go loop()
}
