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
func Helper() *UtilHelper {
	return &UtilHelper{}
}

var (
	localInstances = gmap.NewStrAnyMap(true)
	src            = rand.NewSource(time.Now().UnixNano())
)

type UtilHelper struct{}

// UserAgent is a default http userAgent
func (u *UtilHelper) UserAgent(_ context.Context) string {
	return httpHeaderUserAgent
}

// InitTrxID .根据上下文以及账户标识获取交易订单号
func (u *UtilHelper) InitTrxID(ctx context.Context, ano uint64) uint64 {
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
func (u *UtilHelper) InitOrderID(ctx context.Context, datacenterID, workerID int64) int64 {
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
// datacenterID Datacenter ID must be greater than or equal to zero
// workerID Worker ID must be greater than or equal to 0
func (u *UtilHelper) SnowflakeInstance(ctx context.Context, datacenterID, workerID int64) *snowflake.Snowflake {
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
func (u *UtilHelper) AuthToken(ctx context.Context, accountNo uint64) string {
	return gconv.String(u.InitTrxID(ctx, accountNo%32)) + u.InitRandStr(64) + gtime.TimestampNanoStr()
}

// InitRandStr RandStringBytesMaskImprSrcUnsafe
func (u *UtilHelper) InitRandStr(n int) string {
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
func (u *UtilHelper) UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func (u *UtilHelper) LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// GetOutBoundIP 获取本机 iP
func (u *UtilHelper) GetOutBoundIP(ctx context.Context) string {
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
func (u *UtilHelper) GetLocalIpV4(ctx context.Context) string {
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

// GetLocalIPAddresses 获取本地所有的 IP 地址
func (u *UtilHelper) GetLocalIPAddresses(ctx context.Context) (mp map[int][]net.IP, err error) {
	var (
		ifaces []net.Interface
		logger = g.Log(u.Logger(ctx))
	)
	logger.Debug(ctx, "GetLocalIPAddresses start")
	if ifaces, err = net.Interfaces(); err != nil {
		err = gerror.Wrap(err, "GetLocalIPAddresses net.Interfaces failed")
		return
	}
	mp = map[int][]net.IP{
		4: make([]net.IP, 0),
		6: make([]net.IP, 0),
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			logger.Debug(ctx, "GetLocalIPAddresses iface.Flags&net.FlagUp == 0")
			continue // Skip down interfaces
		}

		var addrs []net.Addr
		if addrs, err = iface.Addrs(); err != nil {
			logger.Errorf(ctx, "GetLocalIPAddresses Error getting addresses for interface: %+v", err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip.IsLoopback() {
				logger.Debug(ctx, "GetLocalIPAddresses ip.IsLoopback() ip:", ip.String())
				continue // Skip loopback addresses
			}

			if ip.To4() != nil {
				mp[4] = append(mp[4], ip)
			} else {
				mp[6] = append(mp[6], ip)
			}
		}
	}
	logger.Debug(ctx, "GetLocalIPAddresses end")

	return
}

// Logger .获取上下文中的 logger
func (u *UtilHelper) Logger(ctx context.Context) string {
	return gconv.String(ctx.Value("logger"))
}

// SetLogger .设置上下文中的 logger
func (u *UtilHelper) SetLogger(ctx context.Context, logger string) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

// EncryptSignData sign data
func (u *UtilHelper) EncryptSignData(ctx context.Context, data interface{}, key []byte) ([]byte, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-Helper-EncryptSignData")
	defer span.End()
	byteInfo, err := gjson.Encode(data)
	if err != nil {
		return nil, gerror.Wrap(err, "EncryptSignData gf json.Encode failed")
	}
	return aes.NewAESCrypt(key).Encrypt(byteInfo, gocrypto.ECB)
}

// Header .
func (u *UtilHelper) Header(_ context.Context) map[string]string {
	return g.MapStrStr{
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Accept":          "image/avif,image/webp,image/apng,image/*,*/*;q=0.8",
		"Connection":      "keep-alive",
		"User-Agent":      httpHeaderUserAgent,
	}
}

// HeaderToMap covert request headers to map.
func (u *UtilHelper) HeaderToMap(header http.Header) map[string]interface{} {
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
func (u *UtilHelper) EncryptPass(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// CompareHashAndPassword 校验密码。
func (u *UtilHelper) CompareHashAndPassword(inputPass, authPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(authPass), []byte(inputPass)); err != nil {
		return false
	}
	return true
}

// RequestTime .request time
func (u *UtilHelper) RequestTime(_ context.Context, ts string) *gtime.Time {
	return gtime.NewFromStrFormat(ts, "YmdHis")
}

// ConcatenateSignSource get sign URL 排序并拼接签名的内容信息
func (u *UtilHelper) ConcatenateSignSource(ctx context.Context, data interface{}) string {
	ctx, span := gtrace.NewSpan(ctx, "tracing-enterprise-utility-ConcatenateSignSource")
	defer span.End()

	var (
		tt     = reflect.TypeOf(data)
		v      = reflect.ValueOf(data)
		count  = v.NumField()
		keys   = make([]string, 0, count)
		params = make(map[string]string)
		logger = g.Log(u.Logger(ctx))
	)

	logger.Debug(ctx, "helper ConcatenateSignSource tt", tt, " v1", v)
	for i := 0; i < count; i++ {
		if v.Field(i).CanInterface() { // 判断是否为可导出字段
			logger.Printf(ctx, "%s %s = %v1 -tag:%s", tt.Field(i).Name, tt.Field(i).Type, v.Field(i).Interface(), tt.Field(i).Tag)
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
		buf.WriteString(params[k])
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	result := buf.String()
	logger.Debug(ctx, "helper ConcatenateSignSource string start:", result)
	return result
}

// DecryptSignDataInfo sign data 数据执行 aes 解密
func (u *UtilHelper) DecryptSignDataInfo(src []byte, key []byte) (dst []byte, err error) {
	return aes.NewAESCrypt(key).Decrypt(src, gocrypto.ECB)
}

// HexDecodeString .
func (u *UtilHelper) HexDecodeString(_ context.Context, data string, key []byte) ([]byte, error) {
	if signData, err := hex.DecodeString(data); err != nil {
		return nil, gerror.Wrap(err, "helper HexDecodeString hex.DecodeString failed")
	} else {
		return u.DecryptSignDataInfo(signData, key)
	}
}

// Sha256Of returns the sha256 of the input string
func (u *UtilHelper) Sha256Of(input []byte) string {
	sum := sha256.Sum256(input)
	return hex.EncodeToString(sum[:])
}

// CheckFileExists .
func (u *UtilHelper) CheckFileExists(ctx context.Context, filePath string) (err error) {
	if !gfile.Exists(filePath) {
		if err = gfile.Mkdir(filePath); err != nil {
			err = gerror.Wrap(err, "CheckFileExists gf file.Mkdir failed")
			return
		}
	} else if !gfile.IsDir(filePath) {
		return gerror.NewCode(gcode.CodeInvalidParameter, `parameter "dirPath" should be a directory path`)
	}
	return nil
}

// UserAgentIPHash user agent ip hash
func (u *UtilHelper) UserAgentIPHash(useragent string, ip string) (string, error) {
	var (
		input     = fmt.Sprintf("%s-%s-%s-%d", useragent, ip, time.Now().String(), rand.Int())
		data, err = u.Sha256OfShort(input)
	)
	if err != nil {
		return "", err
	}

	str := u.Base58Encode(data)
	return str[:10], nil
}

// Sha256OfShort returns the sha256 of the input string
func (u *UtilHelper) Sha256OfShort(input string) ([]byte, error) {
	algorithm := sha256.New()
	if _, err := algorithm.Write([]byte(strings.TrimSpace(input))); err != nil {
		return nil, gerror.Wrap(err, "Sha256OfShort write failed")
	}
	return algorithm.Sum(nil), nil
}

// Base58Encode encodes the input byte array to base58 string
func (u *UtilHelper) Base58Encode(data []byte) string {
	return base58.Encode(data)
}

// PasswordBase58Hash password base58 hash
func (u *UtilHelper) PasswordBase58Hash(password string) (string, error) {
	data, err := u.Sha256OfShort(password)
	if err != nil {
		return "", err
	}
	return u.Base58Encode(data), nil
}

// GenerateShortLink generate short link
func (u *UtilHelper) GenerateShortLink(ctx context.Context, url string) (string, error) {
	var (
		logger    = g.Log(u.Logger(ctx))
		data, err = u.Sha256OfShort(url)
	)
	logger.Debug(ctx, "utilHelper GenerateShortLink start url:", url)

	if err != nil {
		return "", err
	}
	logger.Debug(ctx, "utilHelper GenerateShortLink sha256 data:", string(data))
	str := base58.Encode(data)
	logger.Debug(ctx, "utilHelper GenerateShortLink base58 encode str:", str)
	return str[:8], nil
}

// AESEncrypt encrypts the input byte array with the given key
func (u *UtilHelper) AESEncrypt(_ context.Context, key, data []byte) (string, error) {
	return aes.NewAESCrypt(key).EncryptToString(gocrypto.Base64, data, gocrypto.ECB)
}

// AESDecrypt decrypts the input byte array with the given key
func (u *UtilHelper) AESDecrypt(_ context.Context, key, data []byte) (string, error) {
	return aes.NewAESCrypt(key).DecryptToString(gocrypto.Base64, data, gocrypto.ECB)
}

// CreateAccessToken create access token
func (u *UtilHelper) CreateAccessToken(ctx context.Context, accountNo uint64) (token string, err error) {
	var (
		hash      []byte
		initTrxID = u.InitTrxID(ctx, accountNo)
		logger    = g.Log(u.Logger(ctx))
	)
	logger.Debug(ctx, "utilHelper CreateAccessToken accountNo: ", accountNo, " initTrxID: ", initTrxID)
	if hash, err = u.Sha256OfShort(gconv.String(initTrxID)); err != nil {
		err = gerror.Wrap(err, "utilHelper CreateAccessToken Sha256OfShort error")
		return
	}
	token = hex.EncodeToString(hash)
	logger.Debug(ctx, "utilHelper CreateAccessToken token:", token)
	return
}
