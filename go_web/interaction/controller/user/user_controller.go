package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_web/interaction/controller"
	"go_web/middleware/log"
	"go_web/pkg/resp"
	"strconv"
	"strings"
	"time"
)

type userController struct {
	controller.BaseController
}

var Cont userController
var logger = log.GetInstance()

func (con *userController) GetUsers(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")
	var pageSize int
	if len(pageSizeStr) == 0 {
		pageSize = 10
	}
	pageIndex, err := strconv.ParseInt(pageIndexStr, 10, 64)
	if err != nil {
		con.Error(c, resp.ParamErr)
		return
	}
	if size, err := strconv.ParseUint(pageSizeStr, 10, 64); err != nil {
		con.Error(c, resp.ParamErr)
		return
	} else {
		pageSize = int(size)
	}
	logger.Info("page query:", zap.Int("pageSize", pageSize), zap.Int64("pageIndex", pageIndex))

	var list []User
	if pageSize >= len(users) {
		list = users[:]
	} else {
		list = users[:pageSize]
	}
	pageResult := PageResult{
		Total: int32(len(users)),
		List:  list,
	}

	con.Success(c, pageResult)
}

func (con *userController) ChangeState(c *gin.Context) {
	var state State
	if err := c.BindJSON(&state); err != nil {
		con.Error(c, resp.ParamErr)
		return
	}
	logger.Info("state change", zap.Uint64("id", state.Id), zap.String("mgState", state.MgState))
	for i := 0; i < len(users); i++ {
		if users[i].Id == state.Id {
			if flag, err := strconv.ParseBool(state.MgState); err != nil {
				con.Error(c, resp.ParamErr)
				return
			} else {
				users[i].MgState = flag
			}
			break
		}
	}
	con.Success(c, nil)
}

func (con *userController) AddUser(c *gin.Context) {
	var userDO UserDO
	if err := c.BindJSON(&userDO); err != nil {
		con.Error(c, resp.ParamErr)
		return
	}
	logger.Info("add user", zap.String("username", userDO.Username), zap.String("mobile", userDO.Mobile),
		zap.String("email", userDO.Email))
	user := User{
		Id:        uint64(501 + len(users)),
		Username:  userDO.Username,
		RoleName:  "普通用户",
		Mobile:    userDO.Mobile,
		Email:     userDO.Email,
		Type:      1,
		MgState:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users = append(users, user)
	con.Success(c, nil)
}

func (con *userController) GetById(c *gin.Context) {
	sid := c.Query("id")
	if len(sid) == 0 {
		con.Error(c, resp.ParamMissErr)
		return
	}
	var user User
	for i := 0; i < len(users); i++ {
		if strings.Compare(sid, strconv.FormatUint(users[i].Id, 10)) == 0 {
			user = users[i]
			break
		}
	}
	if user == (User{}) {
		con.Error(c, resp.ResultNotFound)
		return
	}
	con.Success(c, user)
}

func (con *userController) EditUser(c *gin.Context) {
	var editUser EditUser
	if err := c.BindJSON(&editUser); err != nil {
		con.Error(c, resp.ParamErr)
		return
	}
	var flag bool
	for i := 0; i < len(users); i++ {
		if editUser.Id == users[i].Id {
			users[i].Email = editUser.Email
			users[i].Mobile = editUser.Mobile
			flag = true
		}
	}
	con.Success(c, flag)
}

func (con *userController) DeleteUser(c *gin.Context) {
	sid := c.Query("id")
	if len(sid) == 0 {
		con.Error(c, resp.ParamMissErr)
		return
	}
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		con.Error(c, resp.ParamMissErr)
		return
	}
	k := -1
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			k = i
			break
		}
	}
	if k == -1 {
		con.Error(c, resp.ResultNotFound)
		return
	}
	users = append(users[:k], users[k+1:]...)
	con.Success(c, true)
}

var users = []User{
	{
		Id:        500,
		Username:  "admin",
		RoleName:  "超级管理员",
		Mobile:    "13945628276",
		Email:     "admin123@ce.com",
		Type:      1,
		MgState:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        502,
		Username:  "linken",
		RoleName:  "测试角色",
		Mobile:    "13545628270",
		Email:     "ceshi123@ce.com",
		Type:      1,
		MgState:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type PageResult struct {
	List  []User `json:"list"`
	Total int32  `json:"total"`
}

type User struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	RoleName  string    `json:"roleName"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	Type      int8      `json:"type"`
	MgState   bool      `json:"mgState"` // 用户的状态
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type State struct {
	Id      uint64 `json:"id"`
	MgState string `json:"mgState"`
}

type UserDO struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

type EditUser struct {
	Id     uint64 `json:"id"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}
