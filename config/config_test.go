package config

import (
	"os"
	"testing"

	"github.com/luizvnasc/cwbus-hist/test"
)

func TestConfig(t *testing.T) {
	ec := EnvConfigurer{}

	t.Run("Obtendo conexão com o banco", func(t *testing.T) {
		want := os.Getenv("ONIBUSIO_DB_URL")
		got := ec.DBStrConn()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo nome do banco", func(t *testing.T) {
		want := os.Getenv("ONIBUSIO_DB_HIST")
		got := ec.DBName()
		test.AssertStringsEqual(t, want, got)
	})

}
