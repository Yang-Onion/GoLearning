package main

import (
	"fmt"
	"time"
)

func funcSwitch() {
	i := 2
	switch i {
	case 1:
		fmt.Println("偶数")
	case 2:
		fmt.Println("奇数")
	default:
		fmt.Println("负数或非整数")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It'2 workday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon")
	case t.Hour() < 24:
		fmt.Println("Good Ninght")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("unknown type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(2)
	whatAmI("hello")
}
