package config

import "time"

// Config of application
type Config struct {
	Server       Server          `mapstructure:"server,omitempty"`
	Swagger      Swagger         `mapstructure:"swagger,omitempty"`
	Http         Http            `mapstructure:"http,omitempty"`
	Logger       Logger          `mapstructure:"logger,omitempty"`
	Postgresql   Postgresql      `mapstructure:"postgresql,omitempty"`
	Mysql        Mysql           `mapstructure:"mysql,omitempty"`
	MongoDB      MongoDB         `mapstructure:"mongodb,omitempty"`
	Redis        Redis           `mapstructure:"redis,omitempty"`
	Sentinel     Sentinel        `mapstructure:"sentinel,omitempty"`
	Clickhouse   Clickhouse      `mapstructure:"clickhouse,omitempty"`
	Firestore    Firestore       `mapstructure:"firestore,omitempty"`
	Supabase     Supabase        `mapstructure:"supabase,omitempty"`
	Binance      ProviderBinance `mapstructure:"binance,omitempty"`
	Kraken       ProviderKraken  `mapstructure:"kraken,omitempty"`
	FearandGreed FearAndGreed    `mapstructure:"fearandgreed,omitempty"`
	ApiKeys      ApiKeys         `mapstructure:"apikeys,omitempty"`
	ApiUris      ApiUris         `mapstructure:"apiuris,omitempty"`
	Jobs         Jobs            `mapstructure:"jobs,omitempty"`
	Nats         Nats            `mapstructure:"nats,omitempty"`
	Calculation  Calculation     `mapstructure:"calculation,omitempty"`
}

// Swagger config
type Swagger struct {
	SWAGGER_BASIC_AUTH_USERNAME string `mapstructure:"SWAGGER_BASIC_AUTH_USERNAME,omitempty"`
	SWAGGER_BASIC_AUTH_PASSWORD string `mapstructure:"SWAGGER_BASIC_AUTH_PASSWORD,omitempty"`
}

// Server config
type Server struct {
	PROJECT_NAME          string        `mapstructure:"PROJECT_NAME,omitempty"`
	SERVICE_NAME          string        `mapstructure:"SERVICE_NAME,omitempty"`
	APP_ENV               string        `mapstructure:"APP_ENV,omitempty"`
	APP_DEBUG             bool          `mapstructure:"APP_DEBUG,omitempty"`
	APP_PROXY             bool          `mapstructure:"APP_PROXY,omitempty"`
	APP_PROXY_HOST        string        `mapstructure:"APP_PROXY_HOST,omitempty"`
	HTTP_PORT             int           `mapstructure:"HTTP_PORT,omitempty"`
	GRPC_PORT             int           `mapstructure:"GRPC_PORT,omitempty"`
	TIMEOUT               int           `mapstructure:"TIMEOUT,omitempty"`
	APP_SECRET            string        `mapstructure:"APP_SECRET,omitempty"`
	JWT_TOKEN_EXPIRE_TIME int           `mapstructure:"JWT_TOKEN_EXPIRE_TIME,omitempty"`
	APP_VERSION           string        `mapstructure:"APP_VERSION,omitempty"`
	READ_TIMEOUT          time.Duration `mapstructure:"READ_TIMEOUT,omitempty"`
	WRITE_TIMEOUT         time.Duration `mapstructure:"WRITE_TIMEOUT,omitempty"`
	MAX_CONN_IDLE         time.Duration `mapstructure:"MAX_CONN_IDLE,omitempty"`
	MAX_CONN_AGE          time.Duration `mapstructure:"MAX_CONN_AGE,omitempty"`
}

// Http config
type Http struct {
	PORT                string        `mapstructure:"PORT,omitempty"`
	PPROF_PORT          string        `mapstructure:"PPROF_PORT,omitempty"`
	TIMEOUT             time.Duration `mapstructure:"TIMEOUT,omitempty"`
	READ_TIMEOUT        time.Duration `mapstructure:"READ_TIMEOUT,omitempty"`
	WRITE_TIMEOUT       time.Duration `mapstructure:"WRITE_TIMEOUT,omitempty"`
	COOKIE_LIFE_TIME    int           `mapstructure:"COOKIE_LIFE_TIME,omitempty"`
	SESSION_COOKIE_NAME string        `mapstructure:"SESSION_COOKIE_NAME,omitempty"`
	SSL_CERT_PATH       string        `mapstructure:"SSL_CERT_PATH,omitempty"`
	SSL_CERT_KEY        string        `mapstructure:"SSL_CERT_KEY,omitempty"`
}

// Logger config
type Logger struct {
	DISABLE_CALLER     bool   `mapstructure:"DISABLE_CALLER,omitempty"`
	DISABLE_STACKTRACE bool   `mapstructure:"DISABLE_STACKTRACE,omitempty"`
	ENCODING           string `mapstructure:"ENCODING,omitempty"`
	LEVEL              string `mapstructure:"LEVEL,omitempty"`
}

