package model

type Person struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Todo struct {
	ID      uint `gorm:"PrimaryKey"`
	Title   string
	Content string
	UserId  uint
}

type AddTodoReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Login struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
