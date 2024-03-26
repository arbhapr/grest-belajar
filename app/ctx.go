package app

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"gorm.io/gorm"
	"grest.dev/grest"
)

const CtxKey = "ctx"

type Ctx struct {
	Lang   string // language code
	Action Action // general request info

	IsAsync bool     // for async use, autocommit
	mainTx  *gorm.DB // for normal use, commit & rollback from middleware
}

type Action struct {
	Method   string
	EndPoint string
	DataID   string
}

// TxBegin begins a new transaction using the main database connection.
// It returns an error if there is an issue establishing the connection.
func (c *Ctx) TxBegin() error {
	mainTx, err := DB().Conn("main")
	if err != nil {
		return err
	}
	c.mainTx = mainTx.Begin()
	return nil
}

// TxCommit commits the current transaction if it exists (mainTx is not nil).
// Called in middleware when there is no error (http status code is 2xx).
// It does nothing if there is no active transaction.
func (c *Ctx) TxCommit() {
	if c.mainTx != nil {
		c.mainTx.Commit()
	}

	// reset to nil to use gorm autocommit if use goroutine, etc
	c.mainTx = nil
}

// TxRollback rolls back the current transaction if it exists (mainTx is not nil).
// Called on middleware when there is an error (http status code not 2xx)
// It does nothing if there is no active transaction.
func (c *Ctx) TxRollback() {
	if c.mainTx != nil {
		c.mainTx.Rollback()
	}
	// reset to nil to use gorm autocommit if use goroutine, etc
	c.mainTx = nil
}

// Trans translates a given key using the language specified in the context (c.Lang).
// It supports optional parameters for dynamic translation.
func (c Ctx) Trans(key string, params ...map[string]string) string {
	return Translator().Trans(c.Lang, key, params...)
}

// ValidatePermission validates permission for a given ACL key.
// It returns an error if the permission is not granted.
func (c Ctx) ValidatePermission(aclKey string) error {
	// todo
	return nil
}

// This method validates the parameters based on struct tag.
// It returns an error using the language specified in the context (c.Lang) if the validation fails.
func (c Ctx) ValidateParam(v any) error {
	return Validator().ValidateStruct(v, c.Lang)
}

// This method returns the GORM database connection based on the provided connection name (connName).
// If IS_USE_MOCK_DB is true, it returns the mock database connection.
// If c.IsAsync is false and there is an active transaction (c.mainTx), it returns the transaction connection.
// Otherwise, it returns the main database connection.
func (c Ctx) DB(connName ...string) (*gorm.DB, error) {
	if IS_USE_MOCK_DB {
		return Mock().DB()
	}
	// Control the transaction manually (set begin transaction, commit and rollback on middleware)
	if !c.IsAsync && c.mainTx != nil {
		return c.mainTx, nil
	}
	// Autocommit if use goroutine, etc
	return DB().Conn("main")
}

// This method checks if the given error is a "record not found" error from GORM.
// If it is, it creates a new HTTP error with a "not found" status and a translated error message.
// The translated message includes the entity, key, and value involved in the error.
// It returns the created error or nil if the given error is not a "not found" error.
func (c Ctx) NotFoundError(err error, entity, key, value string) error {
	if err != nil && err == gorm.ErrRecordNotFound {
		return Error().New(http.StatusNotFound, c.Trans("not_found",
			map[string]string{
				"entity": c.Trans(entity),
				"key":    c.Trans(key),
				"value":  value,
			},
		))
	}
	return nil
}

// This method performs a hook operation, which involves sleeping for 2 seconds and performing some data manipulation based on the provided parameters.
// It checks if the old value implements the IsFlat() method and determines whether the data is flat.
// If the data is not flat, it converts the old value to a structured format.
// You can do anything you want with this method, for example to save user activity log, send callback/webhook, etc.
func (c Ctx) Hook(method, reason, id string, old any) {

	// kasih jeda 2 detik untuk memastikan db transaction nya sudah di commit
	time.Sleep(2 * time.Second)

	isFlat := false
	flat, ok := old.(interface{ IsFlat() bool })
	if ok {
		isFlat = flat.IsFlat()
	}

	oldData := old
	if !isFlat {
		oldData = grest.NewJSON(old).ToStructured().Data
	}
	oldJSON, _ := json.MarshalIndent(oldData, "", "  ")
	newJSON := []byte{}

	model := reflect.ValueOf(old)
	if m := model.MethodByName("Async"); m.IsValid() {
		useCase := m.Call([]reflect.Value{reflect.ValueOf(c)})
		if len(useCase) > 0 {
			if u := useCase[0].MethodByName("GetByID"); u.IsValid() {
				val := u.Call([]reflect.Value{reflect.ValueOf(id)})
				if len(val) > 0 {
					new := val[0].Interface()
					if !isFlat {
						new = grest.NewJSON(new).ToStructured().Data
					}
					newJSON, _ = json.MarshalIndent(new, "", "  ")
				}
			}
		}
	}
	_ = oldJSON
	_ = newJSON
}
