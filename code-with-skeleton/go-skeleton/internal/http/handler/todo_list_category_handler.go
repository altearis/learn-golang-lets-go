package handler

import (
	"net/http"

	"github.com/rahmatrdn/go-skeleton/internal/parser"
	"github.com/rahmatrdn/go-skeleton/internal/presenter/json"
	todo_list_category_usecase "github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category/entity"

	fiber "github.com/gofiber/fiber/v2"
)

type TodoListCategoryHandler struct {
	parser                      parser.Parser
	presenter                   json.JsonPresenter
	todoListCrudCategoryUsecase todo_list_category_usecase.ICrudTodoListCategoryUsecase
}

func NewTodoListCategoryHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	todoListCrudCategoryUsecase todo_list_category_usecase.ICrudTodoListCategoryUsecase,
) *TodoListCategoryHandler {
	return &TodoListCategoryHandler{parser, presenter, todoListCrudCategoryUsecase}
}

func (w *TodoListCategoryHandler) Register(app fiber.Router) {
	app.Get("/todo-category/:id", w.GetByID)
	app.Get("/todo-category", w.GetAll)
	app.Post("/todo-category", w.Create)
	app.Put("/todo-category/:id", w.Update)
	app.Delete("/todo-category/:id", w.Delete)
}

// @Summary         Get Todo List Category by ID
// @Description     Get a Todo List Category by its ID
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List Category"
// @Success         200 {object} entity.GeneralResponse{data=entity.TodoListCatReq} "Success"
// @Failure         401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure         422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure         500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router          /api/v1/todo-category/{id} [get]
// @Router			/api/v1/todo-lists/{id} [get]
func (w *TodoListCategoryHandler) GetByID(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.todoListCrudCategoryUsecase.GetByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Get All Todo List Categories
// @Description     Get all Todo List Categories
// @Tags            Todo List
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Success         200 {object} entity.GeneralResponse{data=[]entity.TodoListResponse} "Success"
// @Failure         401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure         500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router          /api/v1/todo-category [get]
func (w *TodoListCategoryHandler) GetAll(c *fiber.Ctx) error {

	data, err := w.todoListCrudCategoryUsecase.GetAll(c.Context())
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Create a new Todo List Category
// @Description     Create a new Todo List Category
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           req body entity.TodoListCatReq true "Payload Request Body"
// @Success         201 {object} entity.GeneralResponse{data=entity.TodoListCatReq} "Success"
// @Failure         401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure         422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure         500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router          /api/v1/todo-category [post]
// @Router			/api/v1/todo-list [post]
func (w *TodoListCategoryHandler) Create(c *fiber.Ctx) error {
	var req entity.TodoListCatReq

	err := w.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.todoListCrudCategoryUsecase.Create(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Update a Todo List Category by ID
// @Description     Updates an existing Todo List Category with new name, description or status
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List Category to update"
// @Param           req body entity.TodoListCatReq true "Category data including name, description and status"
// @Success         200 {object} entity.GeneralResponse "Successfully updated Todo List Category"
// @Failure         401 {object} entity.CustomErrorResponse "Unauthorized access"
// @Failure         422 {object} entity.CustomErrorResponse "Invalid request format or validation error"
// @Failure         500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router          /api/v1/todo-category/{id} [put]
// @Router          /api/v1/todo-list [put]
func (w *TodoListCategoryHandler) Update(c *fiber.Ctx) error {
	var req entity.TodoListCatReq
	id, err := w.parser.ParserIntIDFromPathParams(c)
	err = w.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}
	req.ID = id

	err = w.todoListCrudCategoryUsecase.UpdateByID(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

// @Summary         Delete Todo List Category by ID
// @Description     Delete an existing Todo List Category by its ID
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List Category"
// @Success         200 {object} entity.GeneralResponse "Success"
// @Failure         401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure         422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure         500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router          /api/v1/todo-category/{id} [delete]
// @Router			/api/v1/todo-lists/{id} [delete]
func (w *TodoListCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.todoListCrudCategoryUsecase.DeleteByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
