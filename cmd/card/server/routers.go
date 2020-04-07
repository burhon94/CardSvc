package server

func (s *Server) InitRoutes()  {
	s.router.GET(
		"/api/status",
		s.handlerHealth(),
	)

	s.router.GET(
		"/api/get/cards",
		s.handlerGetCards(),
	)
}
