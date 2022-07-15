package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/healthcheck"
	"wartech-studio.com/monster-reacher/libraries/protobuf/wartech"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

var SERVICES_NAME = config.GetNameConfig().MicroServiceName.Wartech

var listenHost = fmt.Sprintf("%s:%d",
	config.GetServiceConfig().Services[SERVICES_NAME].Hosts[0],
	config.GetServiceConfig().Services[SERVICES_NAME].Ports[0])

func main() {
	server := grpc.NewServer()
	healthchecker := healthcheck.NewHealthCheckClient()
	go healthchecker.Start(SERVICES_NAME, listenHost)
	go initOAuth2Server()
	listener, err := net.Listen("tcp", listenHost)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	service := wartech.NewWartechServer()

	wartech.RegisterWartechServer(server, service)
	//reflection.Register(server)
	log.Println("gRPC server listening on " + listenHost)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func initOAuth2Server() {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	if _, ok := config.GetOAuth2Config()["wartech"]; !ok {
		panic("not fount oauth2 config")
	}
	oauth2Config := config.GetOAuth2Config()["wartech"]
	clientStore.Set(oauth2Config.ClientId, &models.Client{
		ID:     oauth2Config.ClientId,
		Secret: oauth2Config.ClientSecret,
		Domain: oauth2Config.Endpoint.DomainUrl,
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetUserAuthorizationHandler(UserAuthorizationHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)

	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		token, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if res, err := wartech.NewWartechServer().GetUser(ctx, &wartech.GetUserRequest{Id: token.GetUserID()}); err == nil {
			data := map[string]interface{}{
				"id":    res.GetData().GetId(),
				"email": res.GetData().GetEmail(),
				"name":  res.GetData().GetUser(),
				"user":  res.GetData().GetUser(),
			}
			e := json.NewEncoder(w)
			e.SetIndent("", "  ")
			e.Encode(data)
			return
		}
		http.Error(w, "something was wrong", http.StatusBadRequest)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.GetServiceConfig().Services[SERVICES_NAME].Ports[1]), nil))
}

func UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	userOrEmail := r.FormValue("userOrEmail")
	password := r.FormValue("password")
	if userOrEmail == "" || password == "" {
		return "", errors.New("not found field userOrEmail and password")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	data, err := wartech.NewWartechServer().Authentication(ctx, &wartech.AuthenticationRequest{UserOrEmail: userOrEmail, Password: password})
	if err != nil {
		return "", err
	}

	return data.GetId(), nil
}
