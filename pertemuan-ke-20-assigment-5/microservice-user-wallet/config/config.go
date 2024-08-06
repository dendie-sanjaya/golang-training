package config

// this is just a sample. in production, should put into kubernetes secret / .env file
const (
	AuthBasicUsername        = "user"
	AuthBasicPassword        = "pass"
	RedisHost                = "127.0.0.1:6378"
	RedisUsername            = ""
	RedisPassword            = ""
	RedisDatabase            = 0
	PostgresStringConnection = "postgresql://postgres:12345678@localhost:5433/golang_user_wallet?sslmode=disable"
	PostgresDB               = "golang_user_wallet"
	PostgresUser             = "postgres"
	//PostgresPassword         = "12345678"
	PostgresPassword = "password"
	PostgressPort    = "5434"
	PostgresSSLMode  = "disable"
	PostgresHost     = "127.0.0.1"
)
