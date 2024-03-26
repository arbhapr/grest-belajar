package src

import (
	"github.com/robfig/cron/v3"

	"grest-belajar/app"
)

func Scheduler() *schedulerUtil {
	if scheduler == nil {
		scheduler = &schedulerUtil{}
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			scheduler.Configure()
		}
		scheduler.isConfigured = true
	}
	return scheduler
}

var scheduler *schedulerUtil

type schedulerUtil struct {
	isConfigured bool
}

func (s *schedulerUtil) Configure() {
	c := cron.New()

	// add scheduler func here, for example :
	// c.AddFunc("CRON_TZ=Asia/Jakarta 5 0 * * *", app.Auth().RemoveExpiredToken)

	c.Start()
}
