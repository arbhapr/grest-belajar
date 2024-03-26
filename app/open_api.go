package app

import "grest.dev/grest"

func OpenAPI() *openAPIUtil {
	if openAPI == nil {
		openAPI = &openAPIUtil{}
	}
	return openAPI
}

var openAPI *openAPIUtil

type openAPIUtil struct {
	grest.OpenAPI
}

func (o *openAPIUtil) Configure() *openAPIUtil {
	o.SetVersion()
	o.Servers = []map[string]any{
		{"description": "Local", "url": "http://localhost:4001"},
	}
	o.Info.Title = "grest-belajar"
	o.Info.Description = "The grest-belajar allows you to perform all the operations that you do with our applications. " +
		"grest-belajar is built using REST principles which ensures predictable URLs, uses standard HTTP response codes, " +
		`authentication, and verbs that makes writing applications easy.

## Query Params

grest-belajar support a common way for pagination, sorting, filtering, searching and other using URL query params on ` + "`GET`" + ` method.

### Pagination

You can use the following query parameters for pagination :

* ` + "`" + `$page` + "`" + `: used to specify the page number to retrieve, default = 1.
* ` + "`" + `$per_page` + "`" + `: used to specify the number of items to retrieve per page, default = 10.
* ` + "`" + `$is_disable_pagination` + "`" + `: used to disable pagination and retrieve all items in one request, default = false.
Example :
` + "`" + `` + "`" + `` + "`" + `
GET /contacts?$page=3&$per_page=10
` + "`" + `` + "`" + `` + "`" + `

### Sorting

You can use the ` + "`" + `$sort` + "`" + ` query parameter for sorting.

* Use the field name according to what you want to sort.
* Use dot notation to sort by the field of the object.
* You can specify multiple fields separated by commas.
* Add ` + "`" + `-` + "`" + ` (minus sign) before the field name to sort in descending order.
* Add ` + "`" + `:i` + "`" + ` after the field name to sort case-insensitively.

This is example if you want to retrieve product data sort by category name, then quantity on hand descending, then case-insensitive name descending :
` + "`" + `` + "`" + `` + "`" + `
GET /products?$sort=category.name,-quantity.on_hand,-name:i
` + "`" + `` + "`" + `` + "`" + `

### Filtering

You can use the field name for filtering the result set based on one or more conditions.

* Use the field name according to what you want to filter.
* Use dot notation to filter by the field of the object.
* Use dot notation with ` + "`" + `*` + "`" + ` to filter by the field on array of the object. (TODO)
* Use dot notation with ` + "`" + `0` + "`" + ` to filter by the field on array of the object and also hide non-matching arrays in the results. (TODO)
* You can use the following operators for filtering :

Operator  | Description               | Example
----------|---------------------------|-----------------------------
none      | Equal to (Exact matches)  | ` + "`" + `/contacts?gender=male` + "`" + `
` + "`" + `$eq` + "`" + `     | Same as above             | ` + "`" + `/contacts?gender.$eq=male` + "`" + `
` + "`" + `$ne` + "`" + `     | Not equal to              | ` + "`" + `/contacts?phone.$ne=null` + "`" + `
` + "`" + `$gt` + "`" + `     | Greater than              | ` + "`" + `/contacts?age.$gt=18` + "`" + `
` + "`" + `$gte` + "`" + `    | Greater than or equal     | ` + "`" + `/contacts?age.$gte=21` + "`" + `
` + "`" + `$lt` + "`" + `     | Less than                 | ` + "`" + `/contacts?age.$lt=17` + "`" + `
` + "`" + `$lte` + "`" + `    | Less than or equal        | ` + "`" + `/contacts?age.$lte=15` + "`" + `
` + "`" + `$like` + "`" + `   | Like                      | ` + "`" + `/contacts?name.$like=john%` + "`" + `
` + "`" + `$nlike` + "`" + `  | Not like                  | ` + "`" + `/contacts?name.$nlike=john%` + "`" + `
` + "`" + `$ilike` + "`" + `  | Case-insensitive Like     | ` + "`" + `/contacts?name.$ilike=john%` + "`" + `
` + "`" + `$nilike` + "`" + ` | Case-insensitive Not Like | ` + "`" + `/contacts?name.$nilike=john%` + "`" + `
` + "`" + `$in` + "`" + `     | In                        | ` + "`" + `/contacts?age.$in=17,21,34` + "`" + `
` + "`" + `$nin` + "`" + `    | Not in                    | ` + "`" + `/contacts?age.$nin=17,21,34` + "`" + `

### Conditional filtering

You can use the ` + "`" + `$or` + "`" + ` query parameter with ` + "`" + `|` + "`" + ` delimiter for conditional filtering.

This is example if you want to filter contact data with condition ` + "`" + `(gender = 'female' or age < 10) and (is_salesman = '1' or is_employee = '1')` + "`" + ` :
` + "`" + `` + "`" + `` + "`" + `
GET /contacts?$or=gender:female|age.$lt:10&$or=is_salesman:true|is_employee:true
` + "`" + `` + "`" + `` + "`" + `

### Searching

You can use the ` + "`" + `$search` + "`" + ` query parameter for searching.

This is example if you want to search contact data with code or name contain character "john" (case-insensitive) :
` + "`" + `` + "`" + `` + "`" + `
GET /contacts?$search=code,name:john
` + "`" + `` + "`" + `` + "`" + `

### Comparing

You can use the ` + "`" + `$field` + "`" + ` key for comparing one field to another field in the same record.

This is example if you want to filter product data with qty_on_order greater than qty_available :
` + "`" + `` + "`" + `` + "`" + `
GET /products?qty_on_order.$gt=$field:qty_available
` + "`" + `` + "`" + `` + "`" + `

### Selection

You can use the ` + "`" + `$select` + "`" + ` query parameter to retrieve specific fields in the response.

* Use the field name according to what you want to retrieve.
* Use dot notation to retrieve the field of the object.
* You can specify multiple fields separated by commas, for example : ` + "`" + `GET /contacts?$select=id,code,name,classification.name` + "`" + `.
* By default, array fields are hidden on get list api for performance reason, you can use ` + "`" + `$nclude` + "`" + ` query parameter to retrieve the specific array field, for example : ` + "`" + `/contacts?$include=families,friends,phones` + "`" + `. Example to retrieve all array field : ` + "`" + `/contacts?$include=all` + "`" + `
* You can use the following operator for aggregation :

Operator | Description | Example                    
---------|-------------|----------------------------
` + "`" + `$count` + "`" + ` | count       | ` + "`" + `/products?$select=$count:id` + "`" + `
` + "`" + `$sum` + "`" + `   | sum         | ` + "`" + `/products?$select=$sum:sold` + "`" + `
` + "`" + `$min` + "`" + `   | minimum     | ` + "`" + `/products?$select=$min:sold` + "`" + `
` + "`" + `$max` + "`" + `   | maximum     | ` + "`" + `/products?$select=$max:sold` + "`" + `
` + "`" + `$avg` + "`" + `   | average     | ` + "`" + `/products?$select=$avg:sold` + "`" + `

### Grouping

You can use the ` + "`" + `$group` + "`" + ` query parameter to grouping.

Example :
` + "`" + `` + "`" + `` + "`" + `
/products?$group=category.id&$select=category.id,$sum:sold&$sort:-$sum:sold
` + "`" + `` + "`" + `` + "`" + `
`

	o.Info.Version = APP_VERSION
	if o.Components == nil {
		o.Components = map[string]any{}
	}
	if o.Components["parameters"] == nil {
		o.Components["parameters"] = map[string]any{}
	}
	param, _ := o.Components["parameters"].(map[string]any)
	param["pathParam.ID"] = map[string]any{
		"in":          "path",
		"name":        "id",
		"description": "An ID of the resources",
		"schema":      map[string]any{"type": "string"},
		"required":    true,
	}
	param["queryParam.Any"] = map[string]any{
		"in":   "query",
		"name": "params",
		"schema": map[string]any{
			"type": "object",
			"additionalProperties": map[string]any{
				"type": "string",
			},
		},
		"explode": true,
	}
	param["headerParam.Accept-Language"] = map[string]any{
		"in":   "header",
		"name": "Accept-Language",
		"schema": map[string]any{
			"type":    "string",
			"default": "en-US",
			"enum":    []string{"en-US", "en", "id-ID", "id"},
		},
	}
	o.Components["parameters"] = param
	o.Components["securitySchemes"] = map[string]any{
		"bearerTokenAuth": map[string]any{
			"type":   "http",
			"scheme": "bearer",
		},
	}
	o.Security = []map[string]any{
		{"bearerTokenAuth": []string{}},
	}
	return o
}

