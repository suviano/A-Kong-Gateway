package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	url = "http://0.0.0.0:8001/apis/"
)

func requestMaker(param KongAPI) interface{} {
	hosts, err := param.HostsAPI()
	if err != nil {
		panic(err)
	}

	uris, err := param.UrisAPI()
	if err != nil {
		panic(err)
	}

	methods, err := param.MethodsAPI()
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(struct {
		Name                   string `json:"name"`
		Hosts                  string `json:"hosts"`
		Uris                   string `json:"uris"`
		Methods                string `json:"methods"`
		UpstreamURL            string `json:"upstream_url"`
		StripURI               bool   `json:"strip_uri,omitempty"`
		PreserveHost           bool   `json:"preserve_host,omitempty"`
		Retries                int    `json:"retries,omitempty"`
		UpstreamConnectTimeout int    `json:"upstream_connect_timeout,omitempty"`
		UpstreamSendTimeout    int    `json:"upstream_send_timeout,omitempty"`
		UpstreamReadTimeout    int    `json:"upstream_read_timeout,omitempty"`
		HTTPSOnly              bool   `json:"https_only,omitempty"`
		HTTPIfTerminated       bool   `json:"http_if_terminated,omitempty"`
	}{
		Name:    param.Name,
		Hosts:   hosts,
		Uris:    uris,
		Methods: methods,
	})

	r, err := http.Post(fmt.Sprintf("%s/apis", ""), "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("%v\n", body)

	return body
}

type KongAPI struct {
	Name                   string
	Hosts                  []string
	Uris                   []string
	Methods                []string
	UpstreamURL            string
	StripURI               bool
	PreserveHost           bool
	Retries                int
	UpstreamConnectTimeout int
	UpstreamSendTimeout    int
	UpstreamReadTimeout    int
	HTTPSOnly              bool
	HTTPIfTerminated       bool
}

// Kong configurable api
type Kong interface {
	Save(name string, host interface{}) error
	HostsAPI() (string, error)
	UrisAPI() (string, error)
	MethodsAPI() (string, error)
}

func (kong *KongAPI) HostsAPI() (string, error) {
	return kongRoutingOptionString(kong.Hosts, kong.Uris, kong.Methods)
}

func (kong *KongAPI) SetHostsAPI(hosts ...string) (err error) {
	_, kong.Hosts, err = emptyValuesInSlice(hosts...)
	return
}

func (kong *KongAPI) UrisAPI() (string, error) {
	return kongRoutingOptionString(kong.Uris, kong.Methods, kong.Hosts)
}

func (kong *KongAPI) SetUrisAPI(uris ...string) (err error) {
	_, kong.Uris, err = emptyValuesInSlice(uris...)
	return
}

func (kong *KongAPI) MethodsAPI() (string, error) {
	return kongRoutingOptionString(kong.Uris, kong.Methods, kong.Hosts)
}

func (kong *KongAPI) SetMethodsAPI(methods ...string) (err error) {
	_, kong.Methods, err = emptyValuesInSlice(methods...)
	return
}

func kongRoutingOptionString(target, alt1, alt2 []string) (string, error) {
	var msg error
	urisEmpty, methodsEmpty := false, false
	if len(target) == 0 {
		urisEmpty = len(alt1) == 0
		if !urisEmpty {
			urisEmpty, _, _ = emptyValuesInSlice(alt1...)
		}

		methodsEmpty = len(alt2) == 0
		if !methodsEmpty {
			methodsEmpty, _, _ = emptyValuesInSlice(alt2...)
		}

		if methodsEmpty && urisEmpty {
			msg = fmt.Errorf("hosts, uris and methods option are empty, aleast one of the should not be empty")
		}
	}

	return strings.Join(target, ","), msg
}

func main() {
	// http://www.michaelbach.de/ot/sze_silhouette/index.html
	// http://www.exploratorium.edu/files/exhibits/fading_dot/fading_dot.html
	// http://www.visnos.com/demos/fractal
	// "iamninja", []string{"imaninja.com"}, nil, []string{"GET", "OPTIONS"}, "http://imaninja.com"
}
