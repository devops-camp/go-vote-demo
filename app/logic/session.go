package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var sessionName = "session-name"

func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName)
	return session.Values
}

func SetSession(c *gin.Context, name string, id int64) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["name"] = name
	session.Values["id"] = id

	return session.Save(c.Request, c.Writer)
}

func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)

	session.Values["name"] = ""
	session.Values["id"] = int64(-1)

	return session.Save(c.Request, c.Writer)
}

// ClearSession 清除 session, 而非像 FlushSession 替换值
// https://github.com/gorilla/sessions/issues/211
// https://github.com/gorilla/sessions/issues/160
func ClearSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)

	// session.Values["name"] = ""
	// session.Values["id"] = int64(-1)

	session.Options.MaxAge = -1
	return session.Save(c.Request, c.Writer)
}
