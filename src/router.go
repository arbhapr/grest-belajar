package src

import (
	"grest-belajar/app"
	"grest-belajar/src/category"
	"grest-belajar/src/user"
	// import : DONT REMOVE THIS COMMENT
)

func Router() *routerUtil {
	if router == nil {
		router = &routerUtil{}
		router.Configure()
		router.isConfigured = true
	}
	return router
}

var router *routerUtil

type routerUtil struct {
	isConfigured bool
}

func (r *routerUtil) Configure() {
	app.Server().AddRoute("/api/version", "GET", app.VersionHandler, nil)

	app.Server().AddRoute("/api/users", "POST", user.REST().Create, user.OpenAPI().Create())
	app.Server().AddRoute("/api/users", "GET", user.REST().Get, user.OpenAPI().Get())
	app.Server().AddRoute("/api/users/{id}", "GET", user.REST().GetByID, user.OpenAPI().GetByID())
	app.Server().AddRoute("/api/users/{id}", "PUT", user.REST().UpdateByID, user.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/users/{id}", "PATCH", user.REST().PartiallyUpdateByID, user.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/users/{id}", "DELETE", user.REST().DeleteByID, user.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/categories", "POST", category.REST().Create, category.OpenAPI().Create())
	app.Server().AddRoute("/api/categories", "GET", category.REST().Get, category.OpenAPI().Get())
	app.Server().AddRoute("/api/categories/{id}", "GET", category.REST().GetByID, category.OpenAPI().GetByID())
	app.Server().AddRoute("/api/categories/{id}", "PUT", category.REST().UpdateByID, category.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/categories/{id}", "PATCH", category.REST().PartiallyUpdateByID, category.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/categories/{id}", "DELETE", category.REST().DeleteByID, category.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/categories", "POST", category.REST().Create, category.OpenAPI().Create())
	app.Server().AddRoute("/api/categories", "GET", category.REST().Get, category.OpenAPI().Get())
	app.Server().AddRoute("/api/categories/{id}", "GET", category.REST().GetByID, category.OpenAPI().GetByID())
	app.Server().AddRoute("/api/categories/{id}", "PUT", category.REST().UpdateByID, category.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/categories/{id}", "PATCH", category.REST().PartiallyUpdateByID, category.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/categories/{id}", "DELETE", category.REST().DeleteByID, category.OpenAPI().DeleteByID())

	// AddRoute : DONT REMOVE THIS COMMENT
}
