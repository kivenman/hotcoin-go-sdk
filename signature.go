package hotcoin

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

// Signature 签名工具
type Signature struct {
	secretKey string
}

// NewSignature 创建签名工具
func NewSignature(secretKey string) *Signature {
	return &Signature{
		secretKey: secretKey,
	}
}

// Sign 根据HOTCOIN签名算法生成签名
func (s *Signature) Sign(method, host, path string, params map[string]string) (string, error) {
	// 添加必需的签名参数
	now := time.Now().UTC()
	timestamp := now.Format("2006-01-02T15:04:05.999Z")

	if params == nil {
		params = make(map[string]string)
	}

	params["SignatureMethod"] = "HmacSHA256"
	params["SignatureVersion"] = "2"
	params["Timestamp"] = timestamp

	// 构建签名字符串
	signString := s.buildSignString(method, host, path, params)

	// 计算HMAC-SHA256签名
	h := hmac.New(sha256.New, []byte(s.secretKey))
	h.Write([]byte(signString))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature, nil
}

// buildSignString 构建签名字符串
func (s *Signature) buildSignString(method, host, path string, params map[string]string) string {
	// 1. 请求方法 + \n
	signString := strings.ToUpper(method) + "\n"

	// 2. 主机名 + \n
	signString += strings.ToLower(host) + "\n"

	// 3. 路径 + \n
	signString += path + "\n"

	// 4. 排序参数并拼接
	signString += s.buildQueryString(params)

	return signString
}

// buildQueryString 构建查询参数字符串
func (s *Signature) buildQueryString(params map[string]string) string {
	// 按ASCII码顺序排序参数名
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 构建查询字符串，进行URL编码
	var queryParts []string
	for _, key := range keys {
		value := params[key]
		encodedKey := url.QueryEscape(key)
		encodedValue := url.QueryEscape(value)
		queryParts = append(queryParts, fmt.Sprintf("%s=%s", encodedKey, encodedValue))
	}

	return strings.Join(queryParts, "&")
}

// BuildAuthURL 构建带认证参数的URL
func (s *Signature) BuildAuthURL(method, baseURL, path, apiKey string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL + path)
	if err != nil {
		return "", err
	}

	if params == nil {
		params = make(map[string]string)
	}

	// 添加AccessKeyId
	params["AccessKeyId"] = apiKey

	// 生成签名
	signature, err := s.Sign(method, u.Host, u.Path, params)
	if err != nil {
		return "", err
	}

	// 添加签名参数
	params["Signature"] = signature

	// 构建完整URL
	query := s.buildQueryString(params)
	if query != "" {
		if strings.Contains(u.RawQuery, "?") {
			u.RawQuery += "&" + query
		} else {
			u.RawQuery = query
		}
	}

	return u.String(), nil
}
