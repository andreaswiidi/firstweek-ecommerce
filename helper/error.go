package helper

import (
	"net/http"
	"one-week-project-ecommerce/data/response"

	"github.com/gofiber/fiber/v2"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandler(ctx *fiber.Ctx, err interface{}) {
	if notFoundError(ctx, err) {
		return
	}

	// if validationErrors(writer, request, err) {
	// 	return
	// }
	// internalServerError(writer, request, err)
}

func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {

		webResponse := response.Response{
			Code:    http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    exception.Error,
		}

		ctx.Status(fiber.StatusNotFound).JSON(webResponse)
		return true
	} else {
		return false
	}
}
