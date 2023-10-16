package global

import (
	"github.com/apibe/cronf/config"
	"github.com/robfig/cron"
)

var C *config.Config

var Cron *cron.Cron
