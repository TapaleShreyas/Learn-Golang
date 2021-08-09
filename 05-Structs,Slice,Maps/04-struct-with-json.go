package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Convert struct to json and json to struct

type Laptop struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	OS    string `json:"operating_system"`
	HDD   uint   `json:"hard_disk_drive"`
	SSD   uint   `json:"solid_state_drive"`
}

func main() {
	hp := Laptop{
		Brand: "HP",
		Model: "15-r007TZ",
		OS:    "windows 10 prefessional",
		HDD:   1000,
		SSD:   250,
	}
	fmt.Println("hp:", hp)
	hpJson, _ := json.Marshal(hp)
	fmt.Println("json version of hp laptop:", bytes.NewBuffer(hpJson))

	fmt.Println("\nInitialize dell laptop using json")
	dellJson := `{"brand":"DELL","model":"10-KKR00B","operating_system":"linux","hard_disk_drive":500,"solid_state_drive":250}`
	fmt.Println("dell laptop dellJson:", dellJson)

	var dell Laptop
	json.Unmarshal([]byte(dellJson), &dell)
	fmt.Println("dell laptop struct:", dell)
}
