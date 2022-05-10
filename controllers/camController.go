package controllers

import (
	"strconv"

	"github.com/1pedrohfreitas/pcams_back_go/database"
	"github.com/1pedrohfreitas/pcams_back_go/dto"
	"github.com/1pedrohfreitas/pcams_back_go/models"
	"github.com/gin-gonic/gin"
)

func ShowCam(c *gin.Context) {
	db := database.GetDataBase()
	var cam models.Cams
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	cam.ID = uint(newid)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de Id inválido",
		})

		return
	}

	err2 := db.QueryRow(`SELECT devicename, alias, streamtype, urlcamstream, status, created_at, updated_at FROM cams;
	where id=$1`, newid).Scan(
		&cam.DeviceName,
		&cam.Alias,
		&cam.StreamType,
		&cam.UrlCamStream,
		&cam.Status,
		&cam.Created_at,
		&cam.Updated_at,
	)

	database.CheckError(err2)

	c.JSON(200, cam)
}

func CreateCam(c *gin.Context) {
	db := database.GetDataBase()
	var cam models.Cams
	err := c.ShouldBindJSON(&cam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro no Json",
		})
		return
	}

	err = db.QueryRow(
		`INSERT INTO pcam.cams (devicename, alias, streamtype, urlcamstream, status) VALUES($1, $2, $3, $4, $5)
		RETURNING id`, cam.DeviceName,
		cam.Alias,
		cam.StreamType,
		cam.UrlCamStream,
		cam.Status,
	).Scan(&cam.ID)

	database.CheckError(err)

	c.JSON(201, cam)
}

func ShowCams(c *gin.Context) {
	db := database.GetDataBase()

	var result dto.PageResultDTO
	var cams []models.Cams

	rows, err := db.Query(`SELECT id, devicename, alias, streamtype, urlcamstream, status, created_at, updated_at FROM cams`)
	database.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var cam models.Cams

		err = rows.Scan(
			&cam.ID,
			&cam.DeviceName,
			&cam.Alias,
			&cam.StreamType,
			&cam.UrlCamStream,
			&cam.Status,
			&cam.Created_at,
			&cam.Updated_at,
		)
		database.CheckError(err)
		cams = append(cams, cam)
	}

	result.Data = append(result.Data, cams)
	c.JSON(200, result)

}
func UpdateCam(c *gin.Context) {
	db := database.GetDataBase()
	var cam models.Cams
	err := c.ShouldBindJSON(&cam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro no Json",
		})
		return
	}

	_, err2 := db.Exec(`UPDATE cams SET devicename=$1, alias=$2, streamtype=$3, urlcamstream=$4, status=$5 WHERE id=$6`,
		cam.DeviceName,
		cam.Alias,
		cam.StreamType,
		cam.UrlCamStream,
		cam.Status,
		cam.ID)

	database.CheckError(err2)

	c.JSON(200, cam)
}

func DeleteCam(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de Id inválido",
		})

		return
	}
	db := database.GetDataBase()
	_, err2 := db.Exec("DELETE FROM cams WHERE id=$1", newid)
	database.CheckError(err2)
	c.Status(204)
}
