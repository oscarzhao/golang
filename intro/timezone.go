 package main

 import (
 	"fmt"
 	"time"
 )

 func main() {

 	now := time.Now()

 	fmt.Println("Now: ", now)

 	fmt.Printf("Location: %#v\n", now.Location())

 	// get the time zone name
 	z, _ := now.Zone()

 	fmt.Printf("Location(Time Zone): %v\n", z)

 	// load different time zone
 	est, err := time.LoadLocation("EST") //<---------------------- here !

 	if err != nil {
 		fmt.Println(err)
 	}

 	fmt.Println("Load Location : ", est)

 	dayInEST := time.Date(2015, 18, 5, 12, 15, 0, 0, est)

 	fmt.Println("This code is created on : ", dayInEST.Format("Monday"))

 }
