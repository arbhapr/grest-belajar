package src

import (
	"grest-belajar/app"
	"grest-belajar/middleware"
)

func Middleware() *middlewareUtil {
	if mdlwr == nil {
		mdlwr = &middlewareUtil{}
		mdlwr.Configure()
		mdlwr.isConfigured = true
	}
	return mdlwr
}

var mdlwr *middlewareUtil

type middlewareUtil struct {
	isConfigured bool
}

func (*middlewareUtil) Configure() {
	app.Server().AddMiddleware(middleware.Ctx().New)
	app.Server().AddMiddleware(middleware.DB().New)
}
