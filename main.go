package main

import (
	"fmt"
	"math"
)

func main() {
	//1 zadaniye
	bannerWidth := 12
	bannerHeight := 8
	bannerArea := bannerWidth * bannerHeight
	fmt.Println(bannerArea)
	halfBannerArea := bannerArea / 2
	fmt.Println(halfBannerArea)
	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println(bannerBorderLength)
	//2 zadaniye
	boxCount := 29
	people := 5
	leftoverBoxes := boxCount / people
	fmt.Println(leftoverBoxes)
	var tempMorning float64 = 18.5
	var tempAfternoon float64 = 25.0
	var tempEvening float64 = 15.3
	totalTemp := tempMorning + tempAfternoon + tempEvening
	fmt.Println(totalTemp)
	averageTemp := tempMorning + tempAfternoon + tempEvening/3
	fmt.Println(averageTemp)
	knownWords := 47
	wordsGoal := 120
	progressPercent := (wordsGoal % knownWords)
	fmt.Println(progressPercent)
	coins := 0
	fmt.Println("Начальное количество монет:", coins)

	coins += 500
	fmt.Println("После выполнения задания (+500):", coins)
	coins += 1200
	fmt.Println("После получения бонуса (+1200):", coins)
	coins /= 2
	fmt.Println("После покупки в магазине (потрачена половина):", coins)
	coins *= 2
	fmt.Println("После игрового события (увеличение в 2 раза):", coins)
	coins -= 300
	fmt.Println("После траты 300 монет:", coins)

	participants := 42
	groupCount := 8
	participantsPerGroup := participants / groupCount
	fmt.Println("Участников в каждой группе:", participantsPerGroup)

	fmt.Println(20 - 4*3)
	fmt.Println((20 - 4) * 3)

	squareValue := 81.0
	sqrtResult := math.Sqrt(squareValue)

	multiplier := 5.0
	exponent := 2.0
	powResult := math.Pow(multiplier, exponent)

	fmt.Println("Квадратный корень из 81:", sqrtResult)
	fmt.Println("5 в степени 2:", powResult)
	//
}