// Postgresql config
type Postgresql struct {
	HOST       string `mapstructure:"HOST,omitempty"`
	PORT       int    `mapstructure:"PORT,omitempty"`
	USER       string `mapstructure:"USER,omitempty"`
	PASS       string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB string `mapstructure:"DEFAULT_DB,omitempty"`
	MAX_CONN   int    `mapstructure:"MAX_CONN,omitempty"`
	DRIVER     string `mapstructure:"DRIVER,omitempty"`
	SCHEMA     string `mapstructure:"SCHEMA,omitempty"`
	LOGGER     bool   `mapstructure:"LOGGER,omitempty"`
}

// Mysql config
type Mysql struct {
	HOST       string `mapstructure:"HOST,omitempty"`
	PORT       int    `mapstructure:"PORT,omitempty"`
	USER       string `mapstructure:"USER,omitempty"`
	PASS       string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB string `mapstructure:"DEFAULT_DB,omitempty"`
	MAX_CONN   int    `mapstructure:"MAX_CONN,omitempty"`
}

// MongoDB config
type MongoDB struct {
	HOST           string `mapstructure:"HOST,omitempty"`
	PORT           int    `mapstructure:"PORT,omitempty"`
	USER           string `mapstructure:"USER,omitempty"`
	PASS           string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB     string `mapstructure:"DEFAULT_DB,omitempty"`
	MONGO_DB_ATLAS string `mapstructure:"MONGO_DB_ATLAS,omitempty"`
}

// Redis config
type Redis struct {
	HOST          string `mapstructure:"HOST,omitempty"`
	PORT          int    `mapstructure:"PORT,omitempty"`
	USER          string `mapstructure:"USER,omitempty"`
	PASS          string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB    int    `mapstructure:"DEFAULT_DB,omitempty"`
	MIN_IDLE_CONN int    `mapstructure:"MIN_IDLE_CONN,omitempty"`
	POOL_SIZE     int    `mapstructure:"POOL_SIZE,omitempty"`
	POOL_TIMEOUT  int    `mapstructure:"POOL_TIMEOUT,omitempty"`
}

// Sentinel config
type Sentinel struct {
	HOST          string `mapstructure:"HOST,omitempty"`
	PORT          int    `mapstructure:"PORT,omitempty"`
	USER          string `mapstructure:"USER,omitempty"`
	PASS          string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB    int    `mapstructure:"DEFAULT_DB,omitempty"`
	MIN_IDLE_CONN int    `mapstructure:"MIN_IDLE_CONN,omitempty"`
	POOL_SIZE     int    `mapstructure:"POOL_SIZE,omitempty"`
	POOL_TIMEOUT  int    `mapstructure:"POOL_TIMEOUT,omitempty"`
}

// Clickhouse config
type Clickhouse struct {
	HOST       string `mapstructure:"HOST,omitempty"`
	PORT       int    `mapstructure:"PORT,omitempty"`
	USER       string `mapstructure:"USER,omitempty"`
	PASS       string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB string `mapstructure:"DEFAULT_DB,omitempty"`
}

// Firestore config
type Firestore struct {
	PROJECT_ID        string `mapstructure:"PROJECT_ID,omitempty"`
	DEFULT_COLLECTION string `mapstructure:"DEFULT_COLLECTION,omitempty"`
	CREDENTIALS_PATH  string `mapstructure:"CREDENTIALS_PATH,omitempty"`
}

// ProviderBinance Binance config
type ProviderBinance struct {
	API_URL                     string `mapstructure:"API_URL,omitempty"`
	WS_URL_SPOT                 string `mapstructure:"WS_URL_SPOT,omitempty"`
	WS_URL_FUTURE               string `mapstructure:"WS_URL_FUTURE,omitempty"`
	FAPI_URL                    string `mapstructure:"FAPI_URL,omitempty"`
	FTAKER_LONG_SHORT_RATIO_URL string `mapstructure:"FTAKER_LONG_SHORT_RATIO_URL,omitempty"`
	API_KEY                     string `mapstructure:"API_KEY,omitempty"`
	API_SECRET                  string `mapstructure:"API_SECRET,omitempty"`
	API_VER                     string `mapstructure:"API_VER,omitempty"`
	FAPI_VER                    string `mapstructure:"FAPI_VER,omitempty"`
}

