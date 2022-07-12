package wartech

import (
	context "context"
	"encoding/json"
	"errors"

	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/database"
)

type wartechUserServer struct{}

type wartechUserServerInfo struct {
	id              string
	createTimestamp int64
}

const NAME_DATABASE = "wartech"
const NAME_TABLE = "accounts"

type wartechUserDBSchema struct {
	ID             string `bson:"_id,omitempty"`
	User           string `bson:"user"`
	Email          string `bson:"email"`
	Password       string `bson:"password"`
	Birthday       string `bson:"birthday"`
	EmailConfirmed bool   `bson:"emailConfirmed"`
}

func NewWartechUserServer() WartechUserServer {
	return &wartechUserServer{}
}

func (*wartechUserServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	id, err := internalRegister(ctx, req.GetUser(), req.GetEmail(), req.GetPassword(), req.GetBirthdayTimestamp().String())
	return &RegisterResponse{Id: id}, err
}
func (w *wartechUserServer) Authentication(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	data := internalAuthentication(ctx, req.GetUserOrEmail(), req.GetPassword())
	if data == nil {
		return nil, errors.New("user not found or password mismatch")
	}
	return &AuthenticationResponse{IsSuccess: true, IsConfirmed: data.EmailConfirmed, Id: data.ID}, nil
}
func (*wartechUserServer) CheckUserOrEmail(ctx context.Context, req *CheckUserOrEmailRequest) (*CheckUserOrEmailResponse, error) {
	if _, found := internalCheckUser(ctx, req.GetUserOrEmail()); !found {
		if _, found := internalCheckEmail(ctx, req.GetUserOrEmail()); !found {
			return &CheckUserOrEmailResponse{IsValid: false}, nil
		}
	}
	return &CheckUserOrEmailResponse{IsValid: true}, nil
}
func (*wartechUserServer) ForgottenPassword(context.Context, *ForgottenPasswordRequest) (*ForgottenPasswordResponse, error) {
	return nil, nil
}
func (*wartechUserServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, nil
}
func (w *wartechUserServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	data := internalGetUser(ctx, req.GetId())
	bytesData, err := json.Marshal(data)
	return &GetUserResponse{Data: bytesData}, err
}
func (*wartechUserServer) ConfirmEmail(context.Context, *ConfirmEmailRequest) (*ConfirmEmailResponse, error) {
	return nil, nil
}
func (*wartechUserServer) mustEmbedUnimplementedWartechUserServer() {}

func getDriver() database.DBDriver {
	driver, err := database.NewMongoDBDriver(config.WartechConfig().Databases["mongodb"].Host, NAME_DATABASE, NAME_TABLE)

	if err != nil {
		panic(err)
	}
	return driver
}

func internalRegister(ctx context.Context, user string, email string, password string, birthday string) (string, error) {
	driver := getDriver()
	defer driver.Close()
	data := &wartechUserDBSchema{
		User:     user,
		Email:    email,
		Password: password,
		Birthday: birthday,
	}
	result, err := driver.PushOne(ctx, data)
	if err != nil {
		return "", err
	}
	return database.MongoDBDecodeResultToID(result), nil
}

func internalCheckUser(ctx context.Context, user string) (string, bool) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("user", user)
	data, err := getwartechUserData(ctx, driver, filter)
	if err != nil {
		return "", false
	}
	return data.ID, true
}

func internalCheckEmail(ctx context.Context, email string) (string, bool) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("email", email)
	data, err := getwartechUserData(ctx, driver, filter)
	if err != nil {
		return "", false
	}
	return data.ID, true
}

func internalGetUser(ctx context.Context, id string) *wartechUserDBSchema {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", id)
	data, err := getwartechUserData(ctx, driver, filter)
	if err != nil {
		return nil
	}
	return data
}

func internalAuthentication(ctx context.Context, userOrEmail string, password string) *wartechUserDBSchema {
	id := ""
	found := false

	if id, found = internalCheckUser(ctx, userOrEmail); !found {
		if id, found = internalCheckEmail(ctx, userOrEmail); !found {
			return nil
		}
	}
	data := internalGetUser(ctx, id)

	if data.Password != password {
		return nil
	}

	return data
}

func getwartechUserData(ctx context.Context, driver database.DBDriver, filter interface{}) (*wartechUserDBSchema, error) {
	result := driver.SelectOne(ctx, filter)
	if err := database.MongoDBSelectOneResultGetError(result); err != nil {
		return nil, err
	}
	data := &wartechUserDBSchema{}
	if err := database.MongoDBDecodeResultToStruct(result, data); err != nil {
		return nil, nil
	}

	return data, nil
}

func SelectizeData(data []byte) (*wartechUserDBSchema, error) {
	w := &wartechUserDBSchema{}
	err := json.Unmarshal(data, w)
	return w, err
}
