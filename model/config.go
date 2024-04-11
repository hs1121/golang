package model

type Config struct {
    DbURL    string `yaml:"dbUrl"`
    DbName   string `yaml:"dbName"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
}
