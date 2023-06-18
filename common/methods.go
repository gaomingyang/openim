package common

import (
	"github.com/google/uuid"
	"strings"
)

func MakeUuid() string {
	id := uuid.New()
	s := id.String()
	s = strings.Replace(s, "-", "", 0)
	return s
}
