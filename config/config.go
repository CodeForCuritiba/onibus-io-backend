// Package config obtém as configurações da app que estão nas variáveis de ambiente
package config

import "os"

const (
	prefix    = "ONIBUSIO_"
	dbStrConn = "DB_URL"
	dbName    = "DB_HIST"
	port      = "PORT"
)

// Configurer é a interface que define um configurador no sistema.
type Configurer interface {
	DBName() string
	DBStrConn() string
	Port() string
}

// EnvConfigurer é um confiurador que  capitura as configurações das variáveis de ambiente.
type EnvConfigurer struct{}

func (ec EnvConfigurer) key(name string) string {
	return prefix + name
}

func (ec EnvConfigurer) getValue(name string) string {
	return os.Getenv(ec.key(name))
}

// DBStrConn retorna a string de conexão do banco de dados.
func (ec EnvConfigurer) DBStrConn() string {
	return ec.getValue(dbStrConn)
}

// DBName retorna o nome do banco de dados.
func (ec EnvConfigurer) DBName() string {
	return ec.getValue(dbName)
}

// Port retorna a porta liberada para rodar o servidor (padrão heroku)
func (ec EnvConfigurer) Port() string {
	return os.Getenv(port)
}
