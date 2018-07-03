package main

import (
	"encoding/json"
	"fmt"
)

//Server is a json format
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
	Port       int    `json:"port,string"`
	Version    string `json:"version,omitempty"`
}

//Serverslice is a json format
type Serverslice struct {
	ID string `json:"-"` // - 将不会处理

	Servers []Server `json:"servers"`
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{"beijing_vpn", "127.0.0.1", 22, "v0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "chegongzhuang", ServerIP: "127.0.0.2", Port: 443})

	b, err := json.Marshal(s)
	checkError(err)
	fmt.Println(string(b))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
