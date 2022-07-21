package home

import (
	"context"

	"github.com/houseme/url-shortenter/app/api/internal/model"
	"github.com/houseme/url-shortenter/app/api/internal/service"
)

type sHome struct {
}

func init() {
	service.RegisterHome(initHome())
}

func initHome() *sHome {
	return &sHome{}
}

// ShortDetail short url detail
func (s *sHome) ShortDetail(ctx context.Context, in *model.HomeInput) (out string, err error) {
	return
}
