package restful

import (
	"github.com/valyala/fasthttp"
)

func New() error {
	router := NewRoutes()

	println("API is running at: http://127.0.0.1:8080")
	return fasthttp.ListenAndServe(":8080", router.Handler)
}
