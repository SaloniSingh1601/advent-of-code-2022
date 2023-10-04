package main

import (
	nilcheck "advent_of_code/commonUtils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(){
	file, err:= os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent := string(file)
	pairs := strings.Split(fileContent, "\n")
	cnt := 0
	for _, ele := range(pairs){
		elves:= strings.Split(ele, ",")
		elf1:=strings.Split(elves[0], "-")
		elf2:=strings.Split(elves[1], "-")
		elf1_0,err := strconv.ParseInt(elf1[0],10,64)
		nilcheck.CheckNilErr(err)
		elf1_1,err := strconv.ParseInt(elf1[1],10,64)
		nilcheck.CheckNilErr(err)
		elf2_0,err := strconv.ParseInt(elf2[0],10,64)
		nilcheck.CheckNilErr(err)
		elf2_1,err := strconv.ParseInt(elf2[1],10,64)
		nilcheck.CheckNilErr(err)
		if(elf1_0>=elf2_0 && elf1_1<=elf2_1) || (elf1_0<=elf2_0 && elf1_1>=elf2_1){
			cnt++
		}
	}
	fmt.Println(cnt)
}
func part2(){
	file, err:= os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent := string(file)
	pairs := strings.Split(fileContent, "\n")
	cnt := 0
	for _, ele := range(pairs){
		elves:= strings.Split(ele, ",")
		elf1:=strings.Split(elves[0], "-")
		elf2:=strings.Split(elves[1], "-")
		elf1_0,err := strconv.ParseInt(elf1[0],10,64)
		nilcheck.CheckNilErr(err)
		elf1_1,err := strconv.ParseInt(elf1[1],10,64)
		nilcheck.CheckNilErr(err)
		elf2_0,err := strconv.ParseInt(elf2[0],10,64)
		nilcheck.CheckNilErr(err)
		elf2_1,err := strconv.ParseInt(elf2[1],10,64)
		nilcheck.CheckNilErr(err)
		if(elf1_0>elf2_1) || (elf1_1<elf2_0){
			continue
		} else{
			cnt++
		}
		
	}
	fmt.Println(cnt)
}
func main(){
	part1()
	part2()
}