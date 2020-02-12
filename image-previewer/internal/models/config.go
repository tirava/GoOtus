// Package models implements models.
package models

import (
	"fmt"
	"reflect"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// implementer's constants.
const (
	LogStd        = "logstd"
	LogRus        = "logrus"
	XDraw         = "xdraw"
	NfntCrop      = "nfnt_crop"
	MD5           = "md5"
	SHA1          = "sha1"
	SHA256        = "sha256"
	LRU           = "lru"
	NoLimit       = "nolimit"
	Disk          = "disk"
	InMemory      = "inmemory"
	maxCacheItems = 128
)

// Configer is the main interface for configs.
type Configer interface {
	GetConfig() Config
	SetConfig(Config)
}

// Config is the base config struct.
type Config struct {
	Source           string                 `yaml:"-"`
	Logger           string                 `yaml:"logger"`
	LogFile          string                 `yaml:"log_file"`
	LogLevel         string                 `yaml:"log_level"`
	ListenHTTP       string                 `yaml:"http_listen"`
	ListenPrometheus string                 `yaml:"prometheus_listen"`
	ListenPprof      string                 `yaml:"pprof_listen"`
	Previewer        string                 `yaml:"previewer"`
	Interpolation    entities.Interpolation `yaml:"interpolation"`
	NoProxyHeaders   []string               `yaml:"no_proxy_headers"`
	ImageURLEncoder  string                 `yaml:"img_url_encoder"`
	Cacher           string                 `yaml:"cache"`
	MaxCacheItems    int                    `yaml:"max_cache_items"`
	Storager         string                 `yaml:"storage"`
	StoragePath      string                 `yaml:"storage_path"`
}

// ConfigDefaults is for default parameters.
type ConfigDefaults map[string]interface{}

// GetDefaults returns default config fields.
func (c Config) GetDefaults() ConfigDefaults {
	config := make(ConfigDefaults)
	config["Source"] = ""
	config["Logger"] = LogStd
	config["LogFile"] = "previewer.log"
	config["LogLevel"] = "info"
	config["ListenHTTP"] = ":8080"
	config["ListenPrometheus"] = ":9180"
	config["ListenPprof"] = ":8181"
	config["Previewer"] = XDraw
	config["ImageURLEncoder"] = MD5
	config["Cacher"] = LRU
	config["MaxCacheItems"] = int64(maxCacheItems)
	config["Storager"] = Disk
	config["StoragePath"] = "cache"

	return config
}

// FillConfig fills empty config's fields.
func (c Config) FillConfig(inConfig Config) (Config, error) {
	defConfig := c.GetDefaults()

	v := reflect.ValueOf(inConfig)
	vPtr := reflect.ValueOf(&inConfig)
	configType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		name := configType.Field(i).Name

		if value == "" || value == 0 {
			switch v := defConfig[name].(type) {
			case string:
				vPtr.Elem().FieldByName(name).SetString(defConfig[name].(string))
			case int64:
				vPtr.Elem().FieldByName(name).SetInt(defConfig[name].(int64))
			default:
				return inConfig, fmt.Errorf("unexpected type '%T' while parsing config parameter '%s'",
					v, name)
			}
		}
	}

	if inConfig.MaxCacheItems < 0 {
		inConfig.MaxCacheItems = maxCacheItems
	}

	return inConfig, nil
}
