package fridge

import (
	"crypto/tls"
	"github.com/shomali11/xredis"
	"time"
)

// RedisOption an option for a storage
type RedisOption func(*RedisSettings)

// WithHost sets redis host
func WithHost(host string) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.Host = host
	}
}

// WithPort sets redis port
func WithPort(port int) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.Port = port
	}
}

// WithPassword sets redis password
func WithPassword(password string) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.Password = password
	}
}

// WithDatabase sets redis database
func WithDatabase(database int) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.Database = database
	}
}

// WithNetwork sets redis network
func WithNetwork(network string) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.Network = network
	}
}

// WithConnectTimeout sets redis connect timeout
func WithConnectTimeout(connectTimeout time.Duration) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.ConnectTimeout = connectTimeout
	}
}

// WithWriteTimeout sets redis write timeout
func WithWriteTimeout(writeTimeout time.Duration) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.WriteTimeout = writeTimeout
	}
}

// WithReadTimeout sets redis read timeout
func WithReadTimeout(readTimeout time.Duration) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.ReadTimeout = readTimeout
	}
}

// WithConnectionIdleTimeout sets redis connection idle timeout
func WithConnectionIdleTimeout(connectionIdleTimeout time.Duration) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.ConnectionIdleTimeout = connectionIdleTimeout
	}
}

// WithConnectionMaxIdle sets redis connection max idle
func WithConnectionMaxIdle(connectionMaxIdle int) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.ConnectionMaxIdle = connectionMaxIdle
	}
}

// WithConnectionMaxActive sets redis connection max active
func WithConnectionMaxActive(connectionMaxActive int) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.ConnectionMaxActive = connectionMaxActive
	}
}

// WithConnectionWait sets redis connection wait
func WithConnectionWait(connectionWait bool) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.ConnectionWait = connectionWait
	}
}

// WithTlsConfig sets redis tls config
func WithTlsConfig(tlsConfig *tls.Config) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.TlsConfig = tlsConfig
	}
}

// WithTlsSkipVerify sets redis tls skip verification
func WithTlsSkipVerify(tlsSkipVerify bool) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.TlsSkipVerify = tlsSkipVerify
	}
}

// WithTestOnBorrowPeriod sets redis test on borrow period
func WithTestOnBorrowPeriod(testOnBorrowPeriod time.Duration) RedisOption {
	return func(redisSettings *RedisSettings) {
		redisSettings.TestOnBorrowPeriod = testOnBorrowPeriod
	}
}

// RedisSettings contains redis settings
type RedisSettings struct {
	Host                  string
	Port                  int
	Password              string
	Database              int
	Network               string
	ConnectTimeout        time.Duration
	WriteTimeout          time.Duration
	ReadTimeout           time.Duration
	ConnectionIdleTimeout time.Duration
	ConnectionMaxIdle     int
	ConnectionMaxActive   int
	ConnectionWait        bool
	TlsConfig             *tls.Config
	TlsSkipVerify         bool
	TestOnBorrowPeriod    time.Duration
}

// RedisClient contains redis client
type RedisClient struct {
	xredisClient *xredis.Client
}

func NewRedisClient(options ...RedisOption) *RedisClient {
	settings := &RedisSettings{}
	for _, option := range options {
		option(settings)
	}

	xredisOptions := &xredis.Options{
		Host:                  settings.Host,
		Port:                  settings.Port,
		Password:              settings.Password,
		Database:              settings.Database,
		Network:               settings.Network,
		ConnectTimeout:        settings.ConnectTimeout,
		WriteTimeout:          settings.WriteTimeout,
		ReadTimeout:           settings.ReadTimeout,
		ConnectionIdleTimeout: settings.ConnectionIdleTimeout,
		ConnectionMaxIdle:     settings.ConnectionMaxIdle,
		ConnectionMaxActive:   settings.ConnectionMaxActive,
		ConnectionWait:        settings.ConnectionWait,
		TlsConfig:             settings.TlsConfig,
		TlsSkipVerify:         settings.TlsSkipVerify,
		TestOnBorrowPeriod:    settings.TestOnBorrowPeriod,
	}
	return &RedisClient{xredisClient: xredis.SetupClient(xredisOptions)}
}
