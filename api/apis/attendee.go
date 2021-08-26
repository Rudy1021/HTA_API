package apis

import (
	model "HTA_api/api/models"
	"net/http"

	//"strconv"
	orm "HTA_api/api/database"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	var attendee model.Attendee
	result := orm.Db.Find(&attendee)
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": result,
	})
}
