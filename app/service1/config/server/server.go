package server

// ServerConfig hold http/grpc server config
type ServerConfig struct {
	GRPC Address `json:"grpc" mapstructure:"grpc" yaml:"grpc"`
	HTTP Address `json:"http" mapstructure:"http" yaml:"http"`
}

// Address represents a network end point address.
type Address struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

// DefaultServerConfig return a default server config
func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		GRPC: Address{
			Host: "0.0.0.0",
			Port: 10443,
		},
		HTTP: Address{
			Host: "0.0.0.0",
			Port: 10080,
		},
	}
}
