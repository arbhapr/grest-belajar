package product

import "grest-belajar/app"

// Product is the main model of Product data. It provides a convenient interface for app.ModelInterface
type Product struct {
	app.Model
	ID           app.NullUUID      `json:"id"                   db:"m.id"              gorm:"column:id;primaryKey"`
	Name         app.NullString    `json:"name"                 db:"m.name"            gorm:"column:name"`
	Stock        app.NullInt64     `json:"stock"                db:"m.stock"           gorm:"column:stock"`
	Price        app.NullFloat64   `json:"price"                db:"m.price"           gorm:"column:price"`
	CategoryID   app.NullUUID      `json:"category.id"          db:"m.category_id"     gorm:"column:category_id"`
	CategoryName app.NullString    `json:"category.name"        db:"c.name"            gorm:"-"`
	CreatedAt    app.NullDateTime  `json:"created_at"           db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt    app.NullDateTime  `json:"updated_at"           db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt    *app.NullDateTime `json:"deleted_at,omitempty" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Product end point, it used for cache key, etc.
func (Product) EndPoint() string {
	return "products"
}

// TableVersion returns the versions of the Product table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Product) TableVersion() string {
	return "24.03.261123"
}

// TableName returns the name of the Product table in the database.
func (Product) TableName() string {
	return "products"
}

// TableAliasName returns the table alias name of the Product table, used for querying.
func (Product) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Product data in the database, used for querying.
func (m *Product) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "categories", "c", []map[string]any{{"column1": "c.id", "column2": "m.category_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Product data in the database, used for querying.
func (m *Product) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	m.AddFilter(map[string]any{"column1": "c.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Product data in the database, used for querying.
func (m *Product) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Product data in the database, used for querying.
func (m *Product) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Product schema, used for querying.
func (m *Product) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Product schema in the open api documentation.
func (Product) OpenAPISchemaName() string {
	return "Product"
}

// GetOpenAPISchema returns the Open API Schema of the Product in the open api documentation.
func (m *Product) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type ProductList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the ProductList schema in the open api documentation.
func (ProductList) OpenAPISchemaName() string {
	return "ProductList"
}

// GetOpenAPISchema returns the Open API Schema of the ProductList in the open api documentation.
func (p *ProductList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Product{})
}

// ParamCreate is the expected parameters for create a new Product data.
type ParamCreate struct {
	UseCaseHandler
	Name       app.NullString  `json:"name"        gorm:"column:name"        validate:"required"`
	Stock      app.NullInt64   `json:"stock"       gorm:"column:stock"       validate:"required"`
	Price      app.NullFloat64 `json:"price"       gorm:"column:price"       validate:"required"`
	CategoryID app.NullUUID    `json:"category_id" gorm:"column:category_id" validate:"required"`
}

// ParamUpdate is the expected parameters for update the Product data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Product data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Product data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
