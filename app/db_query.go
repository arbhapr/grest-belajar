package app

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strings"

	"gorm.io/gorm"
	"grest.dev/grest"
)

// Query returns a pointer to the queryUtil instance (qu).
// If qu is not initialized, it creates a new queryUtil instance, configures it, and assigns it to qu.
// It ensures that only one instance of queryUtil is created and reused.
func Query() *queryUtil {
	if qu == nil {
		qu = &queryUtil{}
	}
	return qu
}

// qu is a pointer to a queryUtil instance.
// It is used to store and access the singleton instance of queryUtil.
var qu *queryUtil

// queryUtil represents a query utility.
type queryUtil struct{}

// Parse parse url to url.Values.
func (queryUtil) Parse(originalURL string) url.Values {
	query := url.Values{}
	_, qs, _ := strings.Cut(originalURL, "?")
	for qs != "" {
		q := ""
		q, qs, _ = strings.Cut(qs, "&")
		if q == "" || strings.Contains(q, ";") {
			continue
		}

		key, value, _ := strings.Cut(q, "=")
		if k, err := url.QueryUnescape(key); err == nil {
			key = k
		}
		if v, err := url.QueryUnescape(value); err == nil {
			value = v
		}
		query.Set(key, value)
	}
	return query
}

// First get first data from database based on model and query.
func (q queryUtil) First(db *gorm.DB, model ModelInterface, query url.Values) error {
	query.Set(grest.QueryInclude, "all")
	res, err := q.Find(db, model, query)
	if err != nil {
		return Error().New(http.StatusInternalServerError, err.Error())
	}
	if len(res) == 0 {
		return gorm.ErrRecordNotFound
	}
	b, err := json.Marshal(res[0])
	if err != nil {
		return Error().New(http.StatusInternalServerError, err.Error())
	}
	err = json.Unmarshal(b, model)
	if err != nil {
		return Error().New(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Find get paginated data from database based on model and query.
func (queryUtil) Find(db *gorm.DB, model ModelInterface, query url.Values) ([]map[string]any, error) {
	q := &grest.DBQuery{}
	q.DB = db
	q.Schema = model.GetSchema()
	q.Query = query
	return q.Find(q.Schema, query)
}

// PaginationInfo get pagination info from database based on model and query.
func (queryUtil) PaginationInfo(db *gorm.DB, model ModelInterface, query url.Values) (int64, int, int, int, error) {
	var err error
	count, page, perPage, pageCount := int64(0), 0, 0, 0
	if query.Get(grest.QueryDisablePagination) == "true" {
		return count, page, -1, pageCount, err
	}

	q := &grest.DBQuery{}
	q.DB = db
	q.Schema = model.GetSchema()
	q.Query = query
	tx, err := q.Prepare(db, q.Schema, query)
	if err != nil {
		return count, page, perPage, pageCount, err
	}
	err = tx.Count(&count).Error
	if err != nil || query.Get(grest.QueryLimit) == "0" {
		return count, page, perPage, pageCount, err
	}
	page, perPage = q.GetPageLimit()
	pageCount = int(math.Ceil(float64(count) / float64(perPage)))
	return count, page, perPage, pageCount, err
}
