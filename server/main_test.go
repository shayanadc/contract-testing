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
		BrokerURL:       "https://shayanadc.pactflow.io",
		BrokerToken:     "v_kaCEiuq-RpacJQflqV4g",
		ProviderBaseURL: "http://127.0.0.1:8080",
		ProviderVersion: "1.0.0",
		ConsumerVersionSelectors: []types.ConsumerVersionSelector{
			{
				Consumer: "example-client",
				Tag:      "1.0.0",
			},
		},
		PublishVerificationResults: true,
	})

	if err != nil {
		t.Log(err)
	}
}
