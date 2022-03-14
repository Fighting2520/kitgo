package nacosconfig

import "github.com/nacos-group/nacos-sdk-go/common/constant"

type (
	Config struct {
		ServerConfig
		ClientConfig
	}

	ClientConfig struct {
		NamespaceId string
		CacheDir    string // nacos 读取的配置缓存目录
		Username    string // the username for nacos auth
		Password    string // the password for nacos auth
		LogDir      string //  nacos 服务日志目录
		TimeoutMs   uint64 // timeout for requesting Nacos server, default value is 10000ms
	}

	ServerConfig struct {
		Scheme      string //the nacos server scheme
		ContextPath string //the nacos server contextpath
		IpAddr      string //the nacos server address
		Port        uint64 //the nacos server port
	}
)

func NewConfig(opts ...Option) *Config {
	var conf Config
	for _, opt := range opts {
		opt(&conf)
	}
	return &conf
}

// Option ...
type Option func(*Config)

// WithTimeoutMs ...
func WithTimeoutMs(timeoutMs uint64) Option {
	return func(config *Config) {
		config.TimeoutMs = timeoutMs
	}
}

// WithNamespaceId ...
func WithNamespaceId(namespaceId string) Option {
	return func(config *Config) {
		config.NamespaceId = namespaceId
	}
}

// WithCacheDir ...
func WithCacheDir(cacheDir string) Option {
	return func(config *Config) {
		config.CacheDir = cacheDir
	}
}

// WithUsername ...
func WithUsername(username string) Option {
	return func(config *Config) {
		config.Username = username
	}
}

// WithPassword ...
func WithPassword(password string) Option {
	return func(config *Config) {
		config.Password = password
	}
}

// WithLogDir ...
func WithLogDir(logDir string) Option {
	return func(config *Config) {
		config.LogDir = logDir
	}
}

//WithScheme set Scheme for server
func WithScheme(scheme string) Option {
	return func(config *Config) {
		config.Scheme = scheme
	}
}

//WithContextPath set contextPath for server
func WithContextPath(contextPath string) Option {
	return func(config *Config) {
		config.ContextPath = contextPath
	}
}

//WithIpAddr set ip address for server
func WithIpAddr(ipAddr string) Option {
	return func(config *Config) {
		config.IpAddr = ipAddr
	}
}

//WithPort set port for server
func WithPort(port uint64) Option {
	return func(config *Config) {
		config.Port = port
	}
}

func convertServerOption(config Config) []constant.ServerOption {
	var serverOpts []constant.ServerOption
	if config.Scheme != "" {
		serverOpts = append(serverOpts, constant.WithScheme(config.Scheme))
	}
	if config.ContextPath != "" {
		serverOpts = append(serverOpts, constant.WithContextPath(config.ContextPath))
	}
	return serverOpts
}

func convertClientOption(config Config) []constant.ClientOption {
	var clientOpts []constant.ClientOption
	if config.Username != "" {
		clientOpts = append(clientOpts, constant.WithUsername(config.Username))
	}
	if config.Password != "" {
		clientOpts = append(clientOpts, constant.WithPassword(config.Password))
	}
	if config.TimeoutMs > 0 {
		clientOpts = append(clientOpts, constant.WithTimeoutMs(config.TimeoutMs))
	}
	if config.CacheDir != "" {
		clientOpts = append(clientOpts, constant.WithCacheDir(config.CacheDir))
	}
	if config.LogDir != "" {
		clientOpts = append(clientOpts, constant.WithLogDir(config.LogDir))
	}
	if config.NamespaceId != "" {
		clientOpts = append(clientOpts, constant.WithNamespaceId(config.NamespaceId))
	}
	return clientOpts
}
