package controllers

import (
	"encoding/json"
	"gin-rdmg-service/commons"
	"gin-rdmg-service/models"
	"gin-rdmg-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name         string   `json:"name"`
	Avatar       string   `json:"avatar"`
	Introduction string   `json:"introduction"`
	Roles        []string `json:"roles"`
}

// @Summary 用户登录
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param login body models.UserLogin true "用户登录名及密码"
// @Success 200 {object} commons.Result
// @Router /noauth/user/login [post]
func Login(c *gin.Context) {

	var errResult = commons.ResultByCodeAndMessage(200, "Account and password are incorrect.")
	b, _ := c.GetRawData()
	var userLogin models.UserLogin
	// 反序列化
	err := json.Unmarshal(b, &userLogin)
	if err != nil {
		errResult.Message = "username or password error."
		c.JSON(http.StatusOK, errResult)
		return
	}
	//密码查询校验
	selectPwd := models.SelectPwd(userLogin.Username)
	if selectPwd == "" {
		errResult.Message = "用户名或密码错误"
		c.JSON(http.StatusOK, errResult)
		return
	}
	if !utils.CheckPwd(userLogin.Password, selectPwd) {
		c.JSON(http.StatusOK, errResult)
		return
	}
	// 获取用户角色列表
	roleList, roleErr := models.SelectUserRoles(userLogin.Username)
	if roleErr != nil || len(roleList) == 0 {
		errResult.Message = "用户角色异常"
		c.JSON(http.StatusOK, errResult)
		return
	}
	token, tokenError := utils.GenToken(userLogin.Username, roleList)
	if tokenError != nil {
		errResult.Message = "create token error."
		c.JSON(http.StatusOK, errResult)
		return
	}
	tokenJson := gin.H{
		"token": "Bearer " + token,
	}
	var result = commons.ResultSuccess(tokenJson)
	c.JSON(http.StatusOK, result)

}

// @Summary 用户详情
// @Tags 用户
// @Produce application/json
// @Success 200 {object} commons.Result
// @Router /auth/user/info [get]
// @Security ApiKeyAuth
func GetUserInfo(c *gin.Context) {
	username, _ := c.Get("username")
	rolesValue, _ := c.Get("roleList")
	roles := rolesValue.([]string)
	var result commons.Result
	var userInfo UserInfo
	userInfo.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	userInfo.Introduction = ""
	userInfo.Name = utils.Strval(username)
	userInfo.Roles = roles
	result.Code = 200
	result.Data = userInfo
	c.JSON(http.StatusOK, result)
}

// @Summary 用户登出
// @Tags 用户
// @Produce application/json
// @Success 200 {object} commons.Result
// @Router /auth/user/logout [post]
// @Security ApiKeyAuth
func Logout(c *gin.Context) {
	var result commons.Result
	c.Set("username", nil)
	result.Code = 200
	result.Data = "success"
	c.JSON(http.StatusOK, result)
}
