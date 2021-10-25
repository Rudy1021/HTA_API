package apis

import (
	orm "HTA_api/api/database"
	model "HTA_api/api/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

/*
func File_r(c *gin.Context) {
	var table []model.File
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
*/

func File_one(c *gin.Context) {
	var table model.File
	name := c.Param("name")
	result := orm.Db.First(&table, "name= ?", name)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else {
		file := table.Path + table.Name
		_, errByOpenFile := os.Open(file)
		if errByOpenFile != nil {
			return
		} else {
			c.Header("Content-Type", "application/octet-stream")
			c.Header("Content-Disposition", "attachment; filename="+table.Name)
			c.Header("Content-Transfer-Encoding", "binary")
			c.File(file)
			return
		}
	}

}

func File_c(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Panicln("c.FormFile('file')", err)
		return
	}
	filename := filepath.Base(file.Filename)
	err = c.SaveUploadedFile(file, "./api/files/"+filename)
	var table model.File
	table.Name = file.Filename
	table.Path = "./api/files/"
	table.Create_time = time.Now()
	result := orm.Db.Create(&table)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": result.Error,
		})
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "上傳檔案出錯：" + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    001,
			"message": result.Value,
		})
	}
}

/*
func File_u(c *gin.Context) {
	var table model.File
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

func File_d(c *gin.Context) {
	var table model.File
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
	} else {
		result := orm.Db.Delete(&table, id)
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
*/
/*
func findAll(table string) (result *gorm.DB) {
	switch table {
	case "table":
		var tables []model.Files
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
		var tables []model.Files
		result = orm.Db.First(&tables, id)
	case "auth":
		var tables []model.Auth
		result = orm.Db.First(&tables, id)
	}
	return result
}
*/
