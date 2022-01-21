package config

type Server struct {
	Origins []string
	Host    string
	Port    int
}

type Auth struct {
	Domain   string
	ClientID string
	AuthKey  string
	Secret   string
}

type Db struct {
	Username string
	Password string
	Url      string
}
