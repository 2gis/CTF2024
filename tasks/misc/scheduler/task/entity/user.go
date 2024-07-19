package entity

type User struct {
	Id string
	AdminLevel int
}

func (user *User) GetId() string {
	return user.Id
}

func (user *User) GetAdminLevel() int {
	return user.AdminLevel
}

func (user *User) SetAdminLevel(level int) {
	user.AdminLevel = level
}