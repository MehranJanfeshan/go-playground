package json_decoder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// if you are calling Rest_API that returns back a valid and defined json you can user Json.NewDecoder to pars and transfer it to the structure.

func DecodeStructure() {
	uri := "http://dummy.restapiexample.com/api/v1/employees"

	resp, err := http.Get(uri)
	if err != nil {
		log.Panic("Error while requesting response!", err)
	}

	defer resp.Body.Close()

	var sampleEmployee [] Employee

	err = json.NewDecoder(resp.Body).Decode(&sampleEmployee)

	fmt.Print(sampleEmployee[0])
}

// Unmarshal Json string

func DecodeWStringJsonStructure() {
	var JSON = `[{
         "employee_name": "Mehran",
         "employee_salary": "739303",
         "employee_age": "32",
         "profile_image": ""
       }]`
	var sampleEmployee [] Employee

	err := json.Unmarshal([]byte(JSON), &sampleEmployee)

	if err != nil {
		log.Panic("Error while requesting response!", err)
	}

	fmt.Print(sampleEmployee[0])
}

// if you are parsing a Json string and you are not sure about it's structure and want to have more flexibility you can use map
func DecodeWStringJsonWithoutStructure() {
	var JSON = `{
         "employee_name": "Mehran",
         "employee_salary": "739303",
         "employee_age": "32",
         "profile_image": ""
       }`
	var sampleEmployee map[string]interface{}

	err := json.Unmarshal([]byte(JSON), &sampleEmployee)

	if err != nil {
		log.Panic("Error while requesting response!", err)
	}

	fmt.Println(sampleEmployee["employee_name"])
	fmt.Println(sampleEmployee["employee_salary"])
	fmt.Println(sampleEmployee["employee_age"])
	fmt.Println(sampleEmployee["profile_image"])

}
