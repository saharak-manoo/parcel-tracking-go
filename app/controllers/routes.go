package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/tracks/kerry-express", s.KerryExpress)
		v1.GET("/tracks/jt-express", s.JTExpress)
		v1.GET("/tracks/flash-express", s.FlashExpress)
	}
}