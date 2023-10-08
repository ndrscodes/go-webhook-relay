package relay

import (
	"webhook-proxy/mappers"
	"webhook-proxy/rendering"
	"webhook-proxy/rendering/resolvers"
)

type Endpoint struct {
	Url      string
	Renderer rendering.Renderer
	Resolver resolvers.Resolver
	Mapper   mappers.Mapper
}
