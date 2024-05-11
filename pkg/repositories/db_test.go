package repositories

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/sumansaurabh24/tb-employee-service/pkg/errors"
	"github.com/sumansaurabh24/tb-employee-service/pkg/models"
	"go.uber.org/zap"
	"sync"
	"testing"
)

func TestDatabase_Update_MultiThreaded(t *testing.T) {
	logger := zap.Must(zap.NewProduction())
	db := NewDatabase[models.Entity](logger.Sugar())
	var wgSet sync.WaitGroup
	wgSet.Add(10)
	for i := range [10]int{} {
		go func(i int) {
			db.Insert(models.Employee{
				Name:     fmt.Sprintf("Suman_%d", i),
				Position: "Lead Member Technical Staff",
				Salary:   10000,
			})
			wgSet.Done()
		}(i)
	}
	wgSet.Wait()
	findOne, err := db.Find(1)
	require.Nil(t, err)
	require.NotNil(t, findOne)
	findTen, err := db.Find(10)
	require.Nil(t, err)
	require.NotNil(t, findTen)
	findFive, err := db.Find(5)
	require.Nil(t, err)
	require.NotNil(t, findFive)
}

func TestDatabase_Find(t *testing.T) {
	logger := zap.Must(zap.NewProduction())
	db := NewDatabase[models.Entity](logger.Sugar())
	emp := db.Insert(models.Employee{
		Name:     "Suman Saurabh",
		Position: "Lead Member Technical Staff",
		Salary:   10000,
	})
	require.NotNil(t, emp)
	require.Equal(t, 1, emp.GetId())
}

func TestDatabase_Update(t *testing.T) {
	logger := zap.Must(zap.NewProduction())
	db := NewDatabase[models.Entity](logger.Sugar())
	entity := db.Insert(models.Employee{
		Name:     "Suman Saurabh",
		Position: "Lead Member Technical Staff",
		Salary:   10000,
	})
	updateEnt, _ := db.Update(entity.GetId(), models.Employee{
		Name:     "Suman Saurabh Updated",
		Position: "Lead Member Technical Staff",
		Salary:   10000,
	})

	require.NotNil(t, updateEnt)
	require.Equal(t, entity.GetId(), updateEnt.GetId())
	emp := updateEnt.(models.Employee)
	require.Equal(t, "Suman Saurabh Updated", emp.Name)
}

func TestDatabase_Delete(t *testing.T) {
	logger := zap.Must(zap.NewProduction())
	db := NewDatabase[models.Entity](logger.Sugar())
	entity := db.Insert(models.Employee{
		Name:     "Suman Saurabh",
		Position: "Lead Member Technical Staff",
		Salary:   10000,
	})
	id := entity.GetId()

	err := db.Delete(id)
	require.Nil(t, err)

	find, err := db.Find(id)
	require.Nil(t, find)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), errors.ErrEntityNotFound.Error())
}
