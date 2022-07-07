package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"wartech-studio.com/monster-reacher/gateway/api/authorization"
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
	enableCors(res)
	res.Write([]byte("This home of auth"))
}

type UserRegister struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"service_name"`
	ID       string `json:"service_id"`
	Token    string `json:"service_token"`
	Secret   string `json:"service_secret"`
}

func (*authApiHandle) register(res http.ResponseWriter, req *http.Request) {
	enableCors(res)
	if strings.ToLower(req.Method) != "post" {
		res.Write([]byte("use POST for register by user,password,email or service_name,service_id,service_token"))
		return
	}

	serivces, _ := ServicesDiscoveryCache.CheckRequireServices([]string{"authentication", "profile"})
	/*
		if !ok {
			res.Write([]byte(`{"success": false ,"message":"service profile,authentication is offline"}`))
			return
		}
	*/
	userRegister := UserRegister{}
	idRegister := ""
	var err error = nil

	if err = json.NewDecoder(req.Body).Decode(&userRegister); err == nil {
		if userRegister.User != "" {
			if idRegister, err = registerByUser(serivces, &userRegister); err != nil {
				res.Write([]byte(fmt.Sprintf(`{"success": false ,"message":"%s"}`, err.Error())))
				return
			}
		}
		if userRegister.Name != "" {
			if idRegister, err = registerByService(serivces, &userRegister); err != nil {
				res.Write([]byte(fmt.Sprintf(`{"success": false ,"message":"%s"}`, err.Error())))
				return
			}
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
	res.Write([]byte(fmt.Sprintf(`{"success": true ,"access_token":"%s" , "id":"%s"}`, resSignUp.GetAccessToken(), idRegister)))
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

func registerByService(serivces map[string]*services_discovery.ServiceInfo, service *UserRegister) (string, error) {

	if service.Name == "" || service.Token == "" {
		return "", fmt.Errorf("some a param is empty. please check params service_name,service_token")
	}

	var autho authorization.Authorization = nil

	switch strings.ToUpper(service.Name) {
	case authorization.SERVICE_MAME_GOOGLE:
		autho = authorization.NewAuthorizationGoogle(service.Token)
	case authorization.SERVICE_MAME_FACEBOOK:
		autho = authorization.NewAuthorizationFacebook(service.Token)
	case authorization.SERVICE_MAME_TWITTER:
		autho = authorization.NewAuthorizationTwitter(service.Token + "--" + service.Secret)
	case authorization.SERVICE_MAME_APPLE:
		autho = authorization.NewAuthorizationApple(service.Token)
	}

	if autho == nil {
		return "", errors.New("services " + service.Name + " not support")
	}

	if err := autho.SubmitAuth(); err != nil {
		return "", errors.New("token is expired")
	}

	if autho.GetData() == nil {
		return "", errors.New("user info is empty")
	}

	cc, err := grpc.Dial(serivces["profile"].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("serivces Dial is error %s", err.Error())
	}
	defer cc.Close()

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	c := profile.NewProfileClient(cc)

	result, err := c.ServiceIsValid(ctx, &profile.ServiceIsValidRequest{
		Name: autho.GetServiceName(),
		Id:   autho.GetData().ID,
	})

	if err != nil {
		return "", fmt.Errorf("serivces UserIsValid is error %s", err.Error())
	}

	if result.GetSuccess() {
		return "", fmt.Errorf("service %s id %s is exist", autho.GetServiceName(), autho.GetData().ID)
	}

	resultRegister, err := c.RegisterByService(ctx, &profile.RegisterByServiceRequest{
		Name: autho.GetServiceName(),
		Id:   autho.GetData().ID,
	})

	if err != nil {
		return "", fmt.Errorf("serivces Register is error %s", err.Error())
	}

	if resultRegister.GetId() == "" {
		return "", fmt.Errorf("service %s id %s register fail", autho.GetServiceName(), autho.GetData().ID)
	}

	return resultRegister.GetId(), nil
}
