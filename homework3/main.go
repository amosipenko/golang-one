package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println("Калькулятор.")
	fmt.Println("Допустимые операции:")
	fmt.Println(" ____________________________________________________________________________________________________")
	fmt.Println(" |  a+b сложение  | a*b умножение |  a%b остаток от деления   | a! факториал         |  a>b больше  |")
	fmt.Println(" |  a-b вычитание | a/b деление   |  a^b возведение в степень | #a квадратный корень |  a<b меньше  |")
	fmt.Println(" ____________________________________________________________________________________________________")
	fmt.Println("Разрешены отрицательные и дробные числа через точку. Пример: -5+2.5")

	for true {
		fmt.Println("Введите числа и операцию. Пример: 5+2. Для завершения введите 0.")

		var scannedString string
		fmt.Scanln(&scannedString)

		if scannedString == "0" {
			break
		}

		var a, b float64
		var operation string
		a, operation, b = calcArgumentsFromString(scannedString)

		switch operation {
		case "+":
			fmt.Printf("%.3g\n", a+b)
		case "-":
			fmt.Printf("%.3g\n", a-b)
		case "*":
			fmt.Printf("%.3g\n", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Нельзя делить на 0")
			} else {
				fmt.Printf("%.3g\n", a/b)
			}
		case "%":
			fmt.Println(int(a) % int(b))
		case "^":
			fmt.Printf("%.3g\n", math.Pow(a, b))
		case "!":
			if a >= 0 {
				fmt.Println(factorial(int64(a)))
			} else {
				fmt.Println("Для расчета факториала введите положительное число")
			}
		case "#":
			fmt.Printf("%.3g\n", math.Sqrt(b))
		case ">":
			fmt.Println(a > b)
		case "<":
			fmt.Println(a < b)
		default:
			fmt.Println("Некорректная операция")
		}

	}
}

func calcArgumentsFromString(str string) (float64, string, float64) {
	var aStr, bStr, oper string
	for i := 0; i < len(str); i++ {
		// Если встретили цифру или точку, то это аргумент
		switch {
		case str[i] >= '0' && str[i] <= '9' || str[i] == '.':
			if oper == "" {
				aStr += string(str[i])
			} else {
				bStr += string(str[i])
			}
		case str[i] == '-':
			// отдельно обрабатываем минус, т.к. он может относиться к аргументу или быть операцией
			switch {
			case aStr == "" && oper == "":
				// еще нет первого аргумента и операции - значит это минус в первом аргументе
				aStr += string(str[i])
			case oper == "":
				// еще нет операции - значит это минус в операции
				oper += string(str[i])
			default:
				// иначе это минус во втором аргументе
				bStr += string(str[i])
			}
		case str[i] == '+' || str[i] == '*' || str[i] == '/' || str[i] == '%' || str[i] == '!' || str[i] == '^' || str[i] == '#' || str[i] == '<' || str[i] == '>':
			// Если один из допустимых символов операции, то это операция: + || * || / || % || ! || ^ || # || < || >
			oper += string(str[i])
		case str[i] != ' ':
			// пробел еще допускаем и игнорим, а остальное не простим
			panic("Некорректная операция")

		}
	}

	var a, b float64
	var err error
	if oper == "#" {
		if aStr != "" {
			// для квадратного корня проверим, что первый аргумент не заполнен
			panic("Для расчета квадратного корня не нужно вводить первый аргумент")
		}
	} else {
		// для остальных операций первый аргумент должен быть заполнен и преобразовываться во float
		if aStr == "" {
			panic("Не введен первый аргумент")
		} else {
			a, err = strconv.ParseFloat(aStr, 32)
			if err != nil {
				panic(err)
			}
		}
	}

	if oper == "!" {
		if bStr != "" {
			// для факториала проверим, что второй аргумент не заполнен
			panic("Для расчета факториала не нужно вводить второй аргумент")
		}
	} else {
		// для остальных операций второй аргумент должен быть заполнен и преобразовываться во float
		if bStr == "" {
			panic("Не введен второй аргумент")
		} else {
			b, err = strconv.ParseFloat(bStr, 32)
			if err != nil {
				panic(err)
			}
		}
	}

	return a, oper, b
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}

	var f int64
	f = 1

	for i := int64(1); i <= n; i++ {
		f *= i
	}
	return f
}
