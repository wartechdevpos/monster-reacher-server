package wartech

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRegister(t *testing.T) {
	server := NewWartechUserServer()
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()
	res, err := server.Register(ctx, &RegisterRequest{User: "test", Email: "test@email.com", Password: "password", BirthdayTimestamp: timestamppb.Now()})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(res)
}
