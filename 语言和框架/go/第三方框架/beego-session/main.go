package main

import (
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/astaxie/beego/session"
)

// 整个web进程生命周期中的session管理器
var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "beegosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func indexHandler(resp http.ResponseWriter, req *http.Request) {
	sess, _ := globalSessions.SessionStart(resp, req)
	defer sess.SessionRelease(resp)
	username := sess.Get("username")
	user := map[string]interface{}{
		"name": username,
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(resp, user)
}

func loginHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	user := map[string]interface{}{}

	req.ParseForm()
	userinfo := req.Form
	log.Printf("认证信息: %+v\n", userinfo)
	// username := userinfo["username"][0]
	// password := userinfo["password"][0]
	username := req.Form.Get("username")
	password := req.Form.Get("password")

	if username == "admin" && password == "123456" {
		sess, _ := globalSessions.SessionStart(resp, req)
		defer sess.SessionRelease(resp)
		sess.Set("username", username)
		user["name"] = username
		// 登录成功, 跳转到首页
		http.Redirect(resp, req, "/", 301)
	} else {
		resp.Write([]byte(`{"msg": "用户名/密码错误"}`))
	}

	return
}

func logoutHandler(resp http.ResponseWriter, req *http.Request) {
	// SessionStart()获取到的session对象有一个Delete()方法,
	// 但那是在当前用户登录状态生命周期时, 可以添加session级别的变量而存在的.
	// SessionDestroy()可以直接销毁session会话本身
	globalSessions.SessionDestroy(resp, req)
	// 注销成功, 跳转到首页
	http.Redirect(resp, req, "/", 301)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	log.Println("http server started")
	http.ListenAndServe(":8079", nil)
}
