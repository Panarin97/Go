package main

import (
     "fmt"
     "runtime"
	 "time"
 )

 func set_x (x *int)  {
		*x = *x * 2

 }

 func set_x_2 (x *int)  {
	   *x = *x + 1

}

// explanation:
/* There are 2 Go routines: 
 1. One that multiplies given value by 2
 2. One that adds 1 to the given value
 Due to Data racing there is no way to tell which operation is going to be performed first. 
 We test by value 0: 
 If the routines execute in this order 1 -> 2, then the output value is (0*2)+1 = 1
 Else if the routines execute in this order 2 -> 1, then the output value is (0+1)*2 = 2
 */

 func main() {

	var x = 0;

	// 1 logical processor for 2 routines
	runtime.GOMAXPROCS(1)

	time.Sleep(100 * time.Millisecond)

	go set_x_2(&x)
	// you can try to decomment the next line of code. 
	// It works absolutely fine in my case however, the program produces different output each time.
	// time.Sleep(100 * time.Millisecond)
	go set_x(&x)

	// proof of performance in the attached screenshots.

	time.Sleep(100 * time.Millisecond)
	
	fmt.Println(x)

 }

 