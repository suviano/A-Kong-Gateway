package main

import "testing"

func TestHostAPI(t *testing.T) {
	t.Run("ShouldThrowErrorWhenAnyRequiredRoutingOptionProvided", func(t *testing.T) {
		api := &KongAPI{}

		hosts, err := api.HostsAPI()
		if hosts != "" {
			t.Error("unnexpected hosts value, should be empty")
		}

		if err == nil {
			t.Error("expected error got any")
		}
	})

	t.Run("ShouldReturnFormattedHosts", func(t *testing.T) {
		api := &KongAPI{Hosts: []string{"localhost", "127.0.0.1", "0.0.0.0"}}

		hosts, err := api.HostsAPI()
		if hosts == "" {
			t.Error("unnexpected hosts value, should not be empty")
		}

		if err != nil {
			t.Error("expected error, should be empty")
		}
	})

	t.Run("ShouldReturnFormattedHostsIfNoRequiredOptionsIsEmpty", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"GET"},
			Uris:    []string{"/test"},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedHostsIfOnlyMethodsIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"GET"},
			Uris:    []string{},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedHostsIfOnlyUrisIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{},
			Uris:    []string{"/test"},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})
}

func TestUrisAPI(t *testing.T) {
	t.Run("ShouldThrowErrorWhenAnyRequiredRoutingOptionProvided", func(t *testing.T) {
		api := &KongAPI{}

		uris, err := api.UrisAPI()
		if uris != "" {
			t.Error("unnexpected uris value, should be empty")
		}

		if err == nil {
			t.Error("expected error got any")
		}
	})

	t.Run("ShouldReturnFormattedUris", func(t *testing.T) {
		api := &KongAPI{Uris: []string{"/test", "/api/test"}}

		uris, err := api.UrisAPI()
		if uris == "" {
			t.Error("unnexpected uris value, should not be empty")
		}

		if err != nil {
			t.Error("expected error, should be empty")
		}
	})

	t.Run("ShouldReturnFormattedUrisIfNoRequiredOptionsIsEmpty", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"GET"},
			Hosts:   []string{"localhost"},
		}

		_, err := api.UrisAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedUrisIfOnlyMethodsIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"GET"},
			Hosts:   []string{},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedUrisIfOnlyMethodsIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"POST"},
			Hosts:   []string{},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})
}

func TestMethodsAPI(t *testing.T) {
	t.Run("ShouldThrowErrorWhenAnyRequiredRoutingOptionProvided", func(t *testing.T) {
		api := &KongAPI{}

		uris, err := api.MethodsAPI()
		if uris != "" {
			t.Error("unnexpected methods value, should be empty")
		}

		if err == nil {
			t.Error("expected error got any")
		}
	})

	t.Run("ShouldReturnFormattedMethods", func(t *testing.T) {
		api := &KongAPI{Methods: []string{"GET", "POST"}}

		uris, err := api.MethodsAPI()
		if uris == "" {
			t.Error("unnexpected methods value, should not be empty")
		}

		if err != nil {
			t.Error("expected error, should be empty")
		}
	})

	t.Run("ShouldReturnFormattedMethodsIfNoRequiredOptionsIsEmpty", func(t *testing.T) {
		api := &KongAPI{
			Uris:  []string{"/test"},
			Hosts: []string{"localhost"},
		}

		_, err := api.UrisAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedMethodsIfOnlyMethodsIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"GET"},
			Hosts:   []string{},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedUrisIfOnlyMethodsIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"POST"},
			Hosts:   []string{},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})
}
