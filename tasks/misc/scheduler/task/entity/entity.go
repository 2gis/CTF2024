package entity

type Entity interface {
	GetId() string
	GetAdminLevel() int
	SetAdminLevel(int)
}

var Entities = map[string]Entity{
	"server": &Server{},
}

func GetEntityById(find string) Entity {
	for _, entity := range Entities {
		if find == entity.GetId() {
			return entity
		}
	}
	return nil
}