// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package hello

import (
	"github.com/houseme/url-shortenter/app/front/api/hello"
)

// ControllerV1 is a demo controller for output "Hello World!", It's used for test purpose.
type ControllerV1 struct{}

// NewV1 creates and returns a new hello controller object.
func NewV1() hello.IHelloV1 {
	return &ControllerV1{}
}
