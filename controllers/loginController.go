package controllers

import (
	"github.com/1pedrohfreitas/pcams_back_go/database"
	"github.com/1pedrohfreitas/pcams_back_go/models"
	"github.com/1pedrohfreitas/pcams_back_go/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDataBase()

	var p models.Login
	var login models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Erro no Json",
		})
		return
	}
	println(p.Username)
	dbErro := db.QueryRow(`select id ,username, password from users where status = 1 and username=$1`,
		p.Username).Scan(&login.ID, &login.Username, &login.Password)

	if dbErro != nil {
		c.JSON(401, gin.H{
			"error": "usuario não encontrado",
		})
		return
	}

	if login.Password != services.SHA256Encoder(p.Password) {
		c.JSON(401, gin.H{
			"error": "Senha invalida",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(login.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
		"id":    login.ID,
	})
}
func ValidateToken(c *gin.Context) {

	var p models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Erro no Json",
		})
		return
	}

	isValidToken := services.NewJWTService().ValidateToken(p.Token)
	if !isValidToken {
		c.JSON(401, gin.H{
			"error": "Token expirado",
		})
		return
	}
	c.JSON(200, gin.H{
		"error": "Token valido",
	})
}
