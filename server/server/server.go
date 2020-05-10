package server

type Config struct {
	Database *DatabaseConfig `yaml:"db"`
	Handler  *HandlerConfig  `yaml:"handler"`
}

func Init(config *Config) error {
	var err error

	err = InitDatabase(config.Database)
	if err != nil {
		return err
	}

	err = InitHandler(config.Handler)
	if err != nil {
		return err
	}
	return nil
}

func Start() {
	myHandler.Start()
}
