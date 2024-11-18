package streamdocksdk

import (
	"encoding/json"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/session"
	"strconv"
)

type Info struct {
	Application struct {
		Font            string `json:"font"`
		Language        string `json:"language"`
		Platform        string `json:"platform"`
		PlatformVersion string `json:"platformVersion"`
		Version         string `json:"version"`
	} `json:"application"`
	Colors struct {
		ButtonMouseOverBackgroundColor string `json:"buttonMouseOverBackgroundColor"`
		ButtonPressedBackgroundColor   string `json:"buttonPressedBackgroundColor"`
		ButtonPressedBorderColor       string `json:"buttonPressedBorderColor"`
		ButtonPressedTextColor         string `json:"buttonPressedTextColor"`
		HighlightColor                 string `json:"highlightColor"`
	} `json:"colors"`
	DevicePixelRatio int `json:"devicePixelRatio"`
	Devices          []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Size struct {
			Columns int `json:"columns"`
			Rows    int `json:"rows"`
		} `json:"size"`
		Type int `json:"type"`
	} `json:"devices"`
	Plugin struct {
		Uuid    string `json:"uuid"`
		Version string `json:"version"`
	} `json:"plugin"`
}

func Run(argv []string, plugin session.Plugin) int {
	var port int
	var pluginUUID, registerEvent, info string
	for i := 0; i+1 < len(argv); i = i + 2 {
		switch argv[i] {
		case "-port":
			port, _ = strconv.Atoi(argv[i+1])
		case "-pluginUUID":
			pluginUUID = argv[i+1]
		case "-registerEvent":
			registerEvent = argv[i+1]
		case "-info":
			info = argv[i+1]
		default:
			log.Message("Unknown argument: " + argv[i])
		}
	}
	if port == 0 || pluginUUID == "" || registerEvent == "" {
		log.Message("port=%d, pluginUUID=%s, registerEvent=%s is required", port, pluginUUID, registerEvent)
		return 1
	}
	var pInfo = &Info{}
	err := json.Unmarshal([]byte(info), pInfo)
	if err != nil {
		log.Message(err.Error())
		return 1
	}
	conn := session.NewConnectionManager(port, pluginUUID, registerEvent, plugin)
	if conn == nil {
		log.Message("NewConnectionManager failed")
		return 1
	}
	conn.Run()
	return 0
}