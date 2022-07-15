module wartech-studio.com/monster-reacher/gateway

go 1.18

require (
	github.com/gorilla/mux v1.8.0
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
	wartech-studio.com/monster-reacher/libraries v1.0.0
)

require github.com/srikrsna/protoc-gen-gotag v0.6.2 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220630215102-69896b714898 // indirect
	golang.org/x/oauth2 v0.0.0-20220630143837-2104d58473e0
	golang.org/x/sys v0.0.0-20220702020025-31831981b65f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220630174209-ad1d48641aa7 // indirect
)

replace wartech-studio.com/monster-reacher/libraries => ../libraries
