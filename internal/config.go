package internal

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	FTP struct {
		ADDR     string `json:"ADDR"`
		USERNAME string `json:"USERNAME"`
		PASSWORD string `json:"PASSWORD"`
	} `json:"FTP"`
	CRONDOWNLOAD []struct {
		NAME      string `json:"NAME"`
		EXEC      bool   `json:"EXEC"`
		JOINCRON  bool   `json:"JOIN_CRON"`
		CRON      string `json:"CRON"`
		FTPPATH   string `json:"FTP_PATH"`
		LOCALPATH string `json:"LOCAL_PATH"`
		RETRY     struct {
			INTERVAL int `json:"INTERVAL"`
			TIMES    int `json:"TIMES"`
		} `json:"RETRY"`
	} `json:"CRON_DOWNLOAD"`
	LOG struct {
		OUTPUT bool `json:"OUTPUT"`
	} `json:"LOG"`
}

var C config

// ConfigInit 读取配置文件
func configInit() {
	// 1. 读取 ./config.json
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("configuration loading failed!!")
	}
	_ = json.Unmarshal(bytes, &C)
}
