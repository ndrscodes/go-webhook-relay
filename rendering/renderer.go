package rendering

import "webhook-proxy/rendering/resolvers"

type Renderer interface {
	Render(payload *map[string]any, resolver resolvers.Resolver) (text *string, err error)
}
