package services

import (
	"github.com/sumansaurabh24/tb-employee-service/pkg/models"
	"github.com/sumansaurabh24/tb-employee-service/pkg/repositories"
	"go.uber.org/zap"
)

type Service[E models.Entity] interface {
	Create(model E) E
	FindById(id int) (*E, error)
	Update(id int, model E) (E, error)
	Delete(id int) error
}

// ServiceImpl Implementation of the base service
type ServiceImpl[E models.Entity] struct {
	// Add any required dependencies or database connection here
	logger   *zap.SugaredLogger
	database *repositories.Database[E]
}

// NewServiceImpl initialize new services
func NewServiceImpl[E models.Entity](logger *zap.SugaredLogger, database *repositories.Database[E]) ServiceImpl[E] {
	return ServiceImpl[E]{
		logger:   logger,
		database: database,
	}
}

// Create helps save/create a model using the provided database
func (s *ServiceImpl[E]) Create(model E) E {
	s.logger.With(zap.Any("entity to be created", model))
	return s.database.Insert(model)
}

// FindById helps in getting individual model by ID using the provided database
func (s *ServiceImpl[E]) FindById(id int) (*E, error) {
	s.logger.With(zap.Any("id", id)).Info("get by id")
	e, err := s.database.Find(id)
	return e, err
}

// Update help to update a model by ID using the provided database
func (s *ServiceImpl[E]) Update(id int, model E) (E, error) {
	s.logger.With(zap.Any("id", id)).Info("update by id")
	e, err := s.database.Update(id, model)
	return e, err
}

// Delete help in deleting a model by ID using the provided database
func (s *ServiceImpl[E]) Delete(id int) error {
	s.logger.With(zap.Any("id", id)).Info("delete by id")
	return s.database.Delete(id)
}
