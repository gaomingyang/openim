package common

import (
	"io/ioutil"
	"net/http"
	"openim/common/define"
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
	c.JSON(http.StatusOK, define.Response{
		Code:    code,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	CommonJSON(c, http.StatusBadRequest, message)
}

func InternalServerError(c *gin.Context, message string) {
	CommonJSON(c, http.StatusInternalServerError, message)
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, define.Response{
		Code:    define.SUCCESSCODE,
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
	body, err := ioutil.ReadAll(resp.Body)
	res = string(body)
	return
}
