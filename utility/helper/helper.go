// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package helper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
	"unicode"
	"unsafe"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/houseme/gocrypto"
	"github.com/houseme/gocrypto/aes"
	"github.com/houseme/snowflake"
	"golang.org/x/crypto/bcrypt"

	"github.com/houseme/url-shortenter/utility/env"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	helperUtilSnowflake = "helper.util.snowflake"

	// userAgent .
	httpHeaderUserAgent = `Mozilla/5.0 (url-shorten; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36`
)

// Helper .
func Helper() *utilHelper {
	return &utilHelper{}
}

var (
	localInstances = gmap.NewStrAnyMap(true)
	src            = rand.NewSource(time.Now().UnixNano())
)

type utilHelper struct{}

// UserAgent is a default http userAgent
func (u *utilHelper) UserAgent(_ context.Context) string {
	return httpHeaderUserAgent
}

// InitTrxID .根据上下文以及账户标识获取交易订单号
func (u *utilHelper) InitTrxID(ctx context.Context, ano uint64) uint64 {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-Helper-InitTrxID")
	defer span.End()

	var appEnv, err = env.NewSnowflakeEnv(ctx)
	if err != nil {
		g.Log(u.Logger(ctx)).Error(ctx, "config get fail err:", err)
		return u.InitTrxID(ctx, ano)
	}
	g.Log(u.Logger(ctx)).Debug(ctx, "appEnv DatacenterID:", appEnv.Datacenter(ctx), " WorkerID:", appEnv.Worker(ctx))
	workerID := appEnv.Worker(ctx)
	if ano > 0 {
		workerID = int64(ano % 32)
	}
	return uint64(u.InitOrderID(ctx, appEnv.Datacenter(ctx), workerID))
}

// InitOrderID init64 order id
func (u *utilHelper) InitOrderID(ctx context.Context, datacenterID, workerID int64) int64 {
	g.Log(u.Logger(ctx)).Debug(ctx, "InitOrderID DatacenterID:", datacenterID, " WorkerID:", workerID)
	if datacenterID < 0 || datacenterID > snowflake.GetDatacenterIDMax() {
		g.Log(u.Logger(ctx)).Info(ctx, "InitOrderID datacenter ID error datacenterID", datacenterID)
		return 0
	}

	if workerID < 0 || workerID > snowflake.GetWorkerIDMax() {
		g.Log(u.Logger(ctx)).Debug(ctx, "InitOrderID worker ID error workerID", workerID)
		return 0
	}
	return int64(u.SnowflakeInstance(ctx, datacenterID, workerID).NextVal())
}

// SnowflakeInstance Get Client Instance
// datacenterID Datacenter ID must be greater than or equal to 0
// workerID Worker ID must be greater than or equal to 0
func (u *utilHelper) SnowflakeInstance(ctx context.Context, datacenterID, workerID int64) *snowflake.Snowflake {
	instanceKey := fmt.Sprintf("%s.%02d.%02d", helperUtilSnowflake, datacenterID, workerID)
	g.Log(u.Logger(ctx)).Debug(ctx, "InitOrderID SnowflakeInstance ", instanceKey, workerID, datacenterID)
	return localInstances.GetOrSetFuncLock(instanceKey, func() interface{} {
		s, err := snowflake.NewSnowflake(datacenterID, workerID)
		if err != nil {
			panic(err)
		}
		return s
	}).(*snowflake.Snowflake)
}

// AuthToken user auth token
func (u *utilHelper) AuthToken(ctx context.Context, accountNo uint64) string {
	return gconv.String(u.InitTrxID(ctx, accountNo%32)) + u.InitRandStr(64) + gtime.TimestampNanoStr()
}

