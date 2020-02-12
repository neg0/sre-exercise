package restful

import "github.com/valyala/fasthttprouter"

func NewRoutes() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.GET("/alerts", alerts)
	router.GET("/alerts/:service_name", alerts)
	router.GET("/objectives", objectives)
	router.GET("/objectives/:service_name", objectives)
	return router
}
