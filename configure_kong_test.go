package main

import "testing"

func TestHostAPI(t *testing.T) {
	t.Run("ShouldThrowErrorWhenAnyRoutingOptionProvided", func(t *testing.T) {
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
		api := &KongAPI{}

		api.SetHostsAPI("localhost", "127.0.0.1", "0.0.0.0")

		host, err := api.HostsAPI()
		if host == "" {
			t.Error("unnexpected host value, should not be empty")
		}

		if err != nil {
			t.Error("expected error, should be empty")
		}
	})

	t.Run("ShouldReturnNotValues", func(t *testing.T) {
		// api := &KongAPI{}

		//TODO: implement this test
	})
}
