package src

import "grest-belajar/app"

func Seeder() *seederUtil {
	if seeder == nil {
		seeder = &seederUtil{}
		seeder.Configure()
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			seeder.Run()
		}
		seeder.isConfigured = true
	}
	return seeder
}

var seeder *seederUtil

type seederUtil struct {
	isConfigured bool
}

func (s *seederUtil) Configure() {

}

func (s *seederUtil) Run() {

}
