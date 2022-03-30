package services

import "fmt"

type Exercises interface {
	ScanMultiple(maxNumber int, anotherNumber int, typeNumber string) int
	Sum(a int, b int) int
	Multiply(a int, b int) int
	SumMultiple(typeNumber string, anotherNumber int, i int, numberToSum int) int
	ImHappyNumber(entryNumber int64, listOfThoseThatHavePassed []int64) bool
	ConvertWordToNumber(wordToConvert string) int
}

type exercisesService struct {}

func NewExercisesService() Exercises {
	return &exercisesService{}
}

func (e *exercisesService) ImHappyNumber(entryNumber int64, listOfThoseThatHavePassed []int64) bool {
	panic("implement me")
}

func (e *exercisesService) ConvertWordToNumber(wordToConvert string) int {
	panic("implement me")
}

func (e *exercisesService) ScanMultiple(maxNumber int, anotherNumber int, typeNumber string) int {
	result := 0

	for i := 1; i < maxNumber; i++ {

		result = e.SumMultiple(typeNumber, anotherNumber, i, result)

		fmt.Println(result)
	}

	return result
}


func (e *exercisesService) SumMultiple(typeNumber string, anotherNumber int, i int, numberToSum int) int {

	if typeNumber == "OR" {

		if i%3 == 0 {
			numberToSum += i
		}

		if i%5 == 0 {
			numberToSum += i
		}
	}
	if typeNumber == "AND" {

		if i%3 == 0 && i%5 == 0 {
			numberToSum += i
		}

	}
	if typeNumber == "OR_AND_OTHER" {
		if (i%3 == 0 || i%5 == 0) && i%anotherNumber == 0 {
			numberToSum += i
		}
	}
	return numberToSum
}

func (e *exercisesService) Sum(a int, b int) int {
	return a + b
}

func (e *exercisesService) Multiply(a int, b int) int {
	return a * b
}
