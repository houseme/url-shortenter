// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package home

import (
	"github.com/houseme/url-shortenter/app/front/api/home"
)

// ControllerV1  is the controller for the home API version 1.
type ControllerV1 struct{}

// NewV1 creates and returns a new home controller.
func NewV1() home.IHomeV1 {
	return &ControllerV1{}
}
