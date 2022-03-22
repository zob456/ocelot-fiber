package ocelotFiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zob456/ocelot-logger"
)

// Error handler funcs

func OErrorHandler(c *fiber.Ctx, err error, errorCode int) error {
	ocelotLogger.ErrorLogger(err)
	return c.SendStatus(errorCode)
}

func OAuthSqlErrorHandler(c *fiber.Ctx, err error) error{
	if err.Error() == SqlErr {
		ocelotLogger.ErrorLogger(err)
		c.Response().SetBodyString(err.Error())
		return c.SendStatus(401)
	}
	ocelotLogger.ErrorLogger(err)
	return c.SendStatus(500)
}

func OSqlErrorHandler(c *fiber.Ctx, err error) error{
		if err.Error() == SqlErr {
			ocelotLogger.ErrorLogger(err)
			c.Response().SetBodyString(err.Error())
			return c.SendStatus(404)
		}
		ocelotLogger.ErrorLogger(err)
		return c.SendStatus(500)
}

func OExpectedNoRowsInSqlErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() != SqlErr {
		ocelotLogger.ErrorLogger(err)
		c.Response().SetBodyString(err.Error())
		return c.SendStatus(500)
	}
	return nil
}