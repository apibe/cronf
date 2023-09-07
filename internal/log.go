package internal

import (
	"log"
	"os"
)

var F *os.File

func logInit() {
	logFile := "ftp.log"
	// 判断日志文件是否存在
	_, err := os.Stat(logFile)
	if os.IsNotExist(err) {
		_, err = os.Create(logFile)
		F, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("创建日志文件失败 %s", err.Error())
			return
		}
	} else {
		F, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("打开日志文件失败 %s", err.Error())
			return
		}
	}
	if !C.LOG.OUTPUT {
		log.SetOutput(F)
	}
}
