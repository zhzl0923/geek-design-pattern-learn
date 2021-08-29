package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	employeeA := NewFullTimeEmployee("A", 2000, 40)
	employeeB := NewPartTimeEmployee("B", 30, 40)
	visitor := &FADepartment{}
	employeeA.Accept(visitor)
	employeeB.Accept(visitor)
}
