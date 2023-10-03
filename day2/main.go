package main

import (
	nilcheck "advent_of_code/commonUtils"
	"fmt"
	"os"
	"strings"
)

func main(){
	part1()
	part2()
}

func part2(){
	file,err := os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileString := string(file)
	fileLine := strings.Split(fileString, "\n")
	score:=0
	
	for _,ele:= range(fileLine){
		choice := strings.Split(ele, " ")
		// fmt.Print(choice)
		if choice[0]=="A" && choice[1]=="X"{
			score+=3
		} else if choice[0]=="A" && choice[1]=="Y"{
			score+=4
		} else if choice[0]=="A" && choice[1]=="Z"{
			score+=8
		} else if choice[0]=="B" && choice[1]=="X"{
			score+=1
		} else if choice[0]=="B" && choice[1]=="Y"{
			score+=5
		} else if choice[0]=="B" && choice[1]=="Z"{
			score+=9
		} else if choice[0]=="C" && choice[1]=="X"{
			score+=2
		} else if choice[0]=="C" && choice[1]=="Y"{
			score+=6
		} else if choice[0]=="C" && choice[1]=="Z"{
			score+=7
		}
	}
	fmt.Println(score)
}

func part1(){
	file,err := os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileString := string(file)
	fileLine := strings.Split(fileString, "\n")
	score:=0
	
	for _,ele:= range(fileLine){
		choice := strings.Split(ele, " ")
		// fmt.Print(choice)
		if choice[0]=="A" && choice[1]=="X"{
			score+=4
		} else if choice[0]=="A" && choice[1]=="Y"{
			score+=8
		} else if choice[0]=="A" && choice[1]=="Z"{
			score+=3
		} else if choice[0]=="B" && choice[1]=="X"{
			score+=1
		} else if choice[0]=="B" && choice[1]=="Y"{
			score+=5
		} else if choice[0]=="B" && choice[1]=="Z"{
			score+=9
		} else if choice[0]=="C" && choice[1]=="X"{
			score+=7
		} else if choice[0]=="C" && choice[1]=="Y"{
			score+=2
		} else if choice[0]=="C" && choice[1]=="Z"{
			score+=6
		}
	}
	fmt.Println(score)
}