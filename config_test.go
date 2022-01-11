package config 

import (
	"testing"
	"log"
)

func TestConfig(t *testing.T){
	cfg, err := read()
	if err != nil {
		log.Fatal("Error : config file", err)
	}
	if cfg.Environment.Env != "Development" {
		log.Fatal("Error : file found but inccorect env variable value : value : ", cfg.Environment.Env)
	}
}