package handler

import (
	"github.com/rahmatrdn/go-skeleton/internal/usecase/generate_file"
	"net/http"

	"github.com/rahmatrdn/go-skeleton/internal/parser"
	"github.com/rahmatrdn/go-skeleton/internal/presenter/json"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/generate_file/entity"

	fiber "github.com/gofiber/fiber/v2"
)

type GenerateFileHandler struct {
	parser              parser.Parser
	presenter           json.JsonPresenter
	generateFileUsecase generate_file.IGenerateFileUsecase
}

func NewGenerateFileHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	generateFileUsecase generate_file.IGenerateFileUsecase,
) *GenerateFileHandler {
	return &GenerateFileHandler{parser, presenter, generateFileUsecase}
}

func (w *GenerateFileHandler) Register(app fiber.Router) {
	app.Post("/generate", w.Create)
}

// @Summary			Generate a new file
// @Description		Generate a new file based on request
// @Tags			File Generation
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param			req body entity.GenerateReq true "Payload Request Body"
// @Success			200 {object} entity.GeneralResponse{data=interface{}} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/generate [post]
func (w *GenerateFileHandler) Create(c *fiber.Ctx) error {
	var req entity.GenerateReq

	err := w.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.generateFileUsecase.GenerateFile(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}
