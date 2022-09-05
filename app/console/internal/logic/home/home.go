package home

import (
	"context"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sHome struct {
}

func init() {
	service.RegisterHome(initHome())
}

func initHome() *sHome {
	return &sHome{}
}

// Index home page.
func (s *sHome) Index(ctx context.Context, in *model.HomeIndexInput) (out *model.HomeIndexOutput, err error) {
	return
}
