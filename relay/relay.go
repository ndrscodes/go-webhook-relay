package relay

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

type Relay struct {
	endpoints []*Endpoint
}

func NewRelay(endpoint ...*Endpoint) Relay {
	return Relay{
		endpoints: endpoint,
	}
}

func (r *Relay) AddEndpoints(endpoint ...*Endpoint) {
	r.endpoints = append(r.endpoints, endpoint...)
}

func (r Relay) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	data := &map[string]any{}

	err := json.NewDecoder(req.Body).Decode(data)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}

	r.relay(data)
}

func stringify(v any) (out string, err error) {
	inBuf := bytes.Buffer{}
	encoder := json.NewEncoder(&inBuf)
	encoder.SetIndent("", " ")
	err = encoder.Encode(v)
	out = inBuf.String()
	return out, err
}

func prepare(in *map[string]any) (err error) {
	(*in)["stringified"], err = stringify(*in)

	return err
}

func processEndpoint(endpoint *Endpoint, data *map[string]any) (err error) {

	err = prepare(data)
	if err != nil {
		log.Println(err)
		return err
	}

	var template *string

	if endpoint.Renderer != nil {
		template, err = endpoint.Renderer.Render(data, endpoint.Resolver)
		if err != nil {
			log.Println(err)
		}
	}

	out, err := endpoint.Mapper.Map(data, template)
	if err != nil {
		log.Println(err)
		return err
	}

	res := &bytes.Buffer{}
	err = json.NewEncoder(res).Encode(out)
	if err != nil {
		return err
	}

	resp, err := http.Post(endpoint.Url, "application/json; charset=utf-8", res)
	if err != nil {
		return err
	}

	st, _ := io.ReadAll(resp.Body)
	log.Print(string(st))

	return nil
}

func (r Relay) relay(data *map[string]any) {
	errors := make(chan error, len(r.endpoints))
	wg := sync.WaitGroup{}
	wg.Add(len(r.endpoints))
	for _, endpoint := range r.endpoints {
		go func(ep *Endpoint) {
			defer wg.Done()
			err := processEndpoint(ep, data)
			if err != nil {
				errors <- err
			}
		}(endpoint)
	}
	wg.Wait()
}
