package rest

import (
	"github.com/gin-gonic/gin"
	"nafue/config"
	"net/http"
)

type LetsEncryptController struct{}

func (self *LetsEncryptController) Init(r *gin.Engine) {
	r.GET(config.LetsEncryptPath, self.letsEncrypt)
}

func (self *LetsEncryptController) letsEncrypt(c *gin.Context) {
	c.String(http.StatusOK, config.LetsEncryptContent)
	return
}
