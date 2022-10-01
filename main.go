package main

import (
	"fmt"
	"log"
	"time"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
)

func main() {
	fmt.Println("Hello world")
	i2c, err := i2c.New(pca9685.Address, "/dev/i2c-1")
	if err != nil {
		log.Fatal(err)
	}
	pca0, err0 := pca9685.New(i2c, nil)
	if err0 != nil {
		log.Fatal(err0)
	}
	pca1, err1 := pca9685.New(i2c, nil)
	if err1 != nil {
		log.Fatal(err1)
	}
	pca1.SetChannel(1, 0, 130)
	pca0.SetChannel(0, 0, 130)
	servo1 := pca1.ServoNew(1, nil)
	servo0 := pca0.ServoNew(0, nil)
	test := 2
	switch {
	case test == 2:
		for x := 0; x < 200; x++ {
			i := 50
			//	for i := 0; i < 64; i++ {
			servo1.Angle(i)
			// if x < 100 || x > 90 {
			if x > 100 {
				servo0.Angle(65)
			} else {
				servo0.Angle(45)

			}
			time.Sleep(10 * time.Millisecond)
			//	}
		}

	case test == 1:
		speed := 10
		for x := 0; x < 4; x++ {
			fmt.Println(x)
			servo1.Angle(speed)
			time.Sleep(10 * time.Millisecond)
		}
	case test == 0:
		fmt.Println("Running Test 0")
		speed := 54
		//		for x := 0; x < 10; x++ {
		for i := 0; i < 130; i++ {
			servo1.Angle(speed)
			servo0.Angle(i)
			time.Sleep(10 * time.Millisecond)
		}
		//		}

	}
	servo0.Fraction(0.5)
	servo1.Fraction(0.5)
}
