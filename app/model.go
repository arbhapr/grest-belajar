package app

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"
)

type ModelInterface interface {
	grest.ModelInterface
}

type Model struct {
	grest.Model
}

type ListModel struct {
	Count       int64 `json:"count"`
	PageContext struct {
		Page      int `json:"page"`
		PerPage   int `json:"per_page"`
		PageCount int `json:"total_pages"`
	} `json:"page_context"`
	Links struct {
		First    string `json:"first"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Last     string `json:"last"`
	} `json:"links"`
	Data []map[string]any `json:"results"`
}

func (list *ListModel) SetData(data []map[string]any, query url.Values) {
	list.Data = data
}

func (list *ListModel) SetLink(c *fiber.Ctx) {
	q := Query().Parse(c.OriginalURL())
	q.Set(grest.QueryLimit, strconv.Itoa(int(list.PageContext.PerPage)))

	path, _, _ := strings.Cut(c.OriginalURL(), "?")

	first := q
	first.Del(grest.QueryPage)
	first.Add(grest.QueryPage, "1")
	firstQS, _ := url.QueryUnescape(first.Encode())
	list.Links.First = c.BaseURL() + path + firstQS

	if list.PageContext.Page > 1 && list.PageContext.PageCount > 1 {
		previous := q
		previous.Set(grest.QueryPage, strconv.Itoa(int(list.PageContext.Page-1)))
		previousQS, _ := url.QueryUnescape(previous.Encode())
		list.Links.Previous = c.BaseURL() + path + previousQS
	}

	if list.PageContext.Page < list.PageContext.PageCount {
		next := q
		next.Set(grest.QueryPage, strconv.Itoa(int(list.PageContext.Page+1)))
		nextQS, _ := url.QueryUnescape(next.Encode())
		list.Links.Next = c.BaseURL() + path + nextQS
	}

	last := q
	last.Set(grest.QueryPage, strconv.Itoa(int(list.PageContext.PageCount)))
	lastQS, _ := url.QueryUnescape(last.Encode())
	list.Links.Last = c.BaseURL() + path + lastQS
}
func (list *ListModel) SetOpenAPISchema(m ModelInterface) map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"count": map[string]any{"type": "integer"},
			"page_context": map[string]any{"type": "object", "properties": map[string]any{
				"page":        map[string]any{"type": "integer"},
				"per_page":    map[string]any{"type": "integer"},
				"total_pages": map[string]any{"type": "integer"},
			}},
			"links": map[string]any{"type": "object", "properties": map[string]any{
				"first":    map[string]any{"type": "string"},
				"previous": map[string]any{"type": "string"},
				"next":     map[string]any{"type": "string"},
				"last":     map[string]any{"type": "string"},
			}},
			"results": map[string]any{
				"type":  "array",
				"items": m.GetOpenAPISchema(),
			},
		},
	}
}

type Setting struct {
	Key   string `gorm:"column:key;primaryKey"`
	Value string `gorm:"column:value"`
}

func (Setting) TableName() string {
	return "settings"
}

func (Setting) KeyField() string {
	return "key"
}

func (Setting) ValueField() string {
	return "value"
}

func (Setting) MigrationKey() string {
	return "table_versions"
}

func (Setting) SeedKey() string {
	return "executed_seeds"
}
