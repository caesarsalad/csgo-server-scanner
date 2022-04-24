package main

type ServerConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

type Config struct {
	ServerList  []ServerConfig `json:"serverList"`
	FavoriteMap string         `json:"favoriteMap"`
}
