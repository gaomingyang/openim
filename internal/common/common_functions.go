package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"openim/internal/common/define"
	"strings"

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

func HttpGet(url string) (res string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	res = string(body)
	return
}

func GetMD5(inputStr string) string {
	h := md5.New()
	io.WriteString(h, inputStr)
	md5Str := fmt.Sprintf("%x", h.Sum(nil))
	return md5Str
}
