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

	s.router.GET(
		"/api/my/cards",
		s.handlerMyCards(),
	)

	s.router.GET(
		"/api/my/card/{id}",
		s.handlerMyCard(),
	)

}
