package controller

import (
	"errors"
	"net/http"
	"strconv"

	"blog.com/config"
	"blog.com/dto"
	"blog.com/model"
	"blog.com/pkg/handle"
	"blog.com/pkg/helper"
	"blog.com/pkg/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct{}

var db = config.GoConnect()

// regeistration
func (con UserController) Register(c *gin.Context) {
	var InputDTO dto.Register

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	var user model.User
	var user2 model.User
	pass_byte, err := bcrypt.GenerateFromPassword([]byte(InputDTO.Password), 10)

	if err != nil {
		response := helper.Error("incription error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if result := db.Where("email=?", InputDTO.Email).Take(&user2); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		} else {
			response := helper.Error("Error", result.Error.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	user.Name = InputDTO.Name
	user.Email = InputDTO.Email
	user.Password = string(pass_byte)

	if result := db.Create(&user); result.Error != nil {
		response := helper.Error("Sql Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", "user created succesfully")
	c.JSON(http.StatusOK, response)

}

// login user
func (con UserController) Login(c *gin.Context) {
	var InputDTO dto.Login
	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	var user model.User
	if result := db.Where("email=?", InputDTO.Email).Take(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response := helper.Error("Error", result.Error.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
			return
		} else {
			response := helper.Error("Error", result.Error.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(InputDTO.Password)); err != nil {
		response := helper.Error("Error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//
	JwtToken := service.GenerateToken(strconv.Itoa(int((user.ID))), 1)

	//user.Password = ""
	user.JwtToken = JwtToken
	//update token
	if _, err := model.UpdateToken(user.ID, user.JwtToken); err != nil {
		msg := helper.Error("SQL error1", err.Error(), helper.EmptyObj{})
		helper.ELog.Error(err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	//
	response := helper.Success(true, "ok", user)
	c.JSON(http.StatusOK, response)
}

// logout..
func (con UserController) Logout(c *gin.Context) {
	user_id := service.GetUserID(c.GetHeader("Token"))

	if result := db.Model(&model.User{}).Where("id=?", user_id).Update("jwt_token", nil); result.Error != nil {
		helper.ELog.Error(result.Error.Error())
		response := helper.Error("Sql Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", "logout successfull")
	c.JSON(http.StatusOK, response)
}
