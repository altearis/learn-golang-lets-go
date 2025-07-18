package generate_file

import (
	"context"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/queue"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/generate_file/entity"
	"regexp"
	"time"
)

type GeneratefileUsecase struct {
	queue queue.Queue
}

func NewGenerateUsecase(
	queue queue.Queue,
) *GeneratefileUsecase {
	return &GeneratefileUsecase{queue}
}

type IGenerateFileUsecase interface {
	GenerateFile(ctx context.Context, GenerateReq entity.GenerateReq) (*entity.GenerateResponse, error)
}

func (t *GeneratefileUsecase) GenerateFile(ctx context.Context, generateReq entity.GenerateReq) (*entity.GenerateResponse, error) {

	re := regexp.MustCompile(`[^A-Za-z0-9]+`)

	// Replace matches with underscore
	timeSimplify := re.ReplaceAllString(helper.ConvertToJakartaTime(time.Now()), "_")
	fileNameSimplify := re.ReplaceAllString(generateReq.Filename, "_")

	generateBody := entity.FilePten{
		Filename:  fileNameSimplify + "_" + timeSimplify + ".txt",
		Container: generateReq.Message,
	}
	filePtenGenerator, _ := helper.Serialize(generateBody)
	err := t.queue.Publish(queue.GenerateReport, filePtenGenerator, 1)
	if err != nil {
		return nil, err
	}

	return &entity.GenerateResponse{
		Filename:   generateBody.Filename,
		Message:    generateBody.Container,
		Status:     "OK",
		GenerateAt: helper.ConvertToJakartaTime(time.Now()),
	}, nil
}
