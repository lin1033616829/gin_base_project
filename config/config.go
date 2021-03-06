package config

type ServConf struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
}

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	ReqLimit      bool   `mapstructure:"req-limit" json:"req-limit" yaml:"req-limit"`
	ReqLimitTimes int    `mapstructure:"req-limit-times" json:"req-limit-times" yaml:"req-limit-times"`
}

type Mysql struct {
	HostName    string `mapstructure:"hostname" json:"hostname" yaml:"hostname"`
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	Config      string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleCons int    `mapstructure:"max-idle-cons" json:"max-idle-cons" yaml:"max-idle-cons"`
	MaxOpenCons int    `mapstructure:"max-open-cons" json:"max-open-cons" yaml:"max-open-cons"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap      string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}
