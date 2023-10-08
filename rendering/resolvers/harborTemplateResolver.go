package resolvers

import (
	"errors"
	"fmt"
)

type HarborTemplateResolver struct {
}

func extract(payload *map[string]interface{}) (t string, err error) {
	t = ""

	tmp, ok := (*payload)["type"]
	if !ok {
		return t, errors.New("type parameter not found")
	}

	t, ok = tmp.(string)
	if !ok {
		return t, errors.New("type parameter malformed")
	}

	return t, nil
}

func (HarborTemplateResolver) Resolve(payload *map[string]interface{}) (path string) {
	t := "unknown"

	ts, err := extract(payload)
	if err == nil {
		t = ts
	}

	return fmt.Sprintf("templates/harbor/%s", t)
}
