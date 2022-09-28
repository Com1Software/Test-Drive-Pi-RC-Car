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
	pca0, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Fatal(err)
	}
	pca0.SetChannel(1, 0, 130)
	servo0 := pca0.ServoNew(0, nil)
	servo1 := pca0.ServoNew(1, nil)
	speed:=10
	for i := 0; i < 130; i++ {
		servo0.Angle(speed)
		servo1.Angle(i)
		time.Sleep(10 * time.Millisecond)
	}

	servo0.Fraction(0.5)
	servo1.Fraction(0.5)
}
