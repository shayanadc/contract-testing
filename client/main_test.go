package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestClientPact_Local(t *testing.T) {

	pact := dsl.Pact{
		Consumer: "example-client",
		Provider: "example-server",
	}
	pact.Setup(true)

	t.Run("get users list", func(t *testing.T) {

		pact.
			AddInteraction().
			Given("user list").
			UponReceiving("User collection is requested").
			WithRequest(dsl.Request{
				Method: "GET",
				Path:   dsl.Term("/users", "/users*"),
			}).
			WillRespondWith(dsl.Response{
				Status: 200,
				Body: dsl.Like([]User{
					{
						ID:        "123",
						FirstName: "XYZ",
					},
					{
						ID:        "124",
						FirstName: "ABS",
					},
				}),
			})

		err := pact.Verify(func() error {

			host := fmt.Sprintf("%s:%d", pact.Host, pact.Server.Port)

			users, err := GetUsers(host)

			if err != nil {
				return errors.New("error is not expected")
			}
			if users == nil {
				return fmt.Errorf("expected user with ID %s but got %v", "12345", users)
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

	pact.Teardown()
}
