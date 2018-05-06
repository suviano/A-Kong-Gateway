package main

import "testing"

func TestHostAPI(t *testing.T) {
	t.Run("ShouldThrowErrorWhenAnyRequiredRoutingOptionProvided", func(t *testing.T) {
		api := &KongAPI{}

		host, err := api.HostsAPI()
		if host != "" {
			t.Error("unnexpected host value, should be empty")
		}

		if err == nil {
			t.Error("expected error got any")
		}
	})

	t.Run("ShouldReturnFormattedHosts", func(t *testing.T) {
		api := &KongAPI{Hosts: []string{"localhost", "127.0.0.1", "0.0.0.0"}}

		host, err := api.HostsAPI()
		if host == "" {
			t.Error("unnexpected host value, should not be empty")
		}

		if err != nil {
			t.Error("expected error, should be empty")
		}
	})

	t.Run("ShouldReturnFormattedHostsIfNoRequiredOptionsIsEmpty", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"GET"},
			Uris: []string{"/test"},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedHostsIfOnlyMethodsIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{"GET"},
			Uris: []string{},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})

	t.Run("ShouldReturnFormattedHostsIfOnlyUrisIsDefined", func(t *testing.T) {
		api := &KongAPI{
			Methods: []string{},
			Uris: []string{"/test"},
		}

		_, err := api.HostsAPI()
		if err != nil {
			t.Errorf("expected error, should be nil %+v", err)
		}
	})
}

func TestUrisAPI(t *testing.T) {}

func TestMethodsAPI(t *testing.T) {}