type OpenAPIOperationInterface interface {
	grest.OpenAPIOperationInterface
}

type OpenAPIOperation struct {
	grest.OpenAPIOperation
}

func OpenAPIError() *openAPIError {
	return &openAPIError{}
}

type openAPIError struct {
	StatusCode  int
	Message     string
	SchemaName  string
	Description string
	Headers     map[string]any
	Links       map[string]any
}

func (o *openAPIError) BadRequest() map[string]any {
	o.StatusCode = 400
	o.Message = "The request cannot be performed because of malformed or missing parameters."
	o.SchemaName = "Error.BadRequest"
	o.Description = "A validation exception has occurred."
	return o.Response()
}

func (o *openAPIError) Unauthorized() map[string]any {
	o.StatusCode = 401
	o.Message = "Invalid authentication token."
	o.SchemaName = "Error.Unauthorized"
	o.Description = "Invalid authorization credentials."
	return o.Response()
}

func (o *openAPIError) Forbidden() map[string]any {
	o.StatusCode = 403
	o.Message = "The user does not have permission to access the resource."
	o.SchemaName = "Error.Forbidden"
	o.Description = "User doesn't have permission to access the resource."
	return o.Response()
}

func (o *openAPIError) Response() map[string]any {
	res := map[string]any{
		"content": map[string]any{
			"application/json": o,
		},
	}
	if o.Description != "" {
		res["description"] = o.Description
	}
	if len(o.Headers) > 0 {
		res["headers"] = o.Headers
	}
	if len(o.Links) > 0 {
		res["links"] = o.Links
	}
	return res
}

func (o *openAPIError) OpenAPISchemaName() string {
	return o.SchemaName
}

func (o *openAPIError) GetOpenAPISchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"error": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"code": map[string]any{
						"type":    "integer",
						"format":  "int32",
						"example": o.StatusCode,
					},
					"message": map[string]any{
						"type":    "string",
						"example": o.Message,
					},
				},
			},
		},
	}
}
