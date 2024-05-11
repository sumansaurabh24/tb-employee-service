package models

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestEmployee(t *testing.T) {
	e := Employee{
		Name:     "testname",
		Position: "test position",
		Salary:   1110,
	}
	copied := e.SetId(100)
	log.Println("set value", copied)
	require.NotNil(t, copied)
	require.Equal(t, copied.GetId(), 100)
}
