package mappers

type Mapper interface {
	Map(in *map[string]any, message *string) (out *map[string]any, err error)
}
