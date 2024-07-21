package config

// this is just a sample. in production, should put into kubernetes secret / .env file
const (
	AuthBasicUsername = "user"
	AuthBasicPassword = "pass"
	RedisHost         = "localhost:6379"
	RedisUsername     = ""
	RedisPassword     = ""
	RedisDatabase     = 0
)
