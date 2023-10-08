package resolvers

type Resolver interface {
	Resolve(payload *map[string]interface{}) (path string)
}
