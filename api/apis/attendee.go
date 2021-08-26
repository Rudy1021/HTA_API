package apis

import (
	model "HTA_api/api/models"
	"net/http"

	//"strconv"
	orm "HTA_api/api/database"

	"github.com/gin-gonic/gin"
)

func Attendee_s(c *gin.Context) {
	var attendee model.Attendee
	result := orm.Db.Find(&attendee)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "取得失敗",
		})
	}

}

func Attendee_c(c *gin.Context) {
	var attendee model.Attendee
	attendee.Customer_id = 55642
	attendee.User_id = 22322
	result := orm.Db.Create(&attendee)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "新增失敗",
		})
	}
}

func Attendee_u(c *gin.Context) {
	var attendee model.Attendee
	orm.Db.Find(&attendee, "A_id = ?", "4")
	attendee.Customer_id = 1234567
	result := orm.Db.Model(&attendee).Where("A_id = ?", attendee.A_id).Update("Customer_id", attendee.Customer_id)
	c.JSON(http.StatusOK, gin.H{
		"code":    001,
		"message": result,
	})
}

func Attendee_d(c *gin.Context) {
	var attendee model.Attendee
	orm.Db.Find(&attendee, "A_id = ?", "11")
	result := orm.Db.Model(&attendee).Where("A_id = ?", &attendee.A_id).Delete(&attendee)
	c.JSON(http.StatusOK, gin.H{
		"code":    001,
		"message": result,
	})
}
