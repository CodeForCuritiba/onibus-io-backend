// Package config obtém as configurações da app que estão nas variáveis de ambiente
package config

//Configuration of the app
type Configuration struct {
	ServiceURL string `env:"CWBUS_URBS_SERVICE_URL"`
	UrbsCode   string `env:"CWBUS_URBS_CODE"`
	MongoDB    struct {
		StrConn string `env:"CWBUS_DB_URL"`
		DBName  string `env:"CWBUS_DB_HIST"`
	}
	Port string `env:"PORT"`
}
