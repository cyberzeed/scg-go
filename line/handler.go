package line

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
)

type broadcastMessages struct {
	messages []string `uri:"messages" binding:"required"`
}

// CreateBroadcastHandler is a factory function for create lineBroadcastHandler
func CreateBroadcastHandler(config *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get broadcast message from request body
		var broadcast broadcastMessages
		if err := c.ShouldBind(&broadcast); err != nil {
			log.Println(err.Error())
			c.Writer.WriteHeader(400)
			return
		}

		// send broadcast message
		resp, err := createBot(config).BroadcastMessage(broadcast.messages...)
		if err != nil {
			log.Println(err.Error())
			c.Writer.WriteHeader(500)
			return
		}

		c.JSON(200, gin.H{"response": resp})
	}
}

// CreateWebhookHandler is a factory function for create lineWebhookHandler
func CreateWebhookHandler(config *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		// parsing request
		err := createBot(config).ParseRequest(c.Request)

		// response status code
		if err == linebot.ErrInvalidSignature {
			log.Println(err.Error())
			c.Writer.WriteHeader(400)
		} else if err != nil {
			log.Println(err.Error())
			c.Writer.WriteHeader(500)
		} else {
			c.Writer.WriteHeader(200)
		}
	}
}

func createBot(config *viper.Viper) *Bot {
	accessToken := config.GetString("linebot.accesstoken")
	secret := config.GetString("linebot.secret")
	return NewBot(accessToken, secret)
}
