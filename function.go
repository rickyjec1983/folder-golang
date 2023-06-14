package main

import (
	"fmt"
)

type HasName interface {
	getname() string
}

func sayHello(HasNamexx HasName) {
	fmt.Println("hello", HasNamexx.getname())
}

func (Customer1 Customer) getname() string {
	return Customer1.name
}

type Customer struct {
	name, address string
	age           int
}

func main() {
	var sCustomerxxx Customer
	sCustomerxxx.name = "ricky lai"
	sayHello(sCustomerxxx)

	var sCustomer Customer
	sCustomer.address = "xxx"
	sCustomer.name = "xxx"
	sCustomer.age = 9

	joko := Customer{
		name:    "ricky",
		address: "jkt",
		age:     3,
	}

	fmt.Println("struct", joko)
	//normalparam("ricky 1", "ricky 2")

	//var sArrayName2 []string
	//sArrayName2 = []string{"ricky", "ronald"}

	//var sTampung string = returnValue("ricky 1", "ricky 2")
	//fmt.Println(sTampung)

	//return function with array
	//var sArrayName = []string{"ricky", "ronald"}
	//var sArrayNameResult = returnValueArray("ricky first", sArrayName)
	//fmt.Println(sArrayNameResult)

	//function return value> 1
	var hitung1, hitung2 = calculate(1)
	fmt.Println("result 1 : ", hitung1, "result 2 : ", hitung2)

	person := map[string]string{
		"nama":   "ricky",
		"alamat": "cikupa",
	}

	fmt.Println(person)

	stampungStudent := variadic("ricky", "ronald", "wisnu", "aang")
	fmt.Printf("%v", stampungStudent)
}

// func normalparam(nama1, nama2 string) {
// 	fmt.Println("nama 1", nama1, "nama 2", nama2)
// }

// func returnValue(name1, name2 string) string {
// 	return "nama 1" + name1 + "nama 2" + name2
// }

// func returnValueArray(name1 string, sArrayName []string) string {
// 	var sJoinArray = strings.Join(sArrayName, " , ")
// 	return "nama 1" + name1 + "nama array : " + sJoinArray
// }

func calculate(angka1 float64) (float64, float64) {

	var hitung1 float64 = angka1 + 1
	var hitung2 float64 = angka1 + 2

	return hitung1, hitung2
}

func variadic(names ...string) []map[string]string {

	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result

}
