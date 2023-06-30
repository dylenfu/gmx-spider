package config

type DBConfig struct {
	URL      string
	User     string
	Password string
	Scheme   string
	Debug    bool
}
