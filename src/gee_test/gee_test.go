package gee

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gee/src/gee"
)

// 构造测试用的 engine
func setupRouter() *gee.Engine {
	r := gee.New()
	r.GET("/json", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"message": "hello json",
		})
	})
	r.GET("/", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			_, _ = c.Writer.Write([]byte(k + ": " + v[0] + "\n"))
		}
	})
	return r
}

// 测试 /json 接口
func TestJSONRoute(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/json", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("期望状态码 200，实际 %d", w.Code)
	}

	// 验证 Content-Type
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("期望 Content-Type application/json，实际 %s", ct)
	}

	// 验证 body 内容
	var resp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err == nil {
		t.Fatalf("JSON 解析失败: %v, body=%s", err, w.Body.String())
	}
	if resp["message"] != "hello json" {
		t.Errorf("message 不匹配，实际 %v", resp["message"])
	}
}

// 测试 / 接口（header 回显）
func TestIndexRoute(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Custom", "abc")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("期望 200，实际 %d", w.Code)
	}

	body := w.Body.String()
	if body == "" {
		t.Error("body 不应为空")
	}
	// 你自己根据输出格式断言，比如包含 X-Custom
	if !contains(body, "X-Custom") {
		t.Errorf("期望 body 包含 X-Custom，实际: %s", body)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
