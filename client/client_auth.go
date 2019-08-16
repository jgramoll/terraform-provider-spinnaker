package client

// Auth for login on spinnaker
type Auth struct {
	Enabled   bool
	CertPath  string
	KeyPath   string
	UserEmail string
	Insecure  bool
}

func NewAuth() *Auth {
	return &Auth{
		Enabled:  true,
		Insecure: true,
	}
}
