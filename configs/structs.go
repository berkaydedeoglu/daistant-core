package configs

type Config struct {
	DB          *dbConfig          `koanf:"db"`
	Port        int                `koanf:"port"`
	GinMode     string             `koanf:"gin_mode"`
	GoogleOAuth *googleOAuthConfig `koanf:"google_oauth"`
}

type dbConfig struct {
	SQLiteFilePath string `koanf:"sqlite_file_path"`
}

type googleOAuthConfig struct {
	ClientID     string `koanf:"client_id"`
	ClientSecret string `koanf:"client_secret"`
	RedirectURL  string `koanf:"redirect_url"`
	Scope        string `koanf:"scope"`
	OAuthURL     string `koanf:"oauth_url"`
}
