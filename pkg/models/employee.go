package models

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func (e Employee) SetId(id int) Entity {
	return Employee{
		ID:       id,
		Name:     e.Name,
		Position: e.Position,
		Salary:   e.Salary,
	}
}

func (e Employee) GetId() int {
	return e.ID
}
