package config

// this is just a sample. in production, should put into kubernetes secret / .env file
const (
	AuthBasicUsername        = "user"
	AuthBasicPassword        = "pass"
	RedisHost                = "redis:6379"
	RedisUsername            = ""
	RedisPassword            = "redispass"
	RedisDatabase            = 0
	PostgresStringConnection = "postgresql://postgres:@postgres:5434/golang_shorturl_test"
	PostgresDB               = "golang_shorturl_test"
	PostgresUser             = "postgres"
	//	PostgresPassword         = "12345678"
	PostgresPassword = "password"
	PostgressPort    = "5432"
	PostgresSSLMode  = "disable"
	PostgresHost     = "postgres"
)
