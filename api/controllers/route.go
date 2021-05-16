package controllers

import (
	"github.com/TanutN/Aqua/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/tracker_sos", middlewares.SetMiddlewareJSON(s.GetAllSos)).Methods("GET")
	s.Router.HandleFunc("/wearing_devices", middlewares.SetMiddlewareJSON(s.GetDevices)).Methods("GET")
	s.Router.HandleFunc("/create_sos", middlewares.SetMiddlewareJSON(s.CreateSos)).Methods("POST")
	//router := gin.Default()
	v2 := s.Gin.Group("v2")
	//v2.Use(middlewares.MiddlewareAuthentication()).Use(middlewares.JSONMiddleware())
	//v2.Use(middlewares.JSONMiddleware()).Use(middlewares.MiddlewareAuthentication())
	v2.Use(middlewares.JSONMiddleware())
	{
		sos := v2.Group("sos")
		{
			//GET ALL SOS
			sos.GET("/", s.RescuerCommit)           // All
			sos.GET("/status", s.RescuerCommit)     // By status / section object
			sos.GET("/:id/detail", s.RescuerCommit) // By status / section object

			//REASON
			sos.POST("/reasons", s.CreateReasons)
			sos.GET("/reasons", s.GetReasons)
			sos.PATCH("/reasons/:id", s.UpdateReasons)
			sos.DELETE("/reasons/:id", s.DeleteReasons)

			//RESCUER
			sos.POST("/rescuers", s.CreateRescuser)
			sos.GET("/rescuers", s.GetRescuser)
			sos.GET("/rescuers/:id", s.GetRescuserByID)
			sos.PATCH("/rescuers/:id", s.UpdateRescuser)
			sos.DELETE("/rescuers/:id", s.DeleteRescue)

			sos.POST("/rescuers/:id/generate-link", s.RescuerCommit)
			sos.POST("/rescuers/:id/commit", s.RescuerCommit)

			/*//RESCUER DETAIL
			sos.GET("/rescuers/:ref_id/status", s.RescuerCommit)
			sos.POST("/rescuer-commit", s.RescuerCommit)
			sos.POST("/choose-rescue", s.ChooseRescue)
			sos.POST("/create_rescue", s.CreateRescuser)
			sos.PATCH("/update_rescue/:id", s.UpdateRescuser)
			sos.DELETE("/delete_rescue/:id", s.DeleteRescue)*/
		}
	}
	//	router.Run()
}
