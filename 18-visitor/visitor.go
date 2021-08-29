package visitor

import (
	"fmt"
	"strconv"
)

//雇员
type Employee interface {
	Accept(visitor Department)
}

type Department interface {
	VisitFullTimeEmployee(employee *FullTimeEmployee)
	VisitPartTimeEmployee(employee *PartTimeEmployee)
}

//全职雇员
type FullTimeEmployee struct {
	Name       string
	WeeklyWage float64
	WorkTime   int
}

func NewFullTimeEmployee(name string, weeklyWage float64, workTime int) *FullTimeEmployee {
	return &FullTimeEmployee{Name: name, WeeklyWage: weeklyWage, WorkTime: workTime}
}

func (e *FullTimeEmployee) Accept(visitor Department) {
	visitor.VisitFullTimeEmployee(e)
}

//兼职雇员
type PartTimeEmployee struct {
	Name     string
	HourWage float64
	WorkTime int
}

func NewPartTimeEmployee(name string, hourWage float64, workTime int) *PartTimeEmployee {
	return &PartTimeEmployee{Name: name, HourWage: hourWage, WorkTime: workTime}
}

func (e *PartTimeEmployee) Accept(visitor Department) {
	visitor.VisitPartTimeEmployee(e)
}

//财务部门
type FADepartment struct{}

//计算全职雇员薪资
func (d *FADepartment) VisitFullTimeEmployee(e *FullTimeEmployee) {
	weekWage := e.WeeklyWage
	if e.WorkTime > 40 {
		weekWage = weekWage + float64((e.WorkTime-40)*100)
	} else if e.WorkTime < 40 {
		weekWage = weekWage - float64((40-e.WorkTime)*80)
		if weekWage < 0 {
			weekWage = 0
		}
	}
	fmt.Println("正式员工：" + e.Name + "，实际工资为：" + strconv.FormatFloat(weekWage, 'f', 10, 64) + "元。")

}

//计算兼职雇员薪资
func (d *FADepartment) VisitPartTimeEmployee(e *PartTimeEmployee) {
	fmt.Println("临时工：" + e.Name + "，实际工资为：" + strconv.FormatFloat(float64(e.WorkTime)*e.HourWage, 'f', 10, 64) + "元。")
}
