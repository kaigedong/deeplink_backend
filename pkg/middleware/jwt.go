package middleware

// 和token有点类似，在服务端加入了一个secret密钥，由用户发送用户名密码给服务端，服务端验证。
// 成功之后就生成三个部分header, payload, signature组成的jwt token给客户端，
// 之后的请求都带上 jwt token，服务端通过secret密钥进行验证。

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var identityKey = "id"

const TokenExpireDuration = time.Hour * 24

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// the jwt middleware
// var AuthMiddleware, err =

func JwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     TokenExpireDuration, // 过期时间
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals UserClaim
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userId := loginVals.Username
			password := loginVals.Password
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return &User{
					UserName:  userId,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"code": code, "message": message})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
