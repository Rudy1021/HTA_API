package apis

import (
	model "HTA_api/api/models"
	"net/http"

	//"strconv"
	orm "HTA_api/api/database"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	var auth model.Auth
	result := orm.Db.Find(&auth)
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": result,
	})
}
