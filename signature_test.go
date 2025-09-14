package hotcoin

import (
	"testing"
	"time"
)

func TestNewSignature(t *testing.T) {
	secretKey := "test_secret_key"
	sig := NewSignature(secretKey)

	if sig == nil {
		t.Fatal("signature should not be nil")
	}

	if sig.secretKey != secretKey {
		t.Errorf("expected secretKey %s, got %s", secretKey, sig.secretKey)
	}
}

func TestBuildSignString(t *testing.T) {
	sig := NewSignature("test_secret")

	method := "GET"
	host := "api-ct.hotcoin.fit"
	path := "/api/v1/perpetual/private/account"
	params := map[string]string{
		"SignatureMethod":  "HmacSHA256",
		"SignatureVersion": "2",
		"Timestamp":        "2023-01-01T00:00:00.000Z",
		"AccessKeyId":      "test_key",
	}

	signString := sig.buildSignString(method, host, path, params)

	expected := "GET\napi-ct.hotcoin.fit\n/api/v1/perpetual/private/account\nAccessKeyId=test_key&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2023-01-01T00%3A00%3A00.000Z"

	if signString != expected {
		t.Errorf("expected signString:\n%s\ngot:\n%s", expected, signString)
	}
}

func TestSign(t *testing.T) {
	sig := NewSignature("test_secret_key")

	method := "GET"
	host := "api-ct.hotcoin.fit"
	path := "/api/v1/perpetual/private/account"
	params := map[string]string{
		"AccessKeyId": "test_key",
	}

	signature, err := sig.Sign(method, host, path, params)
	if err != nil {
		t.Fatalf("Sign should not return error: %v", err)
	}

	if signature == "" {
		t.Error("signature should not be empty")
	}

	// 验证参数是否被正确添加
	if params["SignatureMethod"] != "HmacSHA256" {
		t.Error("SignatureMethod should be added to params")
	}

	if params["SignatureVersion"] != "2" {
		t.Error("SignatureVersion should be added to params")
	}

	if params["Timestamp"] == "" {
		t.Error("Timestamp should be added to params")
	}

	// 验证时间戳格式
	_, err = time.Parse("2006-01-02T15:04:05.999Z", params["Timestamp"])
	if err != nil {
		t.Errorf("Timestamp format should be valid: %v", err)
	}
}

func TestSignWithNilParams(t *testing.T) {
	sig := NewSignature("test_secret_key")

	method := "GET"
	host := "api-ct.hotcoin.fit"
	path := "/api/v1/perpetual/private/account"

	signature, err := sig.Sign(method, host, path, nil)
	if err != nil {
		t.Fatalf("Sign should not return error with nil params: %v", err)
	}

	if signature == "" {
		t.Error("signature should not be empty")
	}
}

func TestBuildAuthURL(t *testing.T) {
	sig := NewSignature("test_secret_key")

	method := "GET"
	baseURL := "https://api-ct.hotcoin.fit"
	path := "/api/v1/perpetual/private/account"
	apiKey := "test_api_key"
	params := map[string]string{
		"symbol": "BTC-USDT",
	}

	authURL, err := sig.BuildAuthURL(method, baseURL, path, apiKey, params)
	if err != nil {
		t.Fatalf("BuildAuthURL should not return error: %v", err)
	}

	if authURL == "" {
		t.Error("authURL should not be empty")
	}

	// 验证URL包含必要的参数
	if !contains(authURL, "AccessKeyId=test_api_key") {
		t.Error("authURL should contain AccessKeyId")
	}

	if !contains(authURL, "SignatureMethod=HmacSHA256") {
		t.Error("authURL should contain SignatureMethod")
	}

	if !contains(authURL, "SignatureVersion=2") {
		t.Error("authURL should contain SignatureVersion")
	}

	if !contains(authURL, "Timestamp=") {
		t.Error("authURL should contain Timestamp")
	}

	if !contains(authURL, "Signature=") {
		t.Error("authURL should contain Signature")
	}

	if !contains(authURL, "symbol=BTC-USDT") {
		t.Error("authURL should contain original params")
	}
}

func TestBuildAuthURLWithNilParams(t *testing.T) {
	sig := NewSignature("test_secret_key")

	method := "GET"
	baseURL := "https://api-ct.hotcoin.fit"
	path := "/api/v1/perpetual/private/account"
	apiKey := "test_api_key"

	authURL, err := sig.BuildAuthURL(method, baseURL, path, apiKey, nil)
	if err != nil {
		t.Fatalf("BuildAuthURL should not return error with nil params: %v", err)
	}

	if authURL == "" {
		t.Error("authURL should not be empty")
	}
}

// 辅助函数
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr ||
		len(s) > len(substr) && (s[len(s)-len(substr):] == substr ||
			findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
