package app

import (
	"log"
	"main/internal/api/auth"
	"main/internal/config"
	"main/internal/repository"
	authRepository "main/internal/repository/auth/postgresql"
	"main/internal/service"
	authService "main/internal/service/auth"
)

type serviceProvider struct {
	config         config.Config
	authRepository repository.AuthRepository
	authService    service.AuthService
	authImpl       *auth.Implementation
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

		s.config = *cfg
	}

	return s.config
}

func (s *serviceProvider) AuthRepository() repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.New(s.config.ConnectionString)
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService() service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository())
	}

	return s.authService
}

func (s *serviceProvider) AuthImplementation() *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService())
	}

	return s.authImpl
}
