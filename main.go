package main

import (
	"net/http"
	"webhook-proxy/mappers"
	"webhook-proxy/relay"
	"webhook-proxy/rendering"
	"webhook-proxy/rendering/resolvers"
)

func main() {
	r := relay.NewRelay(&relay.Endpoint{
		Url:      "http://echo.free.beeceptor.com",
		Renderer: &rendering.GoTemplateRenderer{},
		Resolver: &resolvers.HarborTemplateResolver{},
		Mapper:   &mappers.HarborToTeamsMapper{},
	})
	http.ListenAndServe(":8080", r)
}
