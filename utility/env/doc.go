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
)

// Dev .
func (u *utilEnv) Dev(_ context.Context) string {
	return dev
}

// Prod .
func (u *utilEnv) Prod(_ context.Context) string {
	return prod
}

// Test .
func (u *utilEnv) Test(_ context.Context) string {
	return test
}

// Develop .
func (u *utilEnv) Develop(_ context.Context) string {
	return develop
}

// Production .
func (u *utilEnv) Production(_ context.Context) string {
	return production
}
