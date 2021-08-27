package apis

import (
	model "HTA_api/api/models"
	"net/http"

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
	result := orm.Db.Debug().Delete(&attendee)
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

func Test(c *gin.Context) {
	var attendee model.Attendee
	c.BindJSON(&attendee)
	orm.Db.Find(&attendee, "A_id = ?", attendee.A_id)
	c.JSON(http.StatusOK, gin.H{
		"message": attendee,
		"asd":     attendee.Customer_id,
	})
}
