package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	romNum = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"XI":   11,
		"XX":   20,
		"XXX":  30,
		"XL":   40,
		"L":    50,
		"LX":   60,
		"LXX":  70,
		"LXXX": 80,
		"XC":   90,
		"C":    100,
	}
)

/*
Проверка строки на правильность
*/
func isValid(tmp []string) (err error) {
	var tmpNum int
	arabic, rom := false, false
	if len(tmp) != 3 {
		err = errors.New("too many arguments")
	} else {
		for idx, elem := range tmp {
			if idx != 1 {
				_, ok := romNum[elem]
				if ok {
					rom = true
				} else if !ok {
					arabic = true
					tmpNum, err = strconv.Atoi(elem)
					if err != nil {
						err = errors.New(fmt.Sprintf("this character is not valid %s", elem))
						break
					}
					if tmpNum > 10 || tmpNum < 0 {
						err = errors.New(fmt.Sprintf("num %v is not valid", elem))
						break
					}
				}
				if rom && arabic {
					err = errors.New("different number systems are used simultaneously")
				}
			} else {
				switch elem {
				case "+":
				case "-":
				case "*":
				case "/":
					break
				default:
					err = errors.New(fmt.Sprintf("this character is not valid %s", elem))
					break
				}
			}
		}
	}
	return err
}

/*
Преобразование строки в римскую систему счисления
*/
func makeRom(num int) (res string, err error) {
	var buff []string
	tmp, counter := 0, 1
	for num > 0 {
		tmp = num % 10
		if tmp != 0 {
			tmp *= counter
			for key, val := range romNum {
				if val == tmp {
					buff = append(buff, key)
				}
			}
		}
		num /= 10
	}
	for i := len(buff) - 1; i >= 0; i-- {
		res += buff[i]
	}
	if len(res) == 0 {
		err = errors.New("цифры 0 нет в римской системе")
	}
	return res, err
}

/*
Произведение математической операции
*/
func makeAction(src []string) (res string, err error) {
	var numAnswer, left, right int
	var arabic = true

	if _, ok := romNum[src[0]]; ok {
		left, _ = romNum[src[0]]
		right, _ = romNum[src[2]]
		arabic = false
	} else {
		left, _ = strconv.Atoi(src[0])
		right, _ = strconv.Atoi(src[2])
	}

	switch src[1] {
	case "+":
		numAnswer = left + right
		break
	case "-":
		numAnswer = left - right
		break
	case "*":
		numAnswer = left * right
		break
	case "/":
		if right == 0 {
			err = errors.New("division by zero")
			break
		}
		numAnswer = left / right
		break
	}
	if err == nil {
		if !arabic {
			if res, err = makeRom(numAnswer); err != nil {
				return res, err
			}
		} else {
			return strconv.Itoa(numAnswer), err
		}
	}
	return res, err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	dumpStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	dumpStr = strings.Trim(dumpStr, "\r\n")
	splitStr := strings.Split(dumpStr, " ")

	err = isValid(splitStr)
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println(makeAction(splitStr))
	}

}
