package apis

import (
	model "HTA_api/api/models"
	"net/http"

	//"strconv"
	orm "HTA_api/api/database"

	"github.com/gin-gonic/gin"
)

//select
func Auth_r(c *gin.Context) {
	var auth []model.Auth
	result := orm.Db.Find(&auth)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "讀取失敗",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Value,
		})
	}

}

func Auth_c(c *gin.Context) {
	var auth model.Auth
	c.BindJSON(&auth)
	result := orm.Db.Debug().Create(&auth)
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

func Auth_d(c *gin.Context) {
	var auth model.Auth
	c.BindJSON(&auth)
	result := orm.Db.Debug().Delete(&auth)
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
