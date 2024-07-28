package config

// this is just a sample. in production, should put into kubernetes secret / .env file
const (
	AuthBasicUsername        = "user"
	AuthBasicPassword        = "pass"
	RedisHost                = "redis:6378"
	RedisUsername            = ""
	RedisPassword            = "redispass"
	RedisDatabase            = 0
	PostgresStringConnection = "postgresql://postgres:@postgres:5434/golang_shorturl_test"
	PostgresDB               = "golang_shorturl_test"
	PostgresUser             = "postgres"
	//	PostgresPassword         = "12345678"
	PostgresPassword = "password"
	PostgressPort    = "5434"
	PostgresSSLMode  = "disable"
	PostgresHost     = "172.22.0.4"
)
