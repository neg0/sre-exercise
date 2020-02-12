package restful

import (
	internalServices "uswitch/internal/services"
	"uswitch/internal/services/restful/service"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func alerts(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
	body, _ := service.AlertsHandler(
		params.ByName("service_name"),
		internalServices.Instance(),
	)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}
