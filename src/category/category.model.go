package category

import "grest-belajar/app"

// Category is the main model of Category data. It provides a convenient interface for app.ModelInterface
type Category struct {
	app.Model
	ID        app.NullUUID      `json:"id"                   db:"m.id"              gorm:"column:id;primaryKey"`
	Name      app.NullString    `json:"name"                 db:"m.name"            gorm:"column:name"`
	Products  []Products        `json:"products"             db:"category_id={id}"  gorm:"-"`
	CreatedAt app.NullDateTime  `json:"created_at"           db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime  `json:"updated_at"           db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt *app.NullDateTime `json:"deleted_at,omitempty" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Category end point, it used for cache key, etc.
func (Category) EndPoint() string {
	return "categories"
}

// TableVersion returns the versions of the Category table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Category) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Category table in the database.
func (Category) TableName() string {
	return "categories"
}

// TableAliasName returns the table alias name of the Category table, used for querying.
func (Category) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Category data in the database, used for querying.
func (m *Category) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "products", "p", []map[string]any{{"column1": "p.category_id", "column2": "m.id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Category data in the database, used for querying.
func (m *Category) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Category data in the database, used for querying.
func (m *Category) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Category data in the database, used for querying.
func (m *Category) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Category schema, used for querying.
func (m *Category) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Category schema in the open api documentation.
func (Category) OpenAPISchemaName() string {
	return "Category"
}

// GetOpenAPISchema returns the Open API Schema of the Category in the open api documentation.
func (m *Category) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type CategoryList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the CategoryList schema in the open api documentation.
func (CategoryList) OpenAPISchemaName() string {
	return "CategoryList"
}

// GetOpenAPISchema returns the Open API Schema of the CategoryList in the open api documentation.
func (p *CategoryList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Category{})
}

type Products struct {
	app.Model
	ID         app.NullInt64   `json:"id"          db:"p.id"          gorm:"column:id;primaryKey"`
	CategoryID app.NullUUID    `json:"category_id" db:"p.category_id" gorm:"column:category_id"`
	Name       app.NullString  `json:"name"        db:"p.name"        gorm:"column:name"`
	Stock      app.NullInt64   `json:"stock"       db:"p.stock"       gorm:"column:stock"`
	Price      app.NullFloat64 `json:"price"       db:"p.price"       gorm:"column:price"`
}

func (Products) TableName() string {
	return "products"
}

func (Products) TableAliasName() string {
	return "p"
}

func (o *Products) GetRelations() map[string]map[string]any {
	return o.Relations
}

func (m *Products) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "p.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (o *Products) GetFields() map[string]map[string]any {
	o.SetFields(o)
	return o.Fields
}

func (o *Products) GetSchema() map[string]any {
	return o.SetSchema(o)
}

// ParamCreate is the expected parameters for create a new Category data.
type ParamCreate struct {
	UseCaseHandler
	Name app.NullString `json:"name" gorm:"column:name" validate:"required"`
}

// ParamUpdate is the expected parameters for update the Category data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Category data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Category data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
