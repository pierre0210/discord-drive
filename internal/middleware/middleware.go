package middleware

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func AddSession(session *discordgo.Session) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("bot", session)
		ctx.Next()
	}
}
