package main

import (
	"fmt"
	"golang_tutotorial/dob"
	"time"
)
func spaces(){
	fmt.Println("**********************************")
}

func main() {
	// time library to find out the current time 
	currentTime := time.Now()
	fmt.Println("current time is :", currentTime)
	// fmt.Printf("formatted time is :%T",currentTime) //find out the time type (time.Time)

	// formattedTime := currentTime.Format("2006/01/02 , Monday ,  03:04 PM") //this is way to get time in proper format 2006/01/02 , Monday , 03:04 PM(official launched date of go )
	// fmt.Println("time in format is :", formattedTime) // 2026/04/02 , Thursday ,  05:36 PM

	// layoutTime:= "02/01/2006"
	// timeStr := "23/04/2003"
	// formattedtimeStr,_ := time.Parse(layoutTime, timeStr)
	// fmt.Println("formaatted time is", formattedtimeStr)


	// // add one day to current time 
	// oneDayExtend:= currentTime.Add(24 * time.Hour)
	// formatted := oneDayExtend.Format("2006/01/02")
	// fmt.Println("new Extension time is: ", formatted)

	spaces()

	var inputDOB string
	fmt.Println("Enter your DOB (DD/MM/YYYY):")
	fmt.Scanln(&inputDOB)

	age, err := dob.AgeCalculator(inputDOB)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Your age is:", age)

	spaces()

	


		
}