package midleware

import (
	"api_gateway/api/token"
	"api_gateway/logs"
	"fmt"
	"log/slog"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuhtMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstr, err := ctx.Cookie("access_token")
		if err!= nil || tokenstr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
            ctx.Abort()
            return
        }
		if tokenstr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}
		
		volid, err := token.ValidateToken(tokenstr)
		if err != nil || !volid {
			fmt.Println("salom")
			ctx.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		claims, err := token.ExtractClaimsAccess(tokenstr)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			ctx.Abort()
			return
		}

		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}

}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logs.Logger.Info("logger middleware",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path))

		ctx.Next()
		logs.Logger.Info("Response sent",
			slog.Int("status", ctx.Writer.Status()),
		)
	}
}

func Authorization(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims := ctx.MustGet("claims").(jwt.MapClaims)

		path := ctx.FullPath()

		method := ctx.Request.Method
		ok, err := enforcer.Enforce(claims["role"].(string), path, method)
		fmt.Println(claims["role"].(string))
		if err!= nil {
			fmt.Println(err)
            ctx.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
            ctx.Abort()
            return
        }

		if!ok {
			fmt.Println(ok)
            ctx.AbortWithStatusJSON(403, gin.H{"error": "forbidden false"})
            ctx.Abort()
            return
        }
		ctx.Next()
	}

}
