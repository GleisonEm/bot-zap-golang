package middleware

import (
	"fmt"
	pkgError "github.com/GleisonEm/bot-claudio-zap-golang/pkg/error"
	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Recovery() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer func() {
			err := recover()
			if err != nil {
				var res utils.ResponseData
				res.Status = 500
				res.Code = "INTERNAL_SERVER_ERROR"
				res.Message = fmt.Sprintf("%v", err)

				errValidation, isValidationError := err.(pkgError.GenericError)
				if isValidationError {
					res.Status = errValidation.StatusCode()
					res.Code = errValidation.ErrCode()
					res.Message = errValidation.Error()
				}

				_ = ctx.Status(res.Status).JSON(res)
			}
		}()

		return ctx.Next()
	}
}
