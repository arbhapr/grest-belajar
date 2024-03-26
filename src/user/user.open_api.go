package user

import "grest-belajar/app"

// OpenAPI is constructor for *openAPI, to autogenerate open api document.
func OpenAPI() *OpenAPIOperation {
	return &OpenAPIOperation{}
}

// OpenAPIOperation embed from app.OpenAPIOperation for simplicity, used for autogenerate open api document.
type OpenAPIOperation struct {
	app.OpenAPIOperation
}

// Base is common detail of users open api document component.
func (o *OpenAPIOperation) Base() {
	o.Tags = []string{"User"}
	o.HeaderParams = []map[string]any{{"$ref": "#/components/parameters/headerParam.Accept-Language"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &User{}}, // will auto create schema $ref: '#/components/schemas/User' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	o.Securities = []map[string][]string{}
}

// Get is detail of `GET /api/v3/users` open api document component.
func (o *OpenAPIOperation) Get() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Get User"
	o.Description = "Use this method to get list of User"
	o.QueryParams = []map[string]any{{"$ref": "#/components/parameters/queryParam.Any"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &UserList{}}, // will auto create schema $ref: '#/components/schemas/User.List' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	return o
}

// GetByID is detail of `GET /api/v3/users/{id}` open api document component.
func (o *OpenAPIOperation) GetByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Get User By ID"
	o.Description = "Use this method to get User by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	return o
}

// Create is detail of `POST /api/v3/users` open api document component.
func (o *OpenAPIOperation) Create() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Create User"
	o.Description = "Use this method to create User"
	o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}

// UpdateByID is detail of `PUT /api/v3/users/{id}` open api document component.
func (o *OpenAPIOperation) UpdateByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Update User By ID"
	o.Description = "Use this method to update User by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	o.Body = map[string]any{"application/json": &ParamUpdate{}}
	return o
}

// PartiallyUpdateByID is detail of `PATCH /api/v3/users/{id}` open api document component.
func (o *OpenAPIOperation) PartiallyUpdateByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Partially Update User By ID"
	o.Description = "Use this method to partially update User by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	o.Body = map[string]any{"application/json": &ParamPartiallyUpdate{}}
	return o
}

// DeleteByID is detail of `DELETE /api/v3/users/{id}` open api document component.
func (o *OpenAPIOperation) DeleteByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Delete User By ID"
	o.Description = "Use this method to delete User by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	o.Body = map[string]any{"application/json": &ParamDelete{}}
	return o
}
