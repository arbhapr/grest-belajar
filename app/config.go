package app

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"grest.dev/grest"
)

// Config initializes the configuration (config) if it is not already initialized.
// If config is not initialized, it creates a new configUtil instance, configures it, and assigns it to config.
// It ensures that only one instance of config is created and reused.
func Config() {
	if config == nil {
		config = &configUtil{}
		config.configure()
		config.isConfigured = true
	}
}

// These variables represent various configuration settings used by the application.
// Each variable is assigned a default value or a value loaded from environment variables.
var (
	APP_VERSION = "2024.03.260947"

	APP_ENV  = "local"
	APP_PORT = "4001"
	APP_URL  = "http://localhost:4001"

	IS_MAIN_SERVER = true // set to true to run migration, seed and task scheduling

	IS_GENERATE_OPEN_API_DOC = false

	// for testing
	ENV_FILE            = ""
	IS_USE_MOCK_SERVICE = false
	IS_USE_MOCK_DB      = false

	LOG_CONSOLE_ENABLED     = true           // print log to the terminal
	LOG_FILE_ENABLED        = false          // log to a file. the fields below can be skipped if this value is false
	LOG_FILE_USE_LOCAL_TIME = true           // if false log rotation filename will be use UTC time
	LOG_FILE_FILENAME       = "logs/api.log" //
	LOG_FILE_MAX_SIZE       = 100            // MB
	LOG_FILE_MAX_AGE        = 7              // days
	LOG_FILE_MAX_BACKUPS    = 0              // files

	JWT_KEY     = "45845ccb526944ef8288337fafcccabd"
	CRYPTO_KEY  = "793a9474e8c3418fa7bbe39dd7d8f076"
	CRYPTO_SALT = "8decfa8093174ce7a4b194711a0d510b"
	CRYPTO_INFO = "info"

	DB_DRIVER            = "mysql"
	DB_HOST              = "127.0.0.1"
	DB_HOST_READ         = ""
	DB_PORT              = 5432
	DB_DATABASE          = "data.db"
	DB_USERNAME          = "mysql"
	DB_PASSWORD          = "secret"
	DB_MAX_OPEN_CONNS    = 0
	DB_MAX_IDLE_CONNS    = 5
	DB_CONN_MAX_LIFETIME = time.Hour // on .env = "1h". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	DB_IS_DEBUG          = false

	REDIS_HOST      = "127.0.0.1"
	REDIS_PORT      = "6379"
	REDIS_CACHE_DB  = 3
	REDIS_REPORT_DB = 3
	REDIS_USERNAME  = ""
	REDIS_PASSWORD  = ""

	FS_DRIVER          = "local"
	FS_LOCAL_DIR_PATH  = "storages"
	FS_PUBLIC_DIR_PATH = "storages"
	FS_END_POINT       = "s3.amazonaws.com"
	FS_PORT            = 443
	FS_REGION          = "ap-southeast-3"
	FS_BUCKET_NAME     = "attachments"
	FS_ACCESS_KEY      = ""
	FS_SECRET_KEY      = ""

	TELEGRAM_ALERT_TOKEN   = ""
	TELEGRAM_ALERT_USER_ID = ""
)

// config is a pointer to a configUtil instance.
// It is used to store and access the configuration settings.
var config *configUtil

// configUtil represents the application's configuration utility.
type configUtil struct {
	isConfigured bool
}

