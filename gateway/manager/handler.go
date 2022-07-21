package manager

import (
	"wartech-studio.com/monster-reacher/gateway/bff"
	"wartech-studio.com/monster-reacher/libraries/protobuf/gateway"
)

func NewServerService() gateway.GatewayServer {
	service := gateway.NewGatewayServer()
	service.AuthenticationHandler = bff.Authentication
	service.WartechRegisterHandler = bff.WartechRegister
	service.LinkServiceToAccountHandler = bff.LinkServiceToAccount
	service.GetProfileDataHandler = bff.GetProfileData
	service.GetCharacterDataHandler = bff.GetCharacterData
	service.SetCharacterNameHandler = bff.SetCharacterName
	service.SetCharacterMMRHandler = bff.SetCharacterMMR
	service.IncrementCharacterEXPHandler = bff.IncrementCharacterEXP
	return service
}
