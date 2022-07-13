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

	"wartech-studio.com/monster-reacher/microservices/services/wartech"

	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/healthcheck"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

const SERVICES_NAME = "wartech"

var listenHost = fmt.Sprintf("%s:%d",
	config.WartechConfig().Services[SERVICES_NAME].Hosts[0],
	config.WartechConfig().Services[SERVICES_NAME].Ports[0])

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

	wartech.RegisterWartechUserServer(server, wartech.NewWartechUserServer())
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
	clientStore.Set("85462020023651", &models.Client{
		ID:     "85462020023651",
		Secret: "ac29c66a3bb016d2c632a3a7dc5130b",
		Domain: "https://insuanhouse.ddns.net:3000",
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
		res, _ := wartech.NewWartechUserServer().GetUser(ctx, &wartech.GetUserRequest{Id: token.GetUserID()})
		if userData, err := wartech.SelectizeData(res.Data); err == nil {
			data := map[string]interface{}{
				"id":    userData.ID,
				"email": userData.Email,
				"name":  userData.User,
			}
			e := json.NewEncoder(w)
			e.SetIndent("", "  ")
			e.Encode(data)
			return
		}
		http.Error(w, "something was wrong", http.StatusBadRequest)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.WartechConfig().Services[SERVICES_NAME].Ports[1]), nil))
}

func UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	userOrEmail := r.FormValue("userOrEmail")
	password := r.FormValue("password")
	if userOrEmail == "" || password == "" {
		return "", errors.New("not found field userOrEmail and password")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	data, err := wartech.NewWartechUserServer().Authentication(ctx, &wartech.AuthenticationRequest{UserOrEmail: userOrEmail, Password: password})
	if err != nil {
		return "", err
	}

	return data.GetId(), nil
}
