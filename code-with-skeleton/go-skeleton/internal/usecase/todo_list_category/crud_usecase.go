package todo_list_usecase

import (
	"context"
	"fmt"
	"time"

	errwrap "github.com/pkg/errors"
	generalEntity "github.com/rahmatrdn/go-skeleton/entity"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/repository/mysql"
	mentity "github.com/rahmatrdn/go-skeleton/internal/repository/mysql/entity"
	"github.com/rahmatrdn/go-skeleton/internal/usecase"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category/entity"
)

type CrudTodoListCategoryUsecase struct {
	todoListCatRepo mysql.ITodoListCategoryRepository
}

func NewCrudTodoListCategoryUsecase(
	todoListRepo mysql.ITodoListCategoryRepository,
) *CrudTodoListCategoryUsecase {
	return &CrudTodoListCategoryUsecase{todoListRepo}
}

type ICrudTodoListCategoryUsecase interface {
	GetByID(ctx context.Context, todoListID int64) (*entity.TodoListCatResponse, error)
	GetAll(ctx context.Context) ([]*entity.TodoListCatResponse, error)
	Create(ctx context.Context, todoListReq entity.TodoListCatReq) (*entity.TodoListCatResponse, error)
	UpdateByID(ctx context.Context, todoListReq entity.TodoListCatReq) error
	DeleteByID(ctx context.Context, todoListID int64) error
}

func (t *CrudTodoListCategoryUsecase) GetByID(ctx context.Context, todoListID int64) (*entity.TodoListCatResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.GetByID"
	captureFieldError := generalEntity.CaptureFields{
		"category_id": helper.ToString(todoListID),
	}

	data, err := t.todoListCatRepo.GetByID(ctx, todoListID)
	if err != nil {
		helper.LogError("todoListCatRepo.GetByID", funcName, err, captureFieldError, "")

		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.TodoListCatResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
	}, nil
}
func (t *CrudTodoListCategoryUsecase) GetAll(ctx context.Context) (res []*entity.TodoListCatResponse, err error) {
	funcName := "CrudTodoListCategoryUsecase.GetAll"
	data, err := t.todoListCatRepo.GetAll(ctx)
	if err != nil {
		helper.LogError("todoListCatRepo.GetAll", funcName, err, nil, "")
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	for _, v := range data {
		res = append(res, &entity.TodoListCatResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   helper.ConvertToJakartaTime(v.CreatedAt),
		})
	}

	return res, nil
}

func (t *CrudTodoListCategoryUsecase) Create(ctx context.Context, todoListReq entity.TodoListCatReq) (*entity.TodoListCatResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.Create"
	captureFieldError := generalEntity.CaptureFields{
		"category_id": helper.ToString(todoListReq.ID),
		"payload":     helper.ToString(todoListReq),
	}

	if errMsg := usecase.ValidateStruct(todoListReq); errMsg != "" {
		return nil, errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	todoListPayload := &mentity.TodoListCategoryEnt{
		Name:        todoListReq.Name,
		Description: todoListReq.Description,
		CreatedAt:   time.Now(),
	}

	err := t.todoListCatRepo.Create(ctx, nil, todoListPayload, false)
	if err != nil {
		helper.LogError("todoListCatRepo.Create", funcName, err, captureFieldError, "")
		return nil, err
	}

	return &entity.TodoListCatResponse{
		ID:          todoListPayload.ID,
		Name:        todoListPayload.Name,
		Description: todoListPayload.Description,
		CreatedAt:   helper.ConvertToJakartaTime(todoListPayload.CreatedAt),
	}, nil
}

func (t *CrudTodoListCategoryUsecase) UpdateByID(ctx context.Context, todoListCatReq entity.TodoListCatReq) error {
	funcName := "CrudTodoListCategoryUsecase.UpdateByID"
	todoListID := todoListCatReq.ID

	captureFieldError := generalEntity.CaptureFields{
		"category_id": helper.ToString(todoListCatReq.ID),
		"payload":     helper.ToString(todoListCatReq),
	}

	if err := mysql.DBTransaction(t.todoListCatRepo, func(trx mysql.TrxObj) error {
		lockedData, err := t.todoListCatRepo.LockByID(ctx, trx, todoListID)
		if err != nil {
			helper.LogError("todoListCatRepo.LockByID", funcName, err, captureFieldError, "")

			return err
		}
		if lockedData == nil {
			return fmt.Errorf("DATA IS NOT EXIST")
		}

		if err := t.todoListCatRepo.Update(ctx, trx, lockedData, &mentity.TodoListCategoryEnt{
			Name:        todoListCatReq.Name,
			Description: todoListCatReq.Description,
		}); err != nil {
			helper.LogError("todoListCatRepo.Update", funcName, err, captureFieldError, "")

			return err
		}

		return nil
	}); err != nil {
		helper.LogError("todoListCatRepo.DBTransaction", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}

func (t *CrudTodoListCategoryUsecase) DeleteByID(ctx context.Context, todoListCatID int64) error {
	funcName := "CrudTodoListCategoryUsecase.DeleteByID"
	captureFieldError := generalEntity.CaptureFields{
		"todo_list_category_id": helper.ToString(todoListCatID),
	}

	err := t.todoListCatRepo.DeleteByID(ctx, nil, todoListCatID)
	if err != nil {
		helper.LogError("todoListCatRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
