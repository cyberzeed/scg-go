package business

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type businessSearchOption struct {
	Area     string `uri:"area" binding:"required"`
	Business string
}

// CreateRestaurantSearchHandler is a factory function for create restaurant search
func CreateRestaurantSearchHandler(config *viper.Viper) gin.HandlerFunc {
	const business = "restaurant"
	finder := NewFinder(config.GetString("google.apikey"))

	return func(c *gin.Context) {
		option := businessSearchOption{Business: business}
		if err := c.ShouldBindUri(&option); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		results, err := finder.Search(With(option.Area, option.Business))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"results": results})
	}
}
