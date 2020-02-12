package restful

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
	internalServices "uswitch/internal/services"
	"uswitch/internal/services/restful/service"
)

func objectives(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
	body, _ := service.ObjectivesHandler(
		params.ByName("service_name"),
		internalServices.Instance(),
	)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}
