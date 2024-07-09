// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package env is a utility package for env.
package env

import (
	"context"
)

// Env .
func Env() *UtilEnv {
	return &UtilEnv{}
}

// UtilEnv util env
type UtilEnv struct {
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
func (e *UtilEnv) Dev(_ context.Context) string {
	return dev
}

// Prod .
func (e *UtilEnv) Prod(_ context.Context) string {
	return prod
}

// Test .
func (e *UtilEnv) Test(_ context.Context) string {
	return test
}

// Develop .
func (e *UtilEnv) Develop(_ context.Context) string {
	return develop
}

// Production .
func (e *UtilEnv) Production(_ context.Context) string {
	return production
}
