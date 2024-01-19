package common

import "github.com/gin-gonic/gin"

// 여러 모듈의 컨트롤러를 구현할 수 있는 인터페이스
type Controller interface {
        RegisterRoutes(router *gin.Engine)
}