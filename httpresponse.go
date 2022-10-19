package npgc

import "github.com/gofiber/fiber/v2"

type RestResponse struct {
	Context    *fiber.Ctx
	HTTPStatus int
	Payload    interface{}
	Status     string
}

func (rr RestResponse) SendResponse() error {
	rr.Context.SendStatus(rr.HTTPStatus)
	rr.Context.JSON(fiber.Map{
		"code":   rr.HTTPStatus,
		"status": rr.Status,
		"data":   rr.Payload,
	})
	return nil
}

func DefaultErrorResponse(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return ctx.JSON(fiber.Map{
		"code":    code,
		"status":  "error",
		"message": message,
	})
}
