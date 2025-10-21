package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Student struct {
	Name string
	Age  int
}

// SafeDiv делит a на b и возвращает результат.
// Если b == 0 — верни ошибку.
func SafeDiv(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("нельзя делить на 0")
	}
	return a / b, nil
}

// FirstOrErr возвращает первый элемент среза строк.
// Если срез пустой — верни ошибку.
func FirstOrErr(items []string) (string, error) {
	if len(items) == 0 {
		return "", errors.New("пустой срез")
	}
	return items[0], nil
}

// NewStudent создаёт студента. Если имя пустое или возраст < 0 — верни ошибку.
func NewStudent(name string, age int) (Student, error) {
	if name == "" || age <= 0 {
		return Student{}, errors.New("имя пустое или неккоректный возраст")
	}
	newStudent := Student{Name: name, Age: age}
	return newStudent, nil
}

// ParsePrice возвращает цену из строки в формате "123.45".
// Если формат неверный — верни ошибку из ParseFloat.
func ParsePrice(s string) (float64, error) {
	price, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.New("неверный формат строки")
	}
	return price, nil
}

func main() {

	//*********//
	// Подзадача 1 — Безопасное деление

	if res, err := SafeDiv(10, 2); err != nil {
		fmt.Println("ошибка деления:", err)
	} else {
		fmt.Println("Pезультат:", res)
	} // Pезультат: 5

	//*********//
	// Подзадача 2 — Первый элемент среза

	slice := []string{}
	if res, err := FirstOrErr(slice); err != nil {
		fmt.Println("Ошибка среза:", err)
	} else {
		fmt.Println("Первый элемент среза:", res)
	} // Ошибка среза: пустой срез

	//*********//
	// Подзадача 3

	if res, err := NewStudent("Adam", 19); err != nil {
		fmt.Println("Ошибка структуры:", err)
	} else {
		fmt.Println("Новый студент:", res)
	} // Новый студент: {Adam 19}

	res, err := NewStudent("", 19)
	if err != nil {
		fmt.Println("Ошибка структуры:", err)
	} else {
		fmt.Println("Новый студент:", res)
	} // Ошибка структуры: имя пустое или неккоректный возраст

	if res, err := NewStudent("", 0); err != nil {
		fmt.Println("Ошибка структуры:", err)
	} else {
		fmt.Println("Новый студент:", res)
	} // Ошибка структуры: имя пустое или неккоректный возраст

	//**********//
	// Подзадача 4

	if res, err := ParsePrice("232.34"); err != nil {
		fmt.Println("Ошибка вывода:", err)
	} else {
		fmt.Println("Результат:", res)
	} // Результат: 232.34

	if res, err := ParsePrice("avc"); err != nil {
		fmt.Println("Ошибка вывода:", err)
	} else {
		fmt.Println("Результат:", res)
	} // Ошибка вывода: неверный формат строки

}
