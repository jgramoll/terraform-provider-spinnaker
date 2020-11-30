package client

// Auth for login on spinnaker
type Auth struct {
	CertPath    string
	CertContent string
	KeyPath     string
	KeyContent  string
	UserEmail   string
	Insecure    bool
}

// NewAuth new auth
func NewAuth() *Auth {
	return &Auth{
		Insecure: true,
	}
}
