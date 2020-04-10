package server

func (s *Server) InitRoutes() {
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

	s.router.POST(
		"/api/card/0",
		s.handlerCreateCard(),
	)

	s.router.POST(
		"/api/card/lock",
		s.handlerCardLock(),
	)

	s.router.POST(
		"/api/card/unlock",
		s.handlerCardUnlock(),
	)

	s.router.POST(
		"/api/my/card/{id}/lock",
		s.handlerMyCardLock(),
	)

	s.router.POST(
		"/api/my/card/{id}/unlock",
		s.handlerMyCardUnlock(),
	)
}
