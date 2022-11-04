package controller

import (
	"NetworkControl/app/mapper"
	"NetworkControl/app/response"
	"NetworkControl/model"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

// Pong 接受请求后输出pong!
func Pong(c echo.Context) error {
	if c.Get("User").(string) != "default" {
		return c.String(http.StatusForbidden, "")
	}
	return c.String(http.StatusOK, "pong!")
}

// Print 第一个接口 接受一个query参数，直接print这个参数
func Print(c echo.Context) error {

	// 如果不是默认用户
	if c.Get("User").(string) != "default" {
		return c.String(http.StatusForbidden, "")
	}

	return c.String(http.StatusOK, c.QueryParam("query"))
}

// Analybody 第二个接口 解析request body并print
func Analybody(c echo.Context) error {

	// 如果不是默认用户
	if c.Get("User").(string) != "default" {
		return c.String(http.StatusForbidden, "")
	}

	// 分析
	var m model.Person
	err2 := c.Bind(&m)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(m)
	return response.SendResponse(c, http.StatusOK, "ok", m)
}

// AddTodo 添加todo接口
func AddTodo(e echo.Context) error {
	var req model.AddTodoReq
	err0 := e.Bind(&req)
	if err0 != nil {
		return err0
	}
	id, err1 := mapper.AddTodo(req.Title, req.Content, 1)
	if err1 != nil {
		return err1
	}

	err2 := e.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
	if err2 != nil {
		return err2
	}
	return nil
}

// GetTodo 查询todo接口
func GetTodo(e echo.Context) error {

	todoid, err := strconv.ParseUint(e.QueryParam("todo_id"), 10, 32)
	if err != nil {
		return err
	}

	todo, err := mapper.GetTodoByTodoId(uint(todoid))
	err1 := e.JSON(http.StatusOK, map[string]any{
		"content": todo.Content,
		"title":   todo.Title,
		"id":      todo.ID,
	})
	if err1 != nil {
		return err1
	}

	return nil
}

// Login 实现 login 接口
func Login(e echo.Context) error {
	var req model.Login
	err0 := e.Bind(&req)
	if err0 != nil {
		return err0
	}
	judge, err1 := mapper.GetLogin(req)
	if err1 != nil {
		log.Fatalln(err1)
	}
	if judge == false {
		err2 := e.JSON(http.StatusForbidden, map[string]any{
			"msg": "login fail",
		})
		if err2 != nil {
			return err2
		}
	}
	err2 := e.JSON(http.StatusOK, map[string]any{
		"token": "string",
	})
	if err2 != nil {
		return err2
	}
	return nil

}
