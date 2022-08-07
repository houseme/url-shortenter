// Package short is a short service.
package short

import (
	"github.com/houseme/url-shortenter/app/schedule/internal/service"
)

type sShort struct {
}

func init() {
	service.RegisterShort(initShort())
}

// initShort create a initShort sShort.
func initShort() *sShort {
	return &sShort{}
}
