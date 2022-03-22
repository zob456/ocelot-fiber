package ocelotFiber
import (
	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http/httptest"
	"testing"
)



func testFiberAuthHandler(c *fiber.Ctx) error {
	return OErrorHandler(c, NotAuthorizedErr, NotAuthorizedCode)
}

func testFiberBadRequestHandler(c *fiber.Ctx) error {
	return OErrorHandler(c, BadRequestErr, BadRequestCode)
}

func testFiberAuthSqlHandler(c *fiber.Ctx) error {
	return OAuthSqlErrorHandler(c, TestSqlErr)
}

func testFiberSqlHandler(c *fiber.Ctx) error {
	return OSqlErrorHandler(c, TestSqlErr)
}

func testFiberExpectedSqlHandler(c *fiber.Ctx) error {
	return OExpectedNoRowsInSqlErrorHandler(c, TestSqlErr)
}

func Test_OcelotFiberErrorHandler_Not_Authorized_HAPPY(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberAuthHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.Equal(t, NotAuthorizedCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Not_Authorized_SAD(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberAuthHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.NotEqual(t, InternalServerErrCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Bad_Request_HAPPY(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberBadRequestHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.Equal(t, BadRequestCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Bad_Request_SAD(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberBadRequestHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.NotEqual(t, InternalServerErrCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Auth_Sql_HAPPY(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberAuthSqlHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.Equal(t, NotAuthorizedCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Auth_Sql_SAD(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberAuthSqlHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.NotEqual(t, InternalServerErrCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Sql_HAPPY(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberSqlHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.Equal(t, NotFoundCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Sql_SAD(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberSqlHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.NotEqual(t, InternalServerErrCode, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Expected_Sql_HAPPY(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberExpectedSqlHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.Equal(t, 200, res.StatusCode)
}

func Test_OcelotFiberErrorHandler_Expected_Sql_SAD(t *testing.T) {
	app := fiber.New()
	app.Post("/", testFiberExpectedSqlHandler)
	req := httptest.NewRequest("POST", "/", nil)
	res, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	assert.NotEqual(t, InternalServerErrCode, res.StatusCode)
}
