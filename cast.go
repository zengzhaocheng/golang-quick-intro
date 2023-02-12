package main

import (
	"fmt"
	"unsafe"
)

type Employee struct {
	//基本工资
	BaseWage float64
	eType    E_TYPE //员工类型
}
type E_TYPE string

const (
	WORKER    E_TYPE = "工人"
	FARMER           = "农民"
	TEACHER2         = "老师"
	SCIENTIST        = "科学家"
	WAITER           = "服务员"
)

type Worker struct {
	Employee
}

func (e *Employee) getYearBaseSalary() float64 {
	return 12 * e.BaseWage
}

func (e *Employee) PrintSalary() {
	fmt.Printf("%s全年工资：", e.eType)
	switch e.eType {
	case WORKER:
		fallthrough
	case FARMER:
		fallthrough
	case WAITER:
		fmt.Printf("%.2f\n", e.BaseWage)
	case TEACHER2:
		//下面使用unsafe包的Pointer将父类(Employee)实例转为子类(Teacher2)实例
		//将父类实例转为通用指针
		uP := unsafe.Pointer(e)
		//再转换为本地 Teacher2 结构体
		pT := (*Teacher2)(uP)
		//转换完毕，可以输出子类特有属性
		fmt.Println("测试-->老师的课酬是:", pT.ClassRemuneration)
		fmt.Printf("%.2f\n", e.getYearBaseSalary()+pT.ClassRemuneration)
	case SCIENTIST:
		uP := unsafe.Pointer(e)
		pT := (*Scientist)(uP)
		fmt.Printf("%.2f\n", e.getYearBaseSalary()+pT.YearEndBonus)
	default:
		fmt.Println("未能匹配到对应的员工类型")
	}

}

type Farmer struct {
	Employee
}

type Teacher2 struct {
	Employee
	//课酬 (元/天)
	ClassRemuneration float64
}

type Scientist struct {
	Employee
	//年终奖
	YearEndBonus float64
}

type Waiter struct {
	Employee
}

// ----------------测试------------------------
func main() {

	println("面试题:https://studygolang.com/articles/16681")
	//实例化工人
	var worker *Worker = &Worker{Employee: Employee{BaseWage: 120, eType: WORKER}}
	//实例化农民
	var farmer *Farmer = &Farmer{
		Employee: Employee{
			BaseWage: 150,
			eType:    FARMER,
		},
	}
	//实例化老师
	employee := &Employee{BaseWage: 250, eType: TEACHER2}
	var teacher2 *Teacher2 = &Teacher2{Employee: *employee}
	teacher2.ClassRemuneration = 500
	//实例化科学家
	var scientist *Scientist = &Scientist{
		Employee{BaseWage: 300, eType: SCIENTIST},
		200,
	}
	//实例化服务人员
	var waiter *Waiter = &Waiter{Employee: Employee{BaseWage: 150, eType: WAITER}}
	//测试输出
	worker.PrintSalary()
	farmer.PrintSalary()
	waiter.PrintSalary()
	teacher2.PrintSalary()
	scientist.PrintSalary()
}
