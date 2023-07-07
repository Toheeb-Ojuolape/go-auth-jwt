package models

type TodoList struct {
	ID    int32
	Title string
}

func (TodoList) TableName() string {
	return "todo_list"
}
