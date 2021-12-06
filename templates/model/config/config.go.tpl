package config

type Server struct {
    HTTP HTTP `yaml:"http"`
}

type HTTP struct {
    IP string `yaml:"ip"`
    Port int `yaml:"port"`
}

// Log ...
type Log struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	Prefix     string `yaml:"prefix"`
	Director   string `yaml:"director"`
	ShowLine   bool   `yaml:"showLine"`
	Color      bool   `yaml:"color"`
	Stacktrace string `yaml:"stacktrace"`
}