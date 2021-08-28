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
	result, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
	} else {
		re := orm.Db.Find(&attendee, "A_id = ?", result)
		if re.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    001,
				"message": re,
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
