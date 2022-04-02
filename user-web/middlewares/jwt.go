package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/models"
	"net/http"
	"time"
)

func JWTAuth() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			ctx.Abort()
			return
		}
		j:=NewJWT()
		claims, err := j.ParseToken(token)
		if err!=nil{
			if err == TokenExpired{
				ctx.JSON(http.StatusUnauthorized,map[string]string{
					"msg": "授权已过期",
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusUnauthorized,map[string]string{
				"msg":"未登陆",
			})
			ctx.Abort()
			return
		}
		ctx.Set("claims",claims)
		ctx.Set("userId",claims.Id)
		ctx.Next()
	}
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SigningKey []byte
}

//初始化jwt对象
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JWTInfo.SigningKey),
	}
}

//创建一个token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//解析token
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

//更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	//更新时间
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
