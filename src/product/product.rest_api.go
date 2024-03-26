package product

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"

	"grest-belajar/app"
)

// REST returns a *RESTAPIHandler.
func REST() *RESTAPIHandler {
	return &RESTAPIHandler{}
}

// RESTAPIHandler provides a convenient interface for Product REST API handler.
type RESTAPIHandler struct {
	UseCase UseCaseHandler
}

// injectDeps inject the dependencies of the Product REST API handler.
func (r *RESTAPIHandler) injectDeps(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx is not found")
	}
	r.UseCase = UseCase(*ctx, app.Query().Parse(c.OriginalURL()))
	return nil
}

// GetByID is the REST API handler for `GET /api/products/{id}`.
func (r *RESTAPIHandler) GetByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// Get is the REST API handler for `GET /api/products`.
func (r *RESTAPIHandler) Get(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.Get()
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res.SetLink(c)
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// Create is the REST API handler for `POST /api/products`.
func (r *RESTAPIHandler) Create(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamCreate{}
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.Create(&p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.Status(http.StatusCreated).JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(p.ID.String)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.Status(http.StatusCreated).JSON(res)
	}
	return c.Status(http.StatusCreated).JSON(grest.NewJSON(res).ToStructured().Data)
}

// UpdateByID is the REST API handler for `PUT /api/products/{id}`.
func (r *RESTAPIHandler) UpdateByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamUpdate{}
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.UpdateByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// PartiallyUpdateByID is the REST API handler for `PATCH /api/products/{id}`.
func (r *RESTAPIHandler) PartiallyUpdateByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamPartiallyUpdate{}
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.PartiallyUpdateByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// DeleteByID is the REST API handler for `DELETE /api/products/{id}`.
func (r *RESTAPIHandler) DeleteByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamDelete{}
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.DeleteByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res := map[string]any{
		"code": http.StatusOK,
		"message": r.UseCase.Ctx.Trans("deleted", map[string]string{
			"products": p.EndPoint(),
			"id":       c.Params("id"),
		}),
	}
	return c.JSON(res)
}
