package apis

import (
	orm "HTA_api/api/database"
	model "HTA_api/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Work_item_r(c *gin.Context) {
	var table []model.Work_item
	result := orm.Db.Find(&table)
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

func Work_item_one(c *gin.Context) {
	var table model.Work_item
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
	} else {
		result := orm.Db.First(&table, id)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    001,
				"message": result.Value,
			})
		}

	}
}

func Work_item_c(c *gin.Context) {
	var table model.Work_item
	c.BindJSON(&table)
	result := orm.Db.Create(&table)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result.Value,
		})
	}
}

func Work_item_u(c *gin.Context) {
	var table model.Work_item
	c.BindJSON(&table)
	result := orm.Db.Save(&table)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result.Value,
		})
	}
}

func Work_item_d(c *gin.Context) {
	var table model.Work_item
	c.BindJSON(&table)
	result := orm.Db.Delete(&table)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result.Value,
		})
	}
}

/*
func findAll(table string) (result *gorm.DB) {
	switch table {
	case "table":
		var tables []model.Work_item
		result = orm.Db.Find(&tables)
	case "auth":
		var tables []model.Auth
		result = orm.Db.Find(&tables)
	}

	return result
}

func findone(table string, id int64) (result *gorm.DB) {
	switch table {
	case "table":
		var tables []model.Work_item
		result = orm.Db.First(&tables, id)
	case "auth":
		var tables []model.Auth
		result = orm.Db.First(&tables, id)
	}
	return result
}
*/
