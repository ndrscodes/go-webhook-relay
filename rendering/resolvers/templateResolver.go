package resolvers

type Resolver interface {
	Resolve(payload *map[string]any) (path string)
}
