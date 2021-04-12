package middleware

import (
	"github.com/GriezLiao/griez-go-tour/blog/pkg/app"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/errorcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			token string
			ecode = errorcode.Success
		)
		if s, exist := context.GetQuery("token"); exist {
			token = s
		} else {
			token = context.GetHeader("token")
		}

		if token == "" {
			ecode = errorcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errorcode.UnauthorizedTokenTimeout
				default:
					ecode = errorcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errorcode.Success {
			response := app.NewResponse(context)
			response.ToErrorResponse(ecode)
			context.Abort()
			return
		}

		context.Next()
	}
}
