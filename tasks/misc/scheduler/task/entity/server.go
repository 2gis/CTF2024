package entity

type Server struct {}

func (server *Server) GetId() string {
	return "server"
}

func (server *Server) GetAdminLevel() int {
	return 2
}

func (server *Server) SetAdminLevel(level int) {
	return
}
