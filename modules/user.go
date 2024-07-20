package modules

type User struct {
	Id       int    `json:"id "`
	Username string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email"`
	AddTime  int    `json:"add_time"`
}

func (User) TableName() string {
	return "user"
}
