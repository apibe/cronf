package initialize

import (
	"github.com/apibe/cronf/global"
	"github.com/robfig/cron"
)

func Cron() {
	global.Cron = cron.New()
}
