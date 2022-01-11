package config

/*
// DatabaseConfiguration is the required parameters to set up a DB instance
type DatabaseConfiguration struct {
	Name     string `default:"beauty"`
	Username string `default:"user"`
	Password string `default:"password"`
	Host     string `default:"localhost"`
	Port     string `default:"5432"`
	LogMode  bool   `default:"false"`
	SSLMode  string `default:"disable"`
}

// EmailConfiguration is the required parameters to send emails
type EmailConfiguration struct {
	Host     string `default:"smtp.beauty.io"`
	Port     string `default:"25"`
	Username string `default:"hello@beauty.io"`
	Password string `default:"password"`
	From     string `default:"hello@beauty.io"`
	Admin    string `default:"hello@beauty.io"`
}
*/

type EnvironmentConfiguration struct {
	Env string `default:"Development"`
}

type Configuration struct{
	Environment EnvironmentConfiguration
}