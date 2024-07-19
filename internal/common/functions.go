package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"openim/internal/common/define"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeUuid() string {
	id := uuid.New()
	s := id.String()
	s = strings.ReplaceAll(s, "-", "")
	return s
}

func CommonJSON(c *gin.Context, code int, message string) {
	traceID := c.MustGet("trace_id").(string)
	c.JSON(http.StatusOK, define.Response{
		Code:    code,
		TraceID: traceID,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	CommonJSON(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	CommonJSON(c, http.StatusUnauthorized, message)
}

func InternalServerError(c *gin.Context, message string) {
	CommonJSON(c, http.StatusInternalServerError, message)
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, define.Response{
		Code:    define.SUCCESSCODE,
		TraceID: c.MustGet("trace_id").(string),
		Message: "success",
		Data:    data,
	})
}

func HttpGet(url string) (res []byte, err error) {
	client := &http.Client{Timeout: time.Second * 10} // 设置10秒超时
	resp, err := client.Get(url)
	if err != nil {
		return res, fmt.Errorf("HTTP GET request failed: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("failed to read response body: %v", err)
	}
	return body, nil
}

func HttpGetWithHeader(url string, header map[string]string) (res []byte, err error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set custom headers
	for k, h := range header {
		//log.Println("set header:", k, ":", h)
		req.Header.Set(k, h)
	}

	// Set cookies
	//req.AddCookie(&http.Cookie{Name: "cookie_name", Value: "cookie_value"})

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	//log.Println("response:", string(body))

	return body, err
}

func GetMD5(inputStr string) string {
	h := md5.New()
	io.WriteString(h, inputStr)
	md5Str := fmt.Sprintf("%x", h.Sum(nil))
	return md5Str
}
