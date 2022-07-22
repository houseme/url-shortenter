package env

import (
	"context"
)

var uEnv = utilEnv{}

// Env .
func Env() *utilEnv {
	return &uEnv
}

type utilEnv struct {
}

const (
	// env
	dev  = "dev"
	prod = "prod"
	test = "test"
	// environment
	develop    = "develop"
	production = "production"
	testing    = "testing"
)

// Dev .
func (u *utilEnv) Dev(ctx context.Context) string {
	return dev
}

// Prod .
func (u *utilEnv) Prod(ctx context.Context) string {
	return prod
}

// Test .
func (u *utilEnv) Test(ctx context.Context) string {
	return test
}

// Develop .
func (u *utilEnv) Develop(ctx context.Context) string {
	return develop
}

// Production .
func (u *utilEnv) Production(ctx context.Context) string {
	return production
}

// Testing .
func (u *utilEnv) Testing(ctx context.Context) string {
	return testing
}
