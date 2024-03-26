package middleware

import (
	"github.com/gofiber/fiber/v2"

	"grest-belajar/app"
)

func Ctx() *ctxHandler {
	if ch == nil {
		ch = &ctxHandler{}
	}
	return ch
}

var ch *ctxHandler

type ctxHandler struct{}

func (*ctxHandler) New(c *fiber.Ctx) error {
	lang := c.Get("Accept-Language")
	if lang == "" {
		lang = "en"
	}
	ctx := app.Ctx{
		Lang: lang,
	}
	c.Locals("ctx", &ctx)
	return c.Next()
}
