package item

import (
	"fmt"
	"github.com/shomali11/util/conversions"
	"github.com/shomali11/xredis"
	"time"
)

const (
	configKeyFormat = "%s.config"
)

// NewDao creates a new dao object
func NewDao(client *xredis.Client) *Dao {
	return &Dao{xredisClient: client}
}

// Dao controls access to redis
type Dao struct {
	xredisClient *xredis.Client
}

// Get retrieves an item
func (d *Dao) Get(key string) (string, bool, error) {
	return d.xredisClient.Get(key)
}

// Set stores a value
func (d *Dao) Set(key string, value string, timeout int64) error {
	_, err := d.xredisClient.SetEx(key, value, timeout)
	return err
}

// SetConfig stores a key's config
func (d *Dao) SetConfig(key string, config *Config) error {
	config.Timestamp = time.Now().UTC()
	timestampString, err := conversions.Stringify(config)
	if err != nil {
		return err
	}

	configKey := fmt.Sprintf(configKeyFormat, key)
	_, err = d.xredisClient.Set(configKey, timestampString)
	return err
}

// GetConfig retrieves a key's config
func (d *Dao) GetConfig(key string) (*Config, error) {
	configKey := fmt.Sprintf(configKeyFormat, key)
	configString, found, err := d.xredisClient.Get(configKey)
	if err != nil {
		return nil, err
	}

	if !found {
		return &Config{}, nil
	}

	var config *Config
	err = conversions.Structify(configString, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Remove an item
func (d *Dao) Remove(key string) error {
	timestampKey := fmt.Sprintf(configKeyFormat, key)
	_, err := d.xredisClient.Del(key, timestampKey)
	return err
}

// Ping pings redis
func (d *Dao) Ping() error {
	_, err := d.xredisClient.Ping()
	return err
}

// Close closes resources
func (d *Dao) Close() error {
	return d.xredisClient.Close()
}
