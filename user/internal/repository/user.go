package repository

type User struct {
	UserId         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}
