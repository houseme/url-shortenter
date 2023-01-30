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
func (e *utilEnv) Dev(_ context.Context) string {
	return dev
}

// Prod .
func (e *utilEnv) Prod(_ context.Context) string {
	return prod
}

// Test .
func (e *utilEnv) Test(_ context.Context) string {
	return test
}

// Develop .
func (e *utilEnv) Develop(_ context.Context) string {
	return develop
}

// Production .
func (e *utilEnv) Production(_ context.Context) string {
	return production
}
