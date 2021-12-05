package config

type Server struct {
    HTTP HTTP `yaml:"http"`
}

type HTTP struct {
    IP string `yaml:"ip"`
    Port int `yaml:"port"`
}