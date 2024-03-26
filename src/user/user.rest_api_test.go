package user

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"

	"grest-belajar/app"
)

// prepareTest prepares the test.
func prepareTest(tb testing.TB) {
	app.Test()
	tx := app.Test().Tx
	app.DB().RegisterTable("main", User{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"users.detail",
		"users.list",
		"users.create",
		"users.edit",
		"users.delete",
	}))
	app.Server().AddRoute("/users", "POST", REST().Create, nil)
	app.Server().AddRoute("/users", "GET", REST().Get, nil)
	app.Server().AddRoute("/users/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/users/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/users/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/users/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestUserID returns an available User ID.
func getTestUserID() string {
	return "todo"
}

// tests is test scenario.
var tests = []struct {
	description  string // description of the test case
	method       string // method to test
	path         string // route path to test
	token        string // token to test
	bodyRequest  string // body to test
	expectedCode int    // expected HTTP status code
	expectedBody string // expected body response
}{
	{
		description:  "Get empty list of User",
		method:       "GET",
		path:         "/users",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create User with minimum payload",
		method:       "POST",
		path:         "/users",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get User by ID",
		method:       "GET",
		path:         "/users/" + getTestUserID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update User by ID",
		method:       "PUT",
		path:         "/users/" + getTestUserID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update User by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update User by ID",
		method:       "PATCH",
		path:         "/users/" + getTestUserID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update User by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete User by ID",
		method:       "DELETE",
		path:         "/users/" + getTestUserID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete User by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestUserREST tests the REST API of User data with specified scenario.
func TestUserREST(t *testing.T) {
	prepareTest(t)

	// Iterate through test single test cases
	for _, test := range tests {

		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
		req.Header.Add("Authorization", "Bearer "+test.token)
		req.Header.Add("Content-Type", "application/json")

		// Perform the request plain with the app, the second argument is a request latency (set to -1 for no latency)
		res, err := app.Server().Test(req)

		// Verify if the status code is as expected
		utils.AssertEqual(t, nil, err, "app.Server().Test(req)")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)

		// Verify if the body response is as expected
		body, err := io.ReadAll(res.Body)
		utils.AssertEqual(t, nil, err, "io.ReadAll(res.Body)")
		app.Test().AssertMatchJSONElement(t, []byte(test.expectedBody), body, test.description)
		res.Body.Close()
	}
}

// BenchmarkUserREST tests the REST API of User data with specified scenario.
func BenchmarkUserREST(b *testing.B) {
	b.ReportAllocs()
	prepareTest(b)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
			req.Header.Add("Authorization", "Bearer "+test.token)
			req.Header.Add("Content-Type", "application/json")
			app.Server().Test(req)
		}
	}
}
