package controllers

import (
	"strconv"

	"github.com/1pedrohfreitas/pcams_back_go/database"
	"github.com/1pedrohfreitas/pcams_back_go/dto"
	"github.com/1pedrohfreitas/pcams_back_go/models"
	"github.com/gin-gonic/gin"
)

func ShowClient(c *gin.Context) {
	db := database.GetDataBase()
	var client models.Clients
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	client.ID = uint(newid)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de Id inválido",
		})

		return
	}

	err2 := db.QueryRow(`SELECT fullname, alias, codclient, status, created_at, updated_at FROM clients where id=$1`, newid).Scan(
		&client.FullName,
		&client.Alias,
		&client.CodClient,
		&client.Status,
		&client.Created_at,
		&client.Updated_at)

	database.CheckError(err2)

	c.JSON(200, client)
}

func CreateClient(c *gin.Context) {
	db := database.GetDataBase()
	var client models.Clients
	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro no Json",
		})
		return
	}

	err = db.QueryRow(
		`INSERT INTO pcam.clients (fullname, alias, status) VALUES($1, $2, $3) RETURNING id`,
		client.FullName,
		client.Alias,
		client.Status,
	).Scan(&client.ID)

	database.CheckError(err)

	c.JSON(201, client)
}

func ShowClients(c *gin.Context) {
	db := database.GetDataBase()

	var result dto.PageResultDTO
	var clients []models.Clients

	rows, err := db.Query(`SELECT id, fullname, alias, codclient, status, created_at, updated_at FROM clients`)
	database.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var client models.Clients

		err = rows.Scan(
			&client.ID,
			&client.FullName,
			&client.Alias,
			&client.CodClient,
			&client.Status,
			&client.Created_at,
			&client.Updated_at,
		)
		database.CheckError(err)
		clients = append(clients, client)
	}

	result.Data = append(result.Data, clients)
	c.JSON(200, result)

}
func UpdateClient(c *gin.Context) {
	db := database.GetDataBase()
	var client models.Clients
	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro no Json",
		})
		return
	}

	_, err2 := db.Exec(`UPDATE clients SET fullname=$1, alias=$2, status=$3 WHERE id=$4`,
		client.FullName,
		client.Alias,
		client.Status,
		client.ID)

	database.CheckError(err2)

	c.JSON(200, client)
}

func DeleteClient(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de Id inválido",
		})

		return
	}
	db := database.GetDataBase()
	_, err2 := db.Exec("DELETE FROM clients WHERE id=$1", newid)
	database.CheckError(err2)
	c.Status(204)
}
