package controllers

import (
	"strconv"

	"github.com/1pedrohfreitas/pcams_back_go/database"
	"github.com/1pedrohfreitas/pcams_back_go/dto"
	"github.com/1pedrohfreitas/pcams_back_go/models"
	"github.com/gin-gonic/gin"
)

func ShowLicence(c *gin.Context) {
	db := database.GetDataBase()
	var licence models.Licences
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	licence.ID = uint(newid)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de Id inválido",
		})

		return
	}

	err2 := db.QueryRow(`SELECT uniccode, clientid, installerid, licencetype, created_at, updated_at FROM licences	where id=$1`, newid).Scan(
		&licence.UnicCode,
		&licence.ClientId.ID,
		&licence.InstallerId.ID,
		&licence.LicenceType,
		&licence.Created_at,
		&licence.Updated_at)

	database.CheckError(err2)

	c.JSON(200, licence)
}

func CreateLicence(c *gin.Context) {
	db := database.GetDataBase()
	var licence models.Licences
	err := c.ShouldBindJSON(&licence)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro no Json",
		})
		return
	}

	err = db.QueryRow(
		`INSERT INTO pcam.licences (clientid, installerid, licencetype) VALUES($1, $2, $3)
		RETURNING id`,
		licence.ClientId.ID,
		licence.InstallerId.ID,
		licence.LicenceType,
	).Scan(&licence.ID)

	database.CheckError(err)

	c.JSON(201, licence)
}

func ShowLicences(c *gin.Context) {
	db := database.GetDataBase()

	var result dto.PageResultDTO
	var licences []models.Licences

	rows, err := db.Query(`SELECT id, uniccode, clientid, installerid, licencetype, created_at, updated_at FROM licences`)
	database.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var licence models.Licences

		err = rows.Scan(
			&licence.ID,
			&licence.UnicCode,
			&licence.ClientId.ID,
			&licence.InstallerId.ID,
			&licence.LicenceType,
			&licence.Created_at,
			&licence.Updated_at,
		)
		database.CheckError(err)
		licences = append(licences, licence)
	}

	result.Data = append(result.Data, licences)
	c.JSON(200, result)

}
func UpdateLicence(c *gin.Context) {
	db := database.GetDataBase()
	var licence models.Licences
	err := c.ShouldBindJSON(&licence)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro no Json",
		})
		return
	}
	licence.EndDateValidate = licence.PrevDateEndValidate.AddDate(0, 0, 5)

	_, err2 := db.Exec(`UPDATE pcam.licences SET licencetype='', startvalidate='', endvalidate='', prevendvalidate='' WHERE id=$7`,
		licence.LicenceType,
		licence.StartDateValidate,
		licence.PrevDateEndValidate,
		licence.EndDateValidate,
		licence.ID)

	database.CheckError(err2)

	c.JSON(200, licence)
}

func DeleteLicence(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de Id inválido",
		})

		return
	}
	db := database.GetDataBase()
	_, err2 := db.Exec("DELETE FROM licences WHERE id=$1", newid)
	database.CheckError(err2)
	c.Status(204)
}
