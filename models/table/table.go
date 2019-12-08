package table

import "time"

//验证登录
type Login struct{
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Login    string `gorm:"default:'False'"json:"login"`
	Role     string `json:"role"`
	RoleId 	 uint   `json:"role_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PersonMappings struct{
	StuTeaId uint `json:"stu_tea_id"`
	Pid uint `json:"pid"`
	Cid uint `json:"cid"`
	DeletedAt *time.Time `json:"deleted_at"`
}

