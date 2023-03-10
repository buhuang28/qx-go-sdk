package main

import (
	"encoding/json"
	"qx-go-sdk/log_init"
	"qx-go-sdk/qx"
)

func init() {
	log_init.InitLog("./", 7)
	qx.Info = Info
}

func Info() string {
	m := make(map[string]string)
	m["name"] = "demo"
	m["author"] = "不慌"
	m["description"] = "golang sdk"
	m["version"] = "1.0.0"
	m["skey"] = "QXASDOG5SD7A8SD"
	m["sdk"] = "S1"
	marshal, _ := json.Marshal(m)
	return string(marshal)
}

func main() {}
