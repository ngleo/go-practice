package main

import (
	  "fmt"
		"encoding/json"
)

type SensorReading struct {
  	Name string `json:"name"`
   	Capacity int `json:"capacity"`
  	Time string `json:"timestamp"`
}

func main() {
	  reading1 := SensorReading{Name: "reading1", Capacity: 40, Time: "2021-07-07T09:22:00Z"}
		fmt.Printf("%+v\n", reading1)
	  reading1Bytes, err := json.Marshal(reading1)
		if (err != nil) {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", string(reading1Bytes))

	
	  jsonString := `{"name": "reading2", "capacity": 70, "timestamp":"2021-07-10T11:26:00Z"}`
	  var reading2 SensorReading
	  err = json.Unmarshal([]byte(jsonString), &reading2)
    if err != nil {
			  fmt.Println(err)
		}
		fmt.Printf("%+v\n", reading2)
}