/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package model

// HomeIndexInput is the request struct for index.
type HomeIndexInput struct {
	*Base `json:"-" dc:"基础模型"`
}

// HomeIndexOutput is the response struct for index.
type HomeIndexOutput struct {
}
