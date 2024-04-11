package go_http_client

type ClientConfig struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
	Timeout  int32  `json:"timeout" yaml:"timeout"`
	Debug    bool   `json:"debug" yaml:"debug"`
	Auth     *Auth  `json:"auth" yaml:"auth"`
}

type Auth struct {
	Scheme string `json:"scheme" yaml:"scheme"`
	Token  string `json:"token" yaml:"token"`
}
