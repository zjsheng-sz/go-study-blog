package middleware

import (
	"bytes"
	"encoding/json"
	"go-study-blog/logger"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义响应体记录器
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func DetailedLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 记录请求信息
		logger.Printf("=== 请求开始 ===")
		logger.Printf("方法: %s", c.Request.Method)
		logger.Printf("路径: %s", c.Request.URL.Path)
		logger.Printf("查询参数: %s", c.Request.URL.RawQuery)
		logger.Printf("远程地址: %s", c.Request.RemoteAddr)

		// 记录请求头
		logger.Printf("请求头:")
		for key, values := range c.Request.Header {
			logger.Printf("  %s: %v", key, values)
		}

		// 记录请求体（如果有）
		if c.Request.Body != nil {
			body, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // 重新设置body供后续使用

			if len(body) > 0 {
				logger.Printf("请求体: %s", string(body))

				// 尝试美化JSON输出
				var prettyJSON bytes.Buffer
				if err := json.Indent(&prettyJSON, body, "", "  "); err == nil {
					logger.Printf("格式化请求体:\n%s", prettyJSON.String())
				}
			}
		}

		// 创建响应体记录器
		w := &responseBodyWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		// 处理请求
		c.Next()

		// 记录响应信息
		duration := time.Since(start)
		logger.Printf("=== 响应信息 ===")
		logger.Printf("状态码: %d", c.Writer.Status())
		logger.Printf("处理时间: %v", duration)

		// 记录响应头
		logger.Printf("响应头:")
		for key, values := range c.Writer.Header() {
			logger.Printf("  %s: %v", key, values)
		}

		// 记录响应体
		if w.body.Len() > 0 {
			logger.Printf("响应体: %s", w.body.String())

			// 尝试美化JSON输出
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, w.body.Bytes(), "", "  "); err == nil {
				logger.Printf("格式化响应体:\n%s", prettyJSON.String())
			}
		}

		logger.Printf("=== 请求结束 ===\n")
	}
}
