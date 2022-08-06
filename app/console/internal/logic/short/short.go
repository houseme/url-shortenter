package short

import (
	"context"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sShort struct {
}

// init
func init() {
	service.RegisterShort(initShort())
}

func initShort() *sShort {
	return &sShort{}
}

// CreateShort is the handler for CreateShort
func (s *sShort) CreateShort(ctx context.Context, in *model.CreateShortInput) (out *model.CreateShortOutput, err error) {
	return
}

// ModifyShort is the handler for ModifyShort
func (s *sShort) ModifyShort(ctx context.Context, in *model.ModifyShortInput) (out *model.ModifyShortOutput, err error) {
	return
}

// QueryShort is the handler for QueryShort
func (s *sShort) QueryShort(ctx context.Context, in *model.QueryShortInput) (out *model.QueryShortOutput, err error) {
	return
}

// QueryStat is the handler for QueryStat
func (s *sShort) QueryStat(ctx context.Context, in *model.QueryStatInput) (out *model.QueryStatOutput, err error) {
	return
}
