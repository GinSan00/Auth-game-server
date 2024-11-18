package app

import (
	"log"
	"main/internal/config"
)

type serviceProvider struct {
	config         config.Config
	authRepository AuthRepository
	authService    AuthService
	authImpl       AuthImplementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() config.Config {
	if s.Config == nil {
		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) AuthRepository() AuthRepository {
	if s.authRepository == nil {
		s.authRepository = NewAuthRepository()
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService() AuthService {
	if s.authService == nil {
		s.authService = NewAuthService(s.AuthRepository())
	}

	return s.authService
}

func (s *serviceProvider) AuthImplementation() AuthImplementation {
	if s.authImpl == nil {
		s.authImpl = NewImplementation(s.AuthService())
	}

	return s.authImpl
}
