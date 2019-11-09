package main

import (
	"time"

	"github.com/cyberzeed/scg-go/business"
	"github.com/cyberzeed/scg-go/line"
	"github.com/cyberzeed/scg-go/series"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v8"
)

// Builder function for create new router object
func setupRouter(config *viper.Viper) *gin.Engine {
	// register custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("siriessizeindex", series.SizeIndex)
	}

	// prepare essential variables and objects
	router := gin.Default()
	store := persistence.NewInMemoryStore(time.Second)
	cacheTime := time.Minute * 5

	// create handler functions
	seriesHandler := series.CreateHandler()
	seriesIndexValueHandler := series.CreateIndexValueHandler()
	restaurantSearchHandler := business.CreateRestaurantSearchHandler(config)
	lineBroadcastHandler := line.CreateBroadcastHandler(config)
	lineWebhookHandler := line.CreateWebhookHandler(config)

	// mapping METHOD and PATH with handler
	router.GET("/series", cache.CachePage(store, cacheTime, seriesHandler))
	router.GET("/series/:index", cache.CachePage(store, cacheTime, seriesIndexValueHandler))
	router.GET("/restaurant/:area", cache.CachePage(store, cacheTime, restaurantSearchHandler))
	router.POST("/message/broadcast", cache.CachePage(store, cacheTime, lineBroadcastHandler))
	router.POST("/message/webhook", cache.CachePage(store, cacheTime, lineWebhookHandler))
	return router
}
