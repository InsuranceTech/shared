package flagger

import (
	"flag"
	"os"
)

func InitAppEnvFlag() {
	Filters := new(AppEnvFlagParams)
	flag.StringVar(&Filters.AppEnv, "APP_ENV", "dev", "--APP_ENV=prod")
	flag.Parse()
	os.Setenv("APP_ENV", Filters.AppEnv)
}
