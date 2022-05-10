package controllers

import (
	"fmt"
	"strconv"

	"github.com/1pedrohfreitas/pcams_back_go/database"
	"github.com/1pedrohfreitas/pcams_back_go/dto"
	"github.com/1pedrohfreitas/pcams_back_go/models"
	"github.com/gin-gonic/gin"
)

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de Id inválido",
		})

		return
	}

	db := database.GetDataBase()

	var user models.User
	rows, err := db.Query(`Select * from users where id=?`, newid)
	database.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var fullname string

		err = rows.Scan(&fullname)
		database.CheckError(err)

		fmt.Println(fullname)
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi encontrado esse usuario",
		})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro no Json",
		})
		return
	}

	err = db.QueryRow(
		`INSERT INTO users (fullname, alias, username, usertype, status, "password") VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id`, user.FullName,
		user.Alias,
		user.UserName,
		user.UserType,
		user.Status,
		user.Password,
	).Scan(&user.ID)

	database.CheckError(err)

	c.JSON(201, user)
}

func ShowUsers(c *gin.Context) {
	db := database.GetDataBase()

	var result dto.PageResultDTO
	var users []models.User

	rows, err := db.Query(`Select * from "users"`)
	database.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Alias,
			&user.UserName,
			&user.UserType,
			&user.Status,
			&user.Password,
			&user.Created_at,
			&user.Updated_at,
		)
		database.CheckError(err)
		// user := models.User{
		// 	FullName: fullname,
		// }
		users = append(users, user)
		// fmt.Println(fullname)
	}

	result.Data = append(result.Data, users)
	c.JSON(200, result)

}
func UpdateUser(c *gin.Context) {
	// db := database.GetDataBase()

	var user models.User

	// err := c.ShouldBindJSON(&user)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "Erro no Json",
	// 	})
	// 	return
	// }

	// err = db.Save(&user).Error

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	// id := c.Param("id")

	// newid, err := strconv.Atoi(id)

	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "Formato de Id inválido",
	// 	})

	// 	return
	// }

	// db := database.GetDataBase()

	// err = db.Delete(&models.User{}, newid).Error
	c.Status(204)
}
