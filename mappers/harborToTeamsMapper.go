package mappers

import (
	"fmt"
	"log"
)

type HarborToTeamsMapper struct {
}

func createDefaultText(in *map[string]any) (text string) {
	return fmt.Sprintf("received %v", *in)
}

func (HarborToTeamsMapper) Map(in *map[string]any, text *string) (out *map[string]any, err error) {

	res := map[string]any{}

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
