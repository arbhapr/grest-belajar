package product

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
	app.DB().RegisterTable("main", Product{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Product{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"products.detail",
		"products.list",
		"products.create",
		"products.edit",
		"products.delete",
	}))
	app.Server().AddRoute("/products", "POST", REST().Create, nil)
	app.Server().AddRoute("/products", "GET", REST().Get, nil)
	app.Server().AddRoute("/products/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/products/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/products/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/products/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestProductID returns an available Product ID.
func getTestProductID() string {
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
		description:  "Get empty list of Product",
		method:       "GET",
		path:         "/products",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create Product with minimum payload",
		method:       "POST",
		path:         "/products",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get Product by ID",
		method:       "GET",
		path:         "/products/" + getTestProductID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update Product by ID",
		method:       "PUT",
		path:         "/products/" + getTestProductID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update Product by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update Product by ID",
		method:       "PATCH",
		path:         "/products/" + getTestProductID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update Product by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete Product by ID",
		method:       "DELETE",
		path:         "/products/" + getTestProductID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete Product by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestProductREST tests the REST API of Product data with specified scenario.
func TestProductREST(t *testing.T) {
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

// BenchmarkProductREST tests the REST API of Product data with specified scenario.
func BenchmarkProductREST(b *testing.B) {
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
