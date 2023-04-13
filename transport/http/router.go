package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	v1.POST("/transaction", s.handler.CreateTransaction)
	v1.PUT("/transaction/:id", s.handler.UpdateTransaction)
	v1.GET("/transaction", s.handler.ListTransactions)
	v1.GET("/transaction/:id", s.handler.GetTransaction)
	v1.DELETE("/transaction/:id", s.handler.DeleteTransaction)
}
