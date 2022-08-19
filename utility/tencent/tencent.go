package tencent

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ims "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ims/v20201229"

	"github.com/houseme/url-shortenter/utility"
	"github.com/houseme/url-shortenter/utility/env"
)

// Main .
func Main(ctx context.Context, trxID uint64, fileName string) (string, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-alibaba-Main")
	defer span.End()

	// 请替换成您的AccessKey ID、AccessKey Secret。
	var (
		cpf             = profile.NewClientProfile()
		logger          = utility.Helper().Logger(ctx)
		tencentEnv, err = env.NewTencentEnv(ctx)
	)
	if err != nil {
		g.Log(logger).Error(ctx, "tencentEnv.NewTencentEnv error: ", err)
		return "", err
	}
	g.Log(logger).Debug(ctx, "tencentEnv: ", tencentEnv.String(ctx))

	credential := common.NewCredential(tencentEnv.SecretID(ctx), tencentEnv.SecretKey(ctx))
	cpf.HttpProfile.Endpoint = tencentEnv.Endpoint(ctx)
	var (
		client   *ims.Client
		response *ims.ImageModerationResponse
	)
	if client, err = ims.NewClient(credential, tencentEnv.Region(ctx), cpf); err != nil {
		err = gerror.Wrap(err, "ims.NewClient error")
		return "", err
	}

	request := ims.NewImageModerationRequest()
	request.DataId = common.StringPtr(gconv.String(trxID))
	request.FileUrl = common.StringPtr(fileName)
	response, err = client.ImageModeration(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		g.Log(logger).Error(ctx, "ims.ImageModeration error: ", err)
	}
	if err != nil {
		err = gerror.Wrap(err, "ims.ImageModeration error")
		return "", err
	}
	g.Log(logger).Info(ctx, "response: ", response.ToJsonString())
	return response.ToJsonString(), nil
}
