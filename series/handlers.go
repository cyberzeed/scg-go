package series

import (
	"github.com/gin-gonic/gin"
)

type seriesParam struct {
	Size int `form:"size" binding:"siriessizeindex"`
}

// seriesIndex will bind index value of from URI
type seriesIndex struct {
	Value int `uri:"index" binding:"exists,siriessizeindex"`
}

// CreateHandler is a factory function for create seriesHandler
func CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Binding request parameter
		var param seriesParam
		if err := c.ShouldBind(&param); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Find value of series items and response
		series := make([]int, param.Size)
		for i := 0; i < param.Size; i++ {
			series[i] = getScgSeriesItem(i)
		}
		c.JSON(200, gin.H{"series": series})
	}
}

// CreateIndexValueHandler is a factory function for create seriesIndexHandler
func CreateIndexValueHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Binding URI parameter
		var index seriesIndex
		if err := c.ShouldBindUri(&index); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Find value of series item by index and response
		value := getScgSeriesItem(index.Value)
		c.JSON(200, gin.H{"value": value})
	}
}
