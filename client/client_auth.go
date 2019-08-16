package client

// Auth for login on spinnaker
type Auth struct {
	Enabled     bool
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
		Enabled:  true,
		Insecure: true,
	}
}