// InitRandStr RandStringBytesMaskImprSrcUnsafe
func (u *utilHelper) InitRandStr(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// UcFirst 首字母大些
func (u *utilHelper) UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func (u *utilHelper) LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// GetOutBoundIP 获取本机 iP
func (u *utilHelper) GetOutBoundIP(ctx context.Context) string {
	conn, err := net.Dial("udp", "119.29.29.29:80")
	if err != nil {
		g.Log(u.Logger(ctx)).Error(ctx, " GetOutBoundIP udp get Ip failed err: ", err)
		return ""
	}
	defer func() {
		_ = conn.Close()
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// GetLocalIpV4 获取 IPV4 IP，没有则返回空
func (u *utilHelper) GetLocalIpV4(ctx context.Context) string {
	inters, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, inter := range inters {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			addrs, err := inter.Addrs()
			if err != nil {
				g.Log(u.Logger(ctx)).Error(ctx, " GetLocalIpV4 udp get Ip failed err: ", err)
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					// 判断是否存在 IPV4 IP 如果没有过滤
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

// Logger .获取上下文中的 logger
func (u *utilHelper) Logger(ctx context.Context) string {
	return gconv.String(ctx.Value("logger"))
}

// SetLogger .设置上下文中的 logger
func (u *utilHelper) SetLogger(ctx context.Context, logger string) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

// EncryptSignData sign data
func (u *utilHelper) EncryptSignData(ctx context.Context, data interface{}, key []byte) ([]byte, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-Helper-EncryptSignData")
	defer span.End()
	var (
		logger        = u.Logger(ctx)
		byteInfo, err = gjson.Encode(data)
	)
	g.Log(logger).Debug(ctx, "EncryptSignData data:", data)
	if err != nil {
		err = gerror.Wrap(err, "EncryptSignData gjson.Encode error")
		return byteInfo, err
	}
	return aes.NewAESCrypt(key).Encrypt(byteInfo, gocrypto.ECB)
}

// Header .
func (u *utilHelper) Header(_ context.Context) map[string]string {
	return g.MapStrStr{
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Accept":          "image/avif,image/webp,image/apng,image/*,*/*;q=0.8",
		"Connection":      "keep-alive",
		"User-Agent":      httpHeaderUserAgent,
	}
}

// HeaderToMap covert request headers to map.
func (u *utilHelper) HeaderToMap(header http.Header) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range header {
		if len(v) > 1 {
			m[k] = v
		} else {
			m[k] = v[0]
		}
	}
	return m
}

// EncryptPass .加密处理
func (u *utilHelper) EncryptPass(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// CompareHashAndPassword 校验密码。
func (u *utilHelper) CompareHashAndPassword(inputPass, authPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(authPass), []byte(inputPass)); err != nil {
		return false
	}
	return true
}

// RequestTime .request time
func (u *utilHelper) RequestTime(_ context.Context, ts string) *gtime.Time {
	return gtime.NewFromStrFormat(ts, "YmdHis")
}

// ConcatenateSignSource get sign url 排序并拼接签名的内容信息
func (u *utilHelper) ConcatenateSignSource(ctx context.Context, data interface{}) string {
	ctx, span := gtrace.NewSpan(ctx, "tracing-enterprise-utility-ConcatenateSignSource")
	defer span.End()

	var (
		tt     = reflect.TypeOf(data)
		v      = reflect.ValueOf(data)
		count  = v.NumField()
		keys   = make([]string, 0, count)
		params = make(map[string]string)
		log    = g.Log(u.Logger(ctx))
	)

	log.Debug(ctx, "helper ConcatenateSignSource tt", tt, " v1", v)
	for i := 0; i < count; i++ {
		if v.Field(i).CanInterface() { // 判断是否为可导出字段
			log.Printf(ctx, "%s %s = %v1 -tag:%s", tt.Field(i).Name, tt.Field(i).Type, v.Field(i).Interface(), tt.Field(i).Tag)
			keys = append(keys, u.LcFirst(tt.Field(i).Name))
			params[u.LcFirst(tt.Field(i).Name)] = gconv.String(v.Field(i).Interface())
		}
	}
	// sort params
	sort.Strings(keys)
	var buf bytes.Buffer
	for i := range keys {
		k := keys[i]
		if params[k] == "" || k == "sign" {
			continue
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(string(params[k]))
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	log.Debug(ctx, "helper ConcatenateSignSource string start:", buf.String())
	return buf.String()
}

// DecryptSignDataInfo sign data 数据执行 aes 解密
func (u *utilHelper) DecryptSignDataInfo(src []byte, key []byte) (dst []byte, err error) {
	return aes.NewAESCrypt(key).Decrypt(src, gocrypto.ECB)
}

// HexDecodeString .
func (u *utilHelper) HexDecodeString(ctx context.Context, data string, key []byte) ([]byte, error) {
	if signData, err := hex.DecodeString(data); err != nil {
		err = gerror.Wrap(err, "helper HexDecodeString hex.DecodeString failed")
		return nil, err
	} else {
		return u.DecryptSignDataInfo(signData, key)
	}
}

// Sha256Of returns the sha256 of the input string
func (u *utilHelper) Sha256Of(input []byte) string {
	sum := sha256.Sum256(input)
	return hex.EncodeToString(sum[:])
}

// CheckFileExists .
func (u *utilHelper) CheckFileExists(ctx context.Context, filePath string) (err error) {
	if !gfile.Exists(filePath) {
		if err = gfile.Mkdir(filePath); err != nil {
			g.Log(u.Logger(ctx)).Error(ctx, "CheckFileExists gfile.Mkdir error:", err)
			return err
		}
	} else if !gfile.IsDir(filePath) {
		return gerror.NewCode(gcode.CodeInvalidParameter, `parameter "dirPath" should be a directory path`)
	}
	g.Log(u.Logger(ctx)).Info(ctx, "CheckFileExists filePath:", filePath)
	return nil
}

// UserAgentIPHash user agent ip hash
func (u *utilHelper) UserAgentIPHash(useragent string, ip string) (string, error) {
	var (
		input     = fmt.Sprintf("%s-%s-%s-%d", useragent, ip, time.Now().String(), rand.Int())
		data, err = u.Sha256OfShort(input)
	)
	if err != nil {
		return "", gerror.Wrap(err, "UserAgentIPHash Sha256OfShort failed")
	}

	str := u.Base58Encode(data)
	return str[:10], nil
}

// Sha256OfShort returns the sha256 of the input string
func (u *utilHelper) Sha256OfShort(input string) ([]byte, error) {
	algorithm := sha256.New()
	if _, err := algorithm.Write([]byte(strings.TrimSpace(input))); err != nil {
		return nil, gerror.Wrap(err, "Sha256OfShort write error")
	}
	return algorithm.Sum(nil), nil
}

// Base58Encode encodes the input byte array to base58 string
func (u *utilHelper) Base58Encode(data []byte) string {
	return base58.Encode(data)
}

// PasswordBase58Hash password base58 hash
func (u *utilHelper) PasswordBase58Hash(password string) (string, error) {
	data, err := u.Sha256OfShort(password)
	if err != nil {
		return "", gerror.Wrap(err, "utilHelper PasswordBase58Hash Sha256OfShort error")
	}
	return u.Base58Encode(data), nil
}

// GenerateShortLink generate short link
func (u *utilHelper) GenerateShortLink(ctx context.Context, url string) (string, error) {
	var (
		err     error
		urlHash []byte
		log     = g.Log(u.Logger(ctx))
	)
	log.Debug(ctx, "utilHelper GenerateShortLink url:", url)
	if urlHash, err = u.Sha256OfShort(url); err != nil {
		return "", gerror.Wrap(err, "utilHelper GenerateShortLink Sha256OfShort err")
	}
	// number := new(big.Int).SetBytes(urlHash).Uint64()
	// str := u.Base58Encode(gconv.Bytes(number))
	str := u.Base58Encode(urlHash)
	log.Debug(ctx, "utilHelper GenerateShortLink str:", str)
	return str[:8], nil
}

// AESEncrypt encrypts the input byte array with the given key
func (u *utilHelper) AESEncrypt(_ context.Context, key, data []byte) (string, error) {
	return aes.NewAESCrypt(key).EncryptToString(gocrypto.Base64, data, gocrypto.ECB)
}

// AESDecrypt decrypts the input byte array with the given key
func (u *utilHelper) AESDecrypt(_ context.Context, key, data []byte) (string, error) {
	return aes.NewAESCrypt(key).DecryptToString(gocrypto.Base64, data, gocrypto.ECB)
}

// CreateAccessToken create access token
func (u *utilHelper) CreateAccessToken(ctx context.Context, accountNo uint64) (token string, err error) {
	var (
		hash      []byte
		initTrxID = u.InitTrxID(ctx, accountNo)
		log       = g.Log(u.Logger(ctx))
	)
	log.Debug(ctx, "utilHelper CreateAccessToken accountNo: ", accountNo, " initTrxID: ", initTrxID)
	if hash, err = u.Sha256OfShort(gconv.String(initTrxID)); err != nil {
		err = gerror.Wrap(err, "utilHelper CreateAccessToken Sha256OfShort error")
		return
	}
	token = hex.EncodeToString(hash)
	log.Debug(ctx, "utilHelper CreateAccessToken token:", token)
	return
}
