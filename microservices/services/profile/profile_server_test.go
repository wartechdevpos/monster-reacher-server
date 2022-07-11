package profile

import (
	"context"
	"fmt"
	"testing"
)

func TestGetData(t *testing.T) {

}

func TestAuthentication(t *testing.T) {

}

func TestAuthenticationByService(t *testing.T) {
	server := NewProfileServer()
	res, err := server.AuthenticationByService(context.Background(), &AuthenticationByServiceRequest{
		Name: "google",
		Id:   "112170918475213245662",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestRegister(t *testing.T) {
	server := NewProfileServer()
	res, err := server.Register(context.Background(), &RegisterRequest{
		User:     "test",
		Password: "testpass",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(res.GetId())
}
