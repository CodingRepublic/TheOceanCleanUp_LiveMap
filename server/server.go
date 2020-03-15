package main

type Config struct {
	Database *DatabaseConfig `yaml:"db"`
}

type Server struct {
	Database *Database
}

func CreateServer(config *Config) (*Server, error) {
	var err error
	server := Server{}

	server.Database, err = InitDatabase(config.Database)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (s *Server) Run() error {
	return nil
}
