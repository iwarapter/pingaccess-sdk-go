package pingaccess

import (
	"net/http"
)

type Config struct {
	Username string
	Password string
	BaseURL  string
	Client   http.Client
}
