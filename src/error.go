package main

import "errors"

//START OMIT

func hello(input string) (string, error) {
	if length := len(input); length > 0 {
		return "Hello, " + input, nil
	} else {
		return "", errors.New("MISSING INPUT")
	}

}

func main() {
	val, err := hello("")

	if err != nil {
		println(err.Error())

	} else {
		println(val)
	}
}

//END OMIT
