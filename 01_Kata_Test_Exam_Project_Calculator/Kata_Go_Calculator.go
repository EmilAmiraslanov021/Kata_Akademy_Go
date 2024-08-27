package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	total := 0
	prevValue := 0

	for _, char := range s {
		value, exists := romanMap[char]
		if !exists {
			panic("Ошибка: АяЯй! Неверный символ в римских числах")
		}
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}
	return total
}

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; num > 0; i++ {
		for num >= vals[i] {
			num -= vals[i]
			result += romans[i]
		}
	}
	return result
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Ошибка: АяЯй! Деление на ноль") // в этой краевой ситуации у меня срабатывает паника из условия
			//что Калькулятор должен принимать на вход только числа от 1 до 10 включительно)))
		}
		return a / b
	default:
		panic("Ошибка: АяЯй! Неверный оператор")
	}
}

func isStrNum(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: III + V или 3 + 5):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Ошибка: АяЯй! Неверный формат ввода, строка не является математической операцией. Либо формат математической операции не удовлетворяет заданию")
	}
	// я знаю что в 82 строчке можно было бы использовать слайсы для вывода 2-ой строки с новой строчки, Но я уже не успеваю(((
	var a, b int
	operator := parts[1]

	if strings.ContainsAny(parts[0], "IVXLCDM") && strings.ContainsAny(parts[2], "IVXLCDM") {
		a = romanToInt(parts[0])
		b = romanToInt(parts[2])
		if a < 1 || b < 1 || a > 10 || b > 10 {
			panic("Ошибка: АяЯй! Римские числа должны быть от I до X")
		}
	} else if isStrNum(parts[0]) && isStrNum(parts[2]) {
		a, _ = strconv.Atoi(parts[0])
		b, _ = strconv.Atoi(parts[2])
		if a < 1 || b < 1 || a > 10 || b > 10 {
			panic("Ошибка: АяЯй! Арабские числа должны быть от 1 до 10. На 0 делить тоже нельзя!!")
		}
	} else {
		panic("Ошибка: АяЯй! Неверный формат ввода. Используйте только арабские или только римские числа.")
	}

	result := calculate(a, b, operator)

	if strings.ContainsAny(parts[0], "IVXLCDM") {
		if result < 1 {
			panic("Ошибка: АяЯй! Результат не может быть меньше I в римских числах. В римской системе нет отрицательных чисел!!")
		}
		fmt.Printf("Результат: %s\n", intToRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
