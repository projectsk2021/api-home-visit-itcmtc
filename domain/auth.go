package domain

type UserLoginForm struct {
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
