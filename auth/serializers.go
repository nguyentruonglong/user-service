package auth

import (
	"github.com/gin-gonic/gin"
)

type JWTGenerateSerializer struct {
	C *gin.Context
	JWTClaim
}