// configure configures the application's settings by loading values from environment variables using the grest.LoadEnv function.
// Each configuration setting is loaded from the corresponding environment variable and assigned to the appropriate variable.
// The godotenv package is used to load the .env file if provided.
func (*configUtil) configure() {

	// set ENV_FILE with absolute path for the .env file to run test with .env
	envFile := os.Getenv("ENV_FILE")
	if envFile != "" {
		godotenv.Load(envFile)
	} else {
		godotenv.Load()
	}

	grest.LoadEnv("APP_ENV", &APP_ENV)
	grest.LoadEnv("APP_PORT", &APP_PORT)
	grest.LoadEnv("APP_URL", &APP_URL)

	grest.LoadEnv("IS_MAIN_SERVER", &IS_MAIN_SERVER)

	grest.LoadEnv("ENV_FILE", &ENV_FILE)
	grest.LoadEnv("IS_USE_MOCK_SERVICE", &IS_USE_MOCK_SERVICE)
	grest.LoadEnv("IS_USE_MOCK_DB", &IS_USE_MOCK_DB)

	grest.LoadEnv("LOG_CONSOLE_ENABLED", &LOG_CONSOLE_ENABLED)
	grest.LoadEnv("LOG_FILE_ENABLED", &LOG_FILE_ENABLED)
	grest.LoadEnv("LOG_FILE_USE_LOCAL_TIME", &LOG_FILE_USE_LOCAL_TIME)
	grest.LoadEnv("LOG_FILE_FILENAME", &LOG_FILE_FILENAME)
	grest.LoadEnv("LOG_FILE_MAX_SIZE", &LOG_FILE_MAX_SIZE)
	grest.LoadEnv("LOG_FILE_MAX_AGE", &LOG_FILE_MAX_AGE)
	grest.LoadEnv("LOG_FILE_MAX_BACKUPS", &LOG_FILE_MAX_BACKUPS)

	grest.LoadEnv("JWT_KEY", &JWT_KEY)
	grest.LoadEnv("CRYPTO_KEY", &CRYPTO_KEY)
	grest.LoadEnv("CRYPTO_SALT", &CRYPTO_SALT)
	grest.LoadEnv("CRYPTO_INFO", &CRYPTO_INFO)

	grest.LoadEnv("DB_DRIVER", &DB_DRIVER)
	grest.LoadEnv("DB_HOST", &DB_HOST)
	grest.LoadEnv("DB_HOST_READ", &DB_HOST_READ)
	grest.LoadEnv("DB_PORT", &DB_PORT)
	grest.LoadEnv("DB_DATABASE", &DB_DATABASE)
	grest.LoadEnv("DB_USERNAME", &DB_USERNAME)
	grest.LoadEnv("DB_PASSWORD", &DB_PASSWORD)
	grest.LoadEnv("DB_MAX_OPEN_CONNS", &DB_MAX_OPEN_CONNS)
	grest.LoadEnv("DB_MAX_IDLE_CONNS", &DB_MAX_IDLE_CONNS)
	grest.LoadEnv("DB_CONN_MAX_LIFETIME", &DB_CONN_MAX_LIFETIME)
	grest.LoadEnv("DB_IS_DEBUG", &DB_IS_DEBUG)

	grest.LoadEnv("REDIS_HOST", &REDIS_HOST)
	grest.LoadEnv("REDIS_PORT", &REDIS_PORT)
	grest.LoadEnv("REDIS_CACHE_DB", &REDIS_CACHE_DB)
	grest.LoadEnv("REDIS_REPORT_DB", &REDIS_REPORT_DB)
	grest.LoadEnv("REDIS_USERNAME", &REDIS_USERNAME)
	grest.LoadEnv("REDIS_PASSWORD", &REDIS_PASSWORD)

	grest.LoadEnv("FS_END_POINT", &FS_END_POINT)
	grest.LoadEnv("FS_PORT", &FS_PORT)
	grest.LoadEnv("FS_REGION", &FS_REGION)
	grest.LoadEnv("FS_BUCKET_NAME", &FS_BUCKET_NAME)
	grest.LoadEnv("FS_ACCESS_KEY", &FS_ACCESS_KEY)
	grest.LoadEnv("FS_SECRET_KEY", &FS_SECRET_KEY)

	grest.LoadEnv("TELEGRAM_ALERT_TOKEN", &TELEGRAM_ALERT_TOKEN)
	grest.LoadEnv("TELEGRAM_ALERT_USER_ID", &TELEGRAM_ALERT_USER_ID)
}
