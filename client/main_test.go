package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestClientPact_Local(t *testing.T) {

	pact := dsl.Pact{
		Consumer: "example-client",
		Provider: "example-server",
	}

	pact.Setup(true)

	t.Run("get user list", func(t *testing.T) {

		pact.
			AddInteraction().
			Given("User Lists").
			UponReceiving("User Collection is requested").
			WithRequest(dsl.Request{
				Method: "GET",
				Path:   dsl.Term("/users", "/users"),
			}).
			WillRespondWith(dsl.Response{
				Status: 200,
				Body: dsl.Like([]User{
					{
						ID:   1,
						Name: "John",
					},
					{
						ID:   2,
						Name: "Alice",
					},
				}),
			})

		err := pact.Verify(func() error {
			host := fmt.Sprintf("%s:%d", pact.Host, pact.Server.Port)

			users, err := GetUser(host)

			if err != nil {
				return errors.New("error is not expected")
			}

			if users == nil {
				return fmt.Errorf("expected users but got %v", users)
			}

			return err
		})

		if err != nil {
			t.Fatal(err)
		}
	})

	if err := pact.WritePact(); err != nil {
		t.Fatal(err)
	}

	publisher := dsl.Publisher{}

	err := publisher.Publish(types.PublishRequest{

		PactURLs:        []string{"./pacts/"},
		PactBroker:      "https://shayanadc.pactflow.io",
		BrokerToken:     "v_kaCEiuq-RpacJQflqV4g",
		ConsumerVersion: "1.0.0",
		Tags:            []string{"1.0.0", "latest"},
	})

	if err != nil {
		t.Fatal(err)
	}

	pact.Teardown()
}
