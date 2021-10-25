package apis

import (
	orm "HTA_api/api/database"
	model "HTA_api/api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Labor_hour_r(c *gin.Context) {
	var table []model.Labor_hour
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

func Labor_hour_one(c *gin.Context) {
	var table model.Labor_hour
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
				"message": result.Error,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    001,
				"message": result.Value,
			})
		}

	}
}

func Labor_hour_c(c *gin.Context) {
	var table model.Labor_hour
	c.BindJSON(&table)
	fmt.Println(&table)
	result := orm.Db.Debug().Create(&table)
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

func Labor_hour_u(c *gin.Context) {
	var table model.Labor_hour
	c.BindJSON(&table)
	result := orm.Db.Model(&table).Debug().Update(&table)
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

func Labor_hour_d(c *gin.Context) {
	var table model.Labor_hour
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
