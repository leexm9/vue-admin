package menu

import (
	"github.com/gin-gonic/gin"
	"go_web/interaction/controller"
)

type menuController struct {
	controller.BaseController
}

var Cont menuController

func (con menuController) GetMenus(c *gin.Context) {
	Cont.Success(c, menuList)
}

type menuItem struct {
	Id       uint64     `json:"id"`
	AuthName string     `json:"authName"`
	Path     string     `json:"path"`
	Children []menuItem `json:"children,omitempty"`
}

var menuList = []menuItem{
	{
		Id:       101,
		AuthName: "用户管理",
		Path:     "users",
		Children: []menuItem{
			{
				Id:       110,
				AuthName: "用户列表",
				Path:     "users",
			},
		},
	},
	{
		Id:       102,
		AuthName: "权限管理",
		Path:     "rights",
		Children: []menuItem{
			{
				Id:       111,
				AuthName: "角色列表",
				Path:     "roles",
			},
			{
				Id:       112,
				AuthName: "权限列表",
				Path:     "rights",
			},
		},
	},
	{
		Id:       103,
		AuthName: "商品管理",
		Path:     "goods",
		Children: []menuItem{
			{
				Id:       113,
				AuthName: "商品列表",
				Path:     "roles",
			},
			{
				Id:       114,
				AuthName: "分类参数",
				Path:     "rights",
			},
			{
				Id:       115,
				AuthName: "商品分类",
				Path:     "rights",
			},
		},
	},
	{
		Id:       104,
		AuthName: "订单管理",
		Path:     "orders",
	},
	{
		Id:       105,
		AuthName: "数据统计",
		Path:     "reports",
	},
}
