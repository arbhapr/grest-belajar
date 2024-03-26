package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"
)

// Error returns a pointer to the errorUtil instance (eu).
// If eu is not initialized, it creates a new errorUtil instance and assigns it to eu.
// It ensures that only one instance of errorUtil is created and reused.
func Error() *errorUtil {
	if eu == nil {
		eu = &errorUtil{}
	}
	return eu
}

// eu is a pointer to an errorUtil instance.
// It is used to store and access the singleton instance of errorUtil.
var eu *errorUtil

// errorUtil represents an error utility.
// It embeds grest.Error, which indicates that errorUtil inherits from grest.Error.
type errorUtil struct {
	grest.Error
}

// New creates a new error instance based on the provided status code, message, and optional details.
// It returns the created error instance.
func (errorUtil) New(statusCode int, message string, detail ...any) error {
	return grest.NewError(statusCode, message, detail...)
}

// StatusCode retrieves the status code from an error.
// It checks if the error is an instance of grest.Error or fiber.Error.
// If the error is of either type, it returns the corresponding status code.
// Otherwise, it returns http.StatusInternalServerError.
func (errorUtil) StatusCode(err error) int {
	e, ok := err.(*grest.Error)
	if ok {
		return e.StatusCode()
	}
	f, ok := err.(*fiber.Error)
	if ok {
		return f.Code
	}
	return http.StatusInternalServerError
}

// Detail retrieves the details from an error.
// It checks if the error is an instance of grest.Error.
// If it is, it returns the body of the error.
// Otherwise, it returns nil.
func (errorUtil) Detail(err error) any {
	e, ok := err.(*grest.Error)
	if ok {
		return e.Body()
	}
	return nil
}

// Trace retrieves the trace information from an error.
// It checks if the error is an instance of grest.Error.
// If it is, it returns the trace information.
// Otherwise, it returns nil.
func (errorUtil) Trace(err error) []map[string]any {
	e, ok := err.(*grest.Error)
	if ok {
		return e.Trace()
	}
	return nil
}

// TraceSimple retrieves simplified trace information from an error.
// It checks if the error is an instance of grest.Error.
// If it is, it returns the simplified trace information.
// Otherwise, it returns nil.
func (errorUtil) TraceSimple(err error) map[string]string {
	e, ok := err.(*grest.Error)
	if ok {
		return e.TraceSimple()
	}
	return nil
}

// Handler handles errors by processing them and returning an appropriate response.
// It retrieves the language from the context (c) and assigns it to lang.
// It checks if the error is an instance of grest.Error.
// If it is not, it sets the error code and message based on the received error.
// If the error status code is not in the 4xx or 5xx range, it sets the code to http.StatusInternalServerError.
// If the error status code is http.StatusInternalServerError, it translates the error message and assigns it to e.Message.
// It returns a JSON response with the error status code and body.
func (errorUtil) Handler(c *fiber.Ctx, err error) error {
	lang := "en"
	ctx, ctxOK := c.Locals("ctx").(*Ctx)
	if ctxOK {
		lang = ctx.Lang
	}
	e, ok := err.(*grest.Error)
	if !ok {
		code := http.StatusInternalServerError
		fiberError, isFiberError := err.(*fiber.Error)
		if isFiberError {
			code = fiberError.Code
		}
		e.Code = code
		e.Message = err.Error()
	}
	if e.StatusCode() < 400 || e.StatusCode() > 599 {
		e.Code = http.StatusInternalServerError
	}
	if e.StatusCode() == http.StatusInternalServerError {
		e.Message = Translator().Trans(lang, "500_internal_error")
		if e.Detail == nil {
			e.Detail = map[string]string{"message": e.Error()}
		}
	}
	return c.Status(e.StatusCode()).JSON(e.Body())
}

// Recover recovers from a panic during Fiber request processing.
// It uses a defer statement to catch and recover from panics.
// Inside the deferred function, there is a placeholder for saving logs and sending alerts.
func (errorUtil) Recover(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// todo: save log & send alert to telegram
		}
	}()
	return c.Next()
}
