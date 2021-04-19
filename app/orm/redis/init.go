package redis

import (
	"github.com/locpham24/go-coffee/infra"
)

func InitRedisInstances() {
	Token = &jwtToken{infra.Client}
}
