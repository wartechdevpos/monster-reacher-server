package authorization

const (
	SERVICE_MAME_GOOGLE   = "GOOGLE"
	SERVICE_MAME_FACEBOOK = "FACEBOOK"
	SERVICE_MAME_TWITTER  = "TWITTER"
	SERVICE_MAME_APPLE    = "APPLE"
)

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Authorization interface {
	SubmitAuth() error
	// call SubmitAuth before
	GetData() *UserInfo
	GetServiceName() string
	GetVersion() string
}

type method struct{}

func (*method) SubmitAuth() error      { return nil }
func (*method) GetData() interface{}   { return nil }
func (*method) GetServiceName() string { return "" }
func (*method) GetVersion() string     { return "" }
