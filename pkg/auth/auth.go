package auth

import (
	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/env"
)

var AuthJWTSecret string
var AuthJWTExpiredHour int

func init() {

	AuthJWTSecret, _ = env.GetEnvString("AUTH_JWT_SECRET")
	AuthJWTExpiredHour, _ = env.GetEnvInt("AUTH_JWT_EXPIRED_HOUR")
}
