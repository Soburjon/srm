package configs

import (
	"errors"
	"github.com/spf13/cast"
	"os"
	"sync"
)

var (
	instance *Configuration
	once     sync.Once

	//ErrExpiredPassword error text
	ErrExpiredPassword error = errors.New("expired_password")
)

// Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}

// Configuration ...
type Configuration struct {
	Environment      string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	ServerPort       int
	ServerHost       string

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int

	CasbinConfigPath    string
	MiddlewareRolesPath string

	ServerReadTimeout int
}

func load() *Configuration {
	return &Configuration{
		ServerHost:       cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:       cast.ToInt(getOrReturnDefault("SERVER_PORT", "8000")),
		Environment:      cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDatabase: cast.ToString(getOrReturnDefault("POSTGRES_DB", "srm")),
		PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres")),
		PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12345")),

		CasbinConfigPath:    cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./configs/rbac_model.conf")),
		MiddlewareRolesPath: cast.ToString(getOrReturnDefault("MIDDLEWARE_ROLES_PATH", "./configs/models.csv")),

		JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
		JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
		JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
		JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),

		ServerReadTimeout: cast.ToInt(getOrReturnDefault("SERVER_READ_TIMEOUT", "")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
