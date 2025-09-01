package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "session_id"

// use cookie to store session id
func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey))
	return sessions.Sessions("mysession", store)
}

// user auth session middle
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get(userkey)
		if sessionID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message:": "此页面需要登入",
			})
			return
		}
		c.Next()
	}
}

// save session for user
func SaveSession(c *gin.Context, userID int) {
	session := sessions.Default(c)
	session.Set(userkey, userID)
	session.Save()
}

// clear session for user
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// get session for user
func GetSession(c *gin.Context) int {
	session := sessions.Default(c)
	sessionID := session.Get(userkey)
	if sessionID == nil {
		return 0
	}
	return sessionID.(int)
}

// check sessio for user
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionID := session.Get(userkey)
	return sessionID != nil
}
