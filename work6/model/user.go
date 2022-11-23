package model

type User struct {
	Username string
	Password string
	Question string
	Answer   string
}
type Comments struct {
	ID       uint
	Yourname string //留言对象
	Content  string
	Myname   string //留言人
}
