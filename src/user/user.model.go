package user

import "grest-belajar/app"

// User is the main model of User data. It provides a convenient interface for app.ModelInterface
type User struct {
	app.Model
	ID        app.NullUUID     `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`
	Email     app.NullString   `json:"email"      db:"m.email"           gorm:"column:email"`
	Name      app.NullString   `json:"name"       db:"m.name"            gorm:"column:name"`
	Password  app.NullString   `json:"password"   db:"m.password"        gorm:"column:password"`
	IsActive  app.NullBool     `json:"is_active"  db:"m.is_active"       gorm:"column:is_active"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the User end point, it used for cache key, etc.
func (User) EndPoint() string {
	return "users"
}

// TableVersion returns the versions of the User table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (User) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the User table in the database.
func (User) TableName() string {
	return "users"
}

// TableAliasName returns the table alias name of the User table, used for querying.
func (User) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the User data in the database, used for querying.
func (m *User) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the User data in the database, used for querying.
func (m *User) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the User data in the database, used for querying.
func (m *User) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the User data in the database, used for querying.
func (m *User) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the User schema, used for querying.
func (m *User) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the User schema in the open api documentation.
func (User) OpenAPISchemaName() string {
	return "User"
}

// GetOpenAPISchema returns the Open API Schema of the User in the open api documentation.
func (m *User) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type UserList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the UserList schema in the open api documentation.
func (UserList) OpenAPISchemaName() string {
	return "UserList"
}

// GetOpenAPISchema returns the Open API Schema of the UserList in the open api documentation.
func (p *UserList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&User{})
}

// ParamCreate is the expected parameters for create a new User data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the User data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the User data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the User data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
