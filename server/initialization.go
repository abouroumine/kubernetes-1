package server

func (s *Server) Initialize() {
	s.InitDB()
	s.InitServer()
}

func (s *Server) InitDB() {

}

func (s *Server) InitServer() {
	s.NewRouter()
	s.S.Run("0.0.0.0:8080")
}
