package main

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestServerPact_Verification(t *testing.T) {

	pact := dsl.Pact{
		Provider: "example-server",
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"../client/pacts/example-client-example-server.json"},
	})

	if err != nil {
		t.Log(err)
	}
}
