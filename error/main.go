package main

import (
	"errors"
	"fmt"
)

func function1(arg int) (int, error) {
	if arg == 0 {
		return -1, errors.New("Can not work with 0")
	}

	return arg + 7, nil
}

type argError struct {
	arg  int
	prob string
}

func (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}

func function2(arg int) (int, error) {
	if arg == 0 {
		return -1, &argError{arg: arg, prob: "Cannot work with it"}
	}
	return arg + 7, nil
}

func main() {
	for _, v := range []int{0, 42} {
		if r, e := function1(v); e != nil {
			fmt.Println("function 1 is failed: ", e)
		} else {
			fmt.Println("function 1 is worked: ", r)
		}
	}

	for _, v := range []int{7, 42} {
		if r, err := function2(v); err != nil {
			fmt.Println("Function 2 is failed: ", err)
		} else {
			fmt.Println("Funtion 2 is worked: ", r)
		}
	}

	_, err := function2(42)
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
