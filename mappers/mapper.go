package mappers

type Mapper interface {
	Map(in *map[string]interface{}, message *string) (out *map[string]interface{}, err error)
}
