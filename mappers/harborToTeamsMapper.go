package mappers

import (
	"fmt"
	"log"
)

type HarborToTeamsMapper struct {
}

func createDefaultText(in *map[string]interface{}) (text string) {
	return fmt.Sprintf("received %v", *in)
}

func (HarborToTeamsMapper) Map(in *map[string]interface{}, text *string) (out *map[string]interface{}, err error) {

	res := map[string]interface{}{}

	if text == nil {
		t := createDefaultText(in)
		text = &t
	}

	res["text"] = *text
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	out = &res

	return out, nil
}
