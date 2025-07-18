package consumer

import (
	"context"

	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/generate_file/entity"
)

type ReportQueue struct {
	ctx context.Context
}

type GenerateReport interface {
	Process(payload map[string]interface{}) error
}

func NewGenerateReport(
	ctx context.Context,
) GenerateReport {
	return &ReportQueue{ctx}
}

func (l *ReportQueue) Process(payload map[string]interface{}) error {
	var params entity.FilePten
	params.LoadFromMap(payload)

	helper.Dump(params)

	return nil
}
