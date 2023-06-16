package common

import (
	"github.com/google/uuid"
)

func MakeUuid() string {
	id := uuid.New()
	return id.String()
}
