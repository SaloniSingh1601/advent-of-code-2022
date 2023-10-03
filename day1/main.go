package main

import (
	nilcheck "advent_of_code/commonUtils"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func part1(){
	file, err := os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	maxElement := 0
	fileContent := string(file)
	fileLines := strings.Split(fileContent, "\n\n")
	for _, ele := range(fileLines){
		calories := strings.Split(ele, "\n")
		sum := 0
		for _, element := range(calories){
			elementInt, err := strconv.ParseInt(element,10, 64)
			nilcheck.CheckNilErr(err)
			sum += int(elementInt)
		}
		if sum > maxElement{
			maxElement = sum
		}
	}
	fmt.Println(maxElement)
}

func part2(){
	file, err := os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent := string(file)
	fileLines := strings.Split(fileContent, "\n\n")
	topThree := []int{0,0,0}
	// make([]int, 3)
	for _, ele := range(fileLines){
		calories := strings.Split(ele, "\n")
		sum :=0
		for _, element := range(calories){
			elementInt, err := strconv.ParseInt(element,10, 64)
			nilcheck.CheckNilErr(err)
			sum += int(elementInt)
		}
		if sum >= topThree[0]{
			topThree[2] =topThree[1]
			topThree[1] = topThree[0]
			topThree[0] = sum
		} else if sum >= topThree[1]{
			topThree[2] = topThree[1]
			topThree[1] = sum
		} else if sum>= topThree[2]{
			topThree[2] = sum
		}
	}
	result := 0
	for _, ele := range(topThree){
		result+=ele
	}
	fmt.Println(result)
}

func main(){
	part1()
	part2()
}