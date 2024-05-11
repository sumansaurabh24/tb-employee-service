package repositories

import (
	"github.com/sumansaurabh24/tb-employee-service/pkg/errors"
	"github.com/sumansaurabh24/tb-employee-service/pkg/models"
	"go.uber.org/zap"
	"sync"
)

// Database - struct to hold the entity
type Database[E models.Entity] struct {
	data    map[int]E
	lock    sync.RWMutex
	logger  *zap.SugaredLogger
	counter int
}

// NewDatabase - constructor
func NewDatabase[E models.Entity](logger *zap.SugaredLogger) Database[E] {
	return Database[E]{
		logger: logger,
		data:   make(map[int]E),
		lock:   sync.RWMutex{},
	}
}

// Find - find the entity by its key from data store
func (d *Database[E]) Find(key int) (*E, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	value, found := d.data[key]
	d.logger.With(
		zap.Any("value", value),
		zap.Any("key", key),
	).Info("db get")
	if !found {
		return nil, &errors.ErrEntityNotFound
	}
	return &value, nil
}

// Insert - insert the value into the database
func (d *Database[E]) Insert(value E) E {
	d.lock.Lock()
	defer func() {
		d.lock.Unlock()
	}()
	d.counter++
	copied := value.SetId(d.counter)
	d.data[d.counter] = copied.(E)
	d.logger.With(zap.Any("value", copied)).Info("db set")
	return copied.(E)
}

// Update - update value into the database
func (d *Database[E]) Update(key int, value E) (E, error) {
	d.lock.Lock()
	defer func() {
		d.lock.Unlock()
	}()
	_, found := d.data[key]
	d.logger.With(
		zap.Any("value", value),
		zap.Any("key", key),
	).Info("db update")
	if !found {
		return value, &errors.ErrEntityNotFound
	}
	copied := value.SetId(key)
	d.data[key] = copied.(E)
	d.logger.With(zap.Any("value", copied)).Info("db set")
	return copied.(E), nil
}

// Delete - delete the entry from database by key
func (d *Database[E]) Delete(key int) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, found := d.data[key]
	if !found {
		return &errors.ErrEntityNotFound
	}
	delete(d.data, key)
	return nil
}
