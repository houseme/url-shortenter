// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package alibaba

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/houseme/url-shortenter/utility/env"
	"github.com/houseme/url-shortenter/utility/helper"
)

// Main .
func Main(ctx context.Context, trxID uint64, fileName string) (string, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-alibaba-Main")
	defer span.End()

	// 请替换成您的 AccessKey ID、AccessKey Secret。
	var (
		logger          = helper.Helper().Logger(ctx)
		client          *green.Client
		alibabaEnv, err = env.NewAlibabaEnv(ctx)
	)
	g.Log(logger).Debug(ctx, "alibabaEnv: ", alibabaEnv.String(ctx))
	if err != nil {
		g.Log(logger).Error(ctx, "alibabaEnv.NewAlibabaEnv error: ", err)
		return "", err
	}
	if client, err = green.NewClientWithAccessKey(alibabaEnv.Region(ctx), alibabaEnv.AccessKeyID(ctx), alibabaEnv.AccessKeySecret(ctx)); err != nil {
		g.Log(logger).Error(ctx, "green.NewClientWithAccessKey error: ", err)
		return "", err
	}

	task1 := map[string]interface{}{"dataId": gconv.String(trxID), "url": fileName}
	// scenes：检测场景，支持指定多个场景。
	content, _ := json.Marshal(
		map[string]interface{}{
			"tasks": task1, "scenes": [...]string{"porn", "terrorism", "ad"}, "bizType": "业务场景",
		},
	)

	request := green.CreateImageSyncScanRequest()
	request.SetContent(content)
	response, err := client.ImageSyncScan(request)
	if err != nil {
		err = gerror.Wrap(err, "client.ImageSyncScan failed")
		return "", err
	}
	if response.GetHttpStatus() != 200 {
		return "", gerror.New("response not success. status:" + strconv.Itoa(response.GetHttpStatus()))
	}
	return response.GetHttpContentString(), nil
}
