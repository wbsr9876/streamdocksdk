package streamdocksdk

import (
	"encoding/json"
	"flag"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/proto"
	"github.com/wbsr9876/streamdocksdk/session"
)

func Run(plugin session.Plugin) int {
	port := flag.Int("port", 0, "port")
	pluginUUID := flag.String("pluginUUID", "", "pluginUUID")
	registerEvent := flag.String("registerEvent", "", "registerEvent")
	info := flag.String("info", "", "info")
	flag.Parse()

	if port == nil || pluginUUID == nil || registerEvent == nil {
		flag.PrintDefaults()
		//log.Message("port=%d, pluginUUID=%s, registerEvent=%s is required", port, pluginUUID, registerEvent)
		return 1
	}

	if info != nil {
		var pInfo = &proto.Info{}
		err := json.Unmarshal([]byte(*info), pInfo)
		if err != nil {
			log.Message(err.Error())
			return 1
		}
		plugin.SetInfo(pInfo)
	}

	conn := session.NewConnectionManager(*port, *pluginUUID, *registerEvent, plugin)
	if conn == nil {
		log.Message("NewConnectionManager failed")
		return 1
	}
	conn.Run()
	return 0
}
