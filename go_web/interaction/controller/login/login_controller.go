package login

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_web/interaction/controller"
	"go_web/middleware/log"
	"go_web/pkg/req"
	"go_web/pkg/resp"
)

type loginController struct {
	controller.BaseController
}

var Cont loginController
var logger = log.GetInstance()

func (con *loginController) Login(c *gin.Context) {
	var login req.Login
	if err := c.BindJSON(&login); err != nil {
		con.Error(c, resp.ParamMissErr)
	}
	logger.Info("receiver message:", zap.String("username", login.Username))
	rest := make(map[string]string)
	rest["token"] = "iADYAODt4DTD34DYDiEC8cBdqQDCxDFDljcnqDEDYPDxitDbq"
	con.Success(c, rest)
}
