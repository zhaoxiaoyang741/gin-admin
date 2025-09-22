package config

type Config struct {
	System System `mapstructure:"system" yaml:"system" json:"system"`
	Zap    Zap    `mapstructure:"zap" yaml:"zap" json:"zap"`
	Db     Db     `mapstructure:"db" json:"db" yaml:"db"`
}
