package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahmatrdn/go-skeleton/internal/parser"
	"github.com/rahmatrdn/go-skeleton/internal/presenter/json"
	"net/http"
)

type HelloWorld struct {
	parser    parser.Parser
	presenter json.JsonPresenter
}

func HelloWorldHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
) *HelloWorld {
	return &HelloWorld{parser, presenter}
}

func (w *HelloWorld) Register(app fiber.Router) {
	app.Get("/hello", w.Hello)
}

func (w *HelloWorld) Hello(c *fiber.Ctx) error {
	msg := "Hello, from Golang!"
	return w.presenter.BuildSuccess(c, msg, "Success", http.StatusOK)
}
