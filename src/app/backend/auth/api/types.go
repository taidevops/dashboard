package api

type AuthenticationModes map[*AuthenticationMode]bool

// AuthenticationMode represents auth mode supported by dashboard, i.e. basic.
type AuthenticationMode string

// String returns string representation of auth mode.
func (self AuthenticationMode) String() string {
	return string(self)
}

const (
	Token AuthenticationMode = "token"
	Basic AuthenticationMode = "basic"
)

type AuthManager interface {
	// Login authentications user based on provided LoginSpec and returns AuthReponse. AuthResponse contains
	// generated token and list of non-critical errors such as 'Failed authentication'.
	Login(*LoginSpec) (*AuthResponse, error)

	Refresh(string) (string, error)

	AuthenticationModes() []AuthenticationMode

	AuthenticationSkippable() bool
}

type LoginSpec struct {
	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`

	Token string `json:"token,omitempty"`

	KubeConfig string `json:"kubeconfig,omitempty"`
}

type AuthResponse struct {
	Name string `json:"name,omitempty"`

	JWEToken string `json:"jwtToken"`

	Errors []error `json:"errors"`
}

type TokenRefreshSpec struct {
	JWEToken string `json:"jweToken"`
}

type LoginSkippableResponse struct {
	Skippable bool `json:"skippable"`
}