// Supabase config
type Supabase struct {
	API_URL        string `mapstructure:"API_URL,omitempty"`
	API_KEY        string `mapstructure:"API_KEY,omitempty"`
	AUTH_JWT       string `mapstructure:"AUTH_JWT,omitempty"`
	REST_PATH      string `mapstructure:"REST_PATH,omitempty"`
	HOST           string `mapstructure:"HOST,omitempty"`
	PORT           int    `mapstructure:"PORT,omitempty"`
	USER           string `mapstructure:"USER,omitempty"`
	PASS           string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB     string `mapstructure:"DEFAULT_DB,omitempty"`
	DEFAULT_SCHEMA string `mapstructure:"DEFAULT_SCHEMA,omitempty"`
	MAX_CONN       int    `mapstructure:"MAX_CONN,omitempty"`
	MAX_IDLE_CONN  int    `mapstructure:"MAX_IDLE_CONN,omitempty"`
	DRIVER         string `mapstructure:"DRIVER,omitempty"`
}

// Binance config
type ProviderKraken struct {
	API_URL    string `mapstructure:"API_URL,omitempty"`
	API_KEY    string `mapstructure:"API_KEY,omitempty"`
	API_SECRET string `mapstructure:"API_SECRET,omitempty"`
	API_VER    string `mapstructure:"API_VER,omitempty"`
}

// FearAndGreed Fear and Greed Index config
type FearAndGreed struct {
	X_RAPIDAPI_KEY  string `mapstructure:"X_RAPIDAPI_KEY,omitempty"`
	X_RAPIDAPI_HOST string `mapstructure:"X_RAPIDAPI_HOST,omitempty"`
}

// ApiKeys specific API Keys
type ApiKeys struct {
	COIN_MARKETCAP_KEY string `mapstructure:"COIN_MARKETCAP_KEY,omitempty"`
	COIN_GLASS_KEY     string `mapstructure:"COIN_GLASS_KEY,omitempty"`
	CRYPTO_COMPARE_KEY string `mapstructure:"CRYPTO_COMPARE_KEY,omitempty"`
}

// ApiUris specific API URL's
type ApiUris struct {
	BLOCKCHAIN_CENTER_URI string `mapstructure:"BLOCKCHAIN_CENTER_URI,omitempty"`
}

// Jobs run intervals
type Jobs struct {
	INTERVAL_INDICATORS   time.Duration `mapstructure:"INTERVAL_INDICATORS,omitempty"`
	INTERVAL_FEARGREED    time.Duration `mapstructure:"INTERVAL_FEARGREED,omitempty"`
	INTERVAL_LONGSHORT    time.Duration `mapstructure:"INTERVAL_LONGSHORT,omitempty"`
	INTERVAL_SEASONINDEX  time.Duration `mapstructure:"INTERVAL_SEASONINDEX,omitempty"`
	INTERVAL_WORLDINDICES time.Duration `mapstructure:"INTERVAL_WORLDINDICES,omitempty"`
	INTERVAL_PRIVATEDATA  time.Duration `mapstructure:"INTERVAL_PRIVATEDATA,omitempty"`
	INTERVAL_WHALE_HUNTER time.Duration `mapstructure:"INTERVAL_PRIVATEDATA,omitempty"`
	INTERVAL_PRICE_TICKER time.Duration `mapstructure:"INTERVAL_PRIVATEDATA,omitempty"`
}

// Nats run intervals
type Nats struct {
	HOST         string `mapstructure:"HOST,omitempty"`
	CLIENT_PORT  int    `mapstructure:"CLIENT_PORT,omitempty"`
	CLUSTER_PORT int    `mapstructure:"CLUSTER_PORT,omitempty"`
	USER         string `mapstructure:"USER,omitempty"`
	PASS         string `mapstructure:"PASS,omitempty"`
}

// Calculation run intervals
type Calculation struct {
	INTERVAL_5m  time.Duration `mapstructure:"INTERVAL_5m,omitempty"`
	INTERVAL_15m time.Duration `mapstructure:"INTERVAL_15m,omitempty"`
	INTERVAL_30m time.Duration `mapstructure:"INTERVAL_30m,omitempty"`
	INTERVAL_45m time.Duration `mapstructure:"INTERVAL_45m,omitempty"`
	INTERVAL_1h  time.Duration `mapstructure:"INTERVAL_1h,omitempty"`
	INTERVAL_2h  time.Duration `mapstructure:"INTERVAL_2h,omitempty"`
	INTERVAL_4h  time.Duration `mapstructure:"INTERVAL_4h,omitempty"`
	INTERVAL_6h  time.Duration `mapstructure:"INTERVAL_6h,omitempty"`
	INTERVAL_8h  time.Duration `mapstructure:"INTERVAL_8h,omitempty"`
	INTERVAL_12h time.Duration `mapstructure:"INTERVAL_12h,omitempty"`
	INTERVAL_1d  time.Duration `mapstructure:"INTERVAL_1d,omitempty"`
	INTERVAL_1W  time.Duration `mapstructure:"INTERVAL_1W,omitempty"`
}

type IMTelegram struct {
	BOT_KEY string `mapstructure:"BOT_KEY,omitempty"`
	CHAT_ID string `mapstructure:"CHAT_ID,omitempty"`
}
