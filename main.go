package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/rumblefrog/go-a2s"
	"github.com/spf13/viper"
	"gopkg.in/toast.v1"
)

type ServerInfo struct {
	ServerName       string
	CurrentMap       string
	Players          int
	MaxPlayers       int
	PlayersInfo      string
	SendNotification bool
	IsSick           bool
}

var (
	conf *Config
)

// Desktop notification
func sendFavoriteMapNotification(serverName, serverPlayers, mapName string) {
	var (
		err  error
		path string
	)
	path, err = os.Getwd()
	if err != nil {
		log.Println(err)
	}
	message := "Server Name: " + serverName + "\nPlayers: " + serverPlayers
	title := mapName + " is playing now!"
	notification := toast.Notification{
		AppID:   "CS GO Favorite Server Scanner",
		Title:   title,
		Message: message,
		Icon:    path + "/ct.png",
	}
	err = notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func getConf() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		log.Panicf("%v", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Panicf("unable to decode into config struct, %v", err)
	}

	return conf
}

func init() {
	conf = getConf()
}

func main() {
	serverMap := make(map[string]ServerInfo)
	for {
		for _, server := range conf.ServerList {
			client, err := a2s.NewClient(server.Host)

			if err != nil {
				log.Println(err)
			}

			info, err := client.QueryInfo()
			client.Close()

			if err != nil {
				log.Println(err)
			}
			log.Println(info.Name, info.Map, info.Players, "/", info.MaxPlayers)

			_, existServer := serverMap[server.Host]
			if existServer && info.Map == serverMap[server.Host].CurrentMap {
				continue
			}

			serverMap[server.Host] = ServerInfo{
				ServerName: info.Name,
				CurrentMap: info.Map,
				Players:    int(info.Players),
				MaxPlayers: int(info.MaxPlayers),
			}

			if info.Map == conf.FavoriteMap && info.Players > 0 {
				log.Println("sending notification...")
				serverPlayers := strconv.Itoa(int(info.Players)) + "/" + strconv.Itoa(int(info.MaxPlayers))
				sendFavoriteMapNotification(server.Name, serverPlayers, info.Map)
			}
		}
		time.Sleep(time.Second * 30)
	}

}
