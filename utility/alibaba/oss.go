package alibaba

import (
	"context"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/utility"
	"github.com/houseme/url-shortenter/utility/env"
)

// Upload file upload oss
func Upload(ctx context.Context, fileName, basePath string) error {
	var (
		logger          = utility.Helper().Logger(ctx)
		alibabaEnv, err = env.NewAlibabaEnv(ctx)
	)
	g.Log(logger).Debug(ctx, "alibabaEnv: ", alibabaEnv.String(ctx))
	if err != nil {
		g.Log(logger).Error(ctx, "alibabaEnv.NewAlibabaEnv error: ", err)
		return err
	}

	g.Log(logger).Info(ctx, "Upload file to oss fileName:"+fileName+" basePath:"+basePath)
	// 创建OSSClient实例。
	client, err := oss.New(alibabaEnv.Endpoint(ctx), alibabaEnv.AccessKeyID(ctx), alibabaEnv.AccessKeySecret(ctx))
	if err != nil {
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(alibabaEnv.BucketName(ctx))
	if err != nil {
		return err
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(basePath+fileName, fileName)
	if err != nil {
		return err
	}
	return nil
}
