package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"wartech-studio.com/monster-reacher/gateway/services/authentication"
	"wartech-studio.com/monster-reacher/gateway/services/profile"
	"wartech-studio.com/monster-reacher/gateway/services/services_discovery"
)

type authApiHandle struct{}

func RegisterAuthApiHandle(router *mux.Router) *authApiHandle {
	handler := &authApiHandle{}
	router.HandleFunc("/api/auth", handler.home)
	router.HandleFunc("/api/auth/register", handler.register)
	return handler
}

func (*authApiHandle) home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("This home of auth"))
}

type UserRegister struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type ServiceRegister struct {
	Name  string `json:"service_name,omitempty"`
	ID    string `json:"service_id,omitempty"`
	Token string `json:"service_token,omitempty"`
}

func (*authApiHandle) register(res http.ResponseWriter, req *http.Request) {

	if strings.ToLower(req.Method) != "post" {
		res.Write([]byte("use POST for register by user,password,email or service_name,service_id,service_token"))
		return
	}

	serivces, ok := ServicesDiscoveryCache.CheckRequireServices([]string{"authentication", "profile"})

	if !ok {
		res.Write([]byte(`{"success": false ,"message":"service profile,authentication is offline"}`))
		return
	}

	userRegister := UserRegister{}
	serviceRegister := ServiceRegister{}
	idRegister := ""
	var err error = nil

	if err = json.NewDecoder(req.Body).Decode(&userRegister); err == nil {
		if idRegister, err = registerByUser(serivces, &userRegister); err != nil {
			res.Write([]byte(fmt.Sprintf(`{"success": false ,"message":"%s"}`, err.Error())))
		}
	} else if err = json.NewDecoder(req.Body).Decode(&serviceRegister); err == nil {
		if idRegister, err = registerByService(serivces, &serviceRegister); err != nil {
			res.Write([]byte(fmt.Sprintf(`{"success": false ,"message":"%s"}`, err.Error())))
		}
	} else {
		res.Write([]byte(fmt.Sprintf(`{"success": false ,"message":"please check params user,password,email or service_name,service_id,service_token","error":%s}`, err.Error())))
		return
	}

	cc, err := grpc.Dial(serivces["authentication"].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"success": false ,"message":"serivces is error %s"}`, err.Error())))
		return
	}
	defer cc.Close()

	c := authentication.NewAuthenticationClient(cc)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	resSignUp, err := c.SignUp(ctx, &authentication.SignUpRequest{Id: idRegister})
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"success": false ,"message":"serivces is error %s"}`, err.Error())))
		return
	}
	res.Write([]byte(fmt.Sprintf(`{"success": true ,"access_token":%s , "id":%s}`, resSignUp.GetAccessToken(), idRegister)))
}

func registerByUser(serivces map[string]*services_discovery.ServiceInfo, user *UserRegister) (string, error) {

	if user.User == "" || user.Password == "" || user.Email == "" {
		return "", fmt.Errorf("some a param is empty. please check params user,password,email")
	}

	cc, err := grpc.Dial(serivces["profile"].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("serivces Dial is error %s", err.Error())
	}
	defer cc.Close()

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	c := profile.NewProfileClient(cc)

	result, err := c.UserIsValid(ctx, &profile.UserIsValidRequest{User: user.User})

	if err != nil {
		return "", fmt.Errorf("serivces UserIsValid is error %s", err.Error())
	}

	if result.GetSuccess() {
		return "", fmt.Errorf("user %s is exist", user.User)
	}

	resultRegister, err := c.Register(ctx, &profile.RegisterRequest{
		User:     user.User,
		Password: user.Password,
		Email:    user.Email,
	})

	if err != nil {
		return "", fmt.Errorf("serivces Register is error %s", err.Error())
	}

	if resultRegister.GetId() == "" {
		return "", fmt.Errorf("user %s register fail", user.User)
	}

	return resultRegister.GetId(), nil
}

func registerByService(serivces map[string]*services_discovery.ServiceInfo, service *ServiceRegister) (string, error) {
	return "", nil
}
