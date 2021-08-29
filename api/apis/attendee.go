package apis

import (
	model "HTA_api/api/models"
	"net/http"
	"strconv"

	//"strconv"
	orm "HTA_api/api/database"

	"github.com/gin-gonic/gin"
)

func Attendee_r(c *gin.Context) {
	var attendee []model.Attendee
	result := orm.Db.Find(&attendee)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "讀取失敗",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   001,
			"result": result.Value,
		})
	}

}

func Attendee_one(c *gin.Context) {
	var attendee model.Attendee
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
	} else {
		result := orm.Db.First(&attendee, id)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    001,
				"message": result,
			})
		}

	}
}

func Attendee_c(c *gin.Context) {
	var attendee model.Attendee
	c.BindJSON(&attendee)
	result := orm.Db.Create(&attendee)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result,
		})
	}
}

func Attendee_u(c *gin.Context) {
	var attendee model.Attendee
	c.BindJSON(&attendee)
	result := orm.Db.Save(&attendee)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result,
		})
	}
}

func Attendee_d(c *gin.Context) {
	var attendee model.Attendee
	c.BindJSON(&attendee)
	result := orm.Db.Delete(&attendee)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result,
		})
	}
}

/*
func findAll(table string) (result *gorm.DB) {
	switch table {
	case "attendee":
		var tables []model.Attendee
		result = orm.Db.Find(&tables)
	case "auth":
		var tables []model.Auth
		result = orm.Db.Find(&tables)
	}

	return result
}

func findone(table string, id int64) (result *gorm.DB) {
	switch table {
	case "attendee":
		var tables []model.Attendee
		result = orm.Db.First(&tables, id)
	case "auth":
		var tables []model.Auth
		result = orm.Db.First(&tables, id)
	}
	return result
}
*/
