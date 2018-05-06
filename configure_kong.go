package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"github.com/pkg/errors"
)

const (
	url = "http://0.0.0.0:8001/apis/"
)

func requestMaker(param KongAPI) interface{} {
	fmt.Println(fmt.Sprintf("starting %s", param.Name))

	hosts, err := param.HostsAPI()
	if err != nil {
		panic(errors.Wrap(err, "host has a error"))
	}

	uris, err := param.UrisAPI()
	if err != nil {
		panic(errors.Wrap(err, "uris has a error"))
	}

	methods, err := param.MethodsAPI()
	if err != nil {
		panic(errors.Wrap(err, "methods has a error"))
	}

	newApiConfig := struct {
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
		Name:        param.Name,
		Hosts:       hosts,
		Uris:        uris,
		Methods:     methods,
		UpstreamURL: param.UpstreamURL,
	}

	client := &http.Client{}

	ApiURL := fmt.Sprintf("%s/%s", url, param.Name)
	requestDel, err := http.NewRequest("DELETE", ApiURL, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("pre config request operator: %+v", err.Error()))
	}

	responseDelete, err := client.Do(requestDel)
	if err != nil {
		log.Fatal(fmt.Sprintf("pre config request: %+v", err.Error()))
	}
	defer responseDelete.Body.Close()
	bodyDelete, err := ioutil.ReadAll(responseDelete.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("pre config response error: %+v", err.Error()))
	}
	fmt.Println(fmt.Sprintf("pre config response %s", bodyDelete))

	payloadsApi := new(bytes.Buffer)
	json.NewEncoder(payloadsApi).Encode(newApiConfig)

	requestPost, err := http.NewRequest("POST", url, payloadsApi)
	if err != nil {
		log.Fatal(fmt.Sprintf("config request operator %+v", err.Error()))
	}
	requestPost.Header.Add("Content-Type", "application/json")

	responsePost, err := client.Do(requestPost)
	if err != nil {
		log.Fatal(fmt.Sprintf("config request operation %+v", err.Error()))
	}
	defer responsePost.Body.Close()
	bodyPost, _ := ioutil.ReadAll(responsePost.Body)
	fmt.Println(fmt.Sprintf("config request operation response %s", bodyPost))

	fmt.Println(fmt.Sprintf("configuration finished for %s\n", param.Name))
	return bodyPost
}

// KongAPI Configuration structure
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

// HostsAPI kong host api getter
func (kong *KongAPI) HostsAPI() (string, error) {
	return kongRoutingOptionString(kong.Hosts, kong.Uris, kong.Methods)
}

// UrisAPI kong uris api getter
func (kong *KongAPI) UrisAPI() (string, error) {
	return kongRoutingOptionString(kong.Uris, kong.Methods, kong.Hosts)
}

// MethodsAPI kong methods api getter
func (kong *KongAPI) MethodsAPI() (string, error) {
	return kongRoutingOptionString(kong.Uris, kong.Methods, kong.Hosts)
}

func kongRoutingOptionString(target, alt1, alt2 []string) (string, error) {
	var msg error
	alt1Empty, alt2Empty := false, false
	if len(target) == 0 {
		alt1Empty = len(alt1) == 0
		if !alt1Empty {
			alt1Empty, _, _ = emptyValuesInSlice(alt1...)
		}

		alt2Empty = len(alt2) == 0
		if !alt2Empty {
			alt2Empty, _, _ = emptyValuesInSlice(alt2...)
		}

		if alt2Empty && alt1Empty {
			msg = fmt.Errorf("hosts, uris and methods option are empty, aleast one of the should not be empty")
		}
	}

	return strings.Join(target, ","), msg
}
