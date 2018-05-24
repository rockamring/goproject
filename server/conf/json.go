package conf

import (
	"io/ioutil"
	"github.com/rockamring/goproject/server/log"
	"encoding/json"
)

var Server struct{
	// log
	LogLevel	string
	LogPath		string

	// ws
	CertFile	string
	KeyFile		string

	// tcp
	TCPAddr		string

	MaxConnNum	int
	ConsolePort	int
	ProfilePath	string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}

	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}