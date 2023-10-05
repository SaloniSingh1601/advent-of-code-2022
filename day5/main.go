package main

import (
	nilcheck "advent_of_code/commonUtils"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func part1(){
	file, err := os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent := string(file)
	parts := strings.Split(fileContent, "\n\n")
	crate_group := strings.Split(parts[0], "\n")

	stacks := [][] string{}
	for _, line := range(crate_group){
		stack_idx := 0
		for idx:=1; idx<len(line);idx+=4{
			element:=string(line[idx])
			if(stack_idx >= len(stacks)){
				if(element != " "){
					stacks = append(stacks, []string{element})
				} else{
					stacks = append(stacks, []string{})
				}
				
			} else{
				if(element != " "){
					stacks[stack_idx] = append(stacks[stack_idx], element)
				}
			}
			stack_idx++
		}
	}
	
	for _,s:= range(stacks){
		slices.Reverse(s)
	}
	steplines:=strings.Split(parts[1], "\n")
	for _,line:= range(steplines){
		markers := strings.Split(line, " ")
		if(len(markers)<5){
			continue
		}
		c, err:=strconv.ParseInt(markers[1],10,64)
		nilcheck.CheckNilErr(err)
		a, err:=strconv.ParseInt(markers[3],10,64)
		nilcheck.CheckNilErr(err)
		b, err:=strconv.ParseInt(markers[5],10,64)
		nilcheck.CheckNilErr(err)
		a-- 
		b--
		inTransit:= stacks[a][len(stacks[a])-int(c):]
		slices.Reverse(inTransit)
		stacks[a]=stacks[a][:len(stacks[a])-int(c)]
		stacks[b] = append(stacks[b], inTransit...)
	}
	for _, line:= range(stacks){
		fmt.Print(line[len(line)-1])
	}
}

func part2(){
	file, err := os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent := string(file)
	parts := strings.Split(fileContent, "\n\n")
	crate_group := strings.Split(parts[0], "\n")

	stacks := [][] string{}
	for _, line := range(crate_group){
		stack_idx := 0
		for idx:=1; idx<len(line);idx+=4{
			element:=string(line[idx])
			if(stack_idx >= len(stacks)){
				if(element != " "){
					stacks = append(stacks, []string{element})
				} else{
					stacks = append(stacks, []string{})
				}
				
			} else{
				if(element != " "){
					stacks[stack_idx] = append(stacks[stack_idx], element)
				}
			}
			stack_idx++
		}
	}
	
	for _,s:= range(stacks){
		slices.Reverse(s)
	}
	steplines:=strings.Split(parts[1], "\n")
	for _,line:= range(steplines){
		markers := strings.Split(line, " ")
		if(len(markers)<5){
			continue
		}
		c, err:=strconv.ParseInt(markers[1],10,64)
		nilcheck.CheckNilErr(err)
		a, err:=strconv.ParseInt(markers[3],10,64)
		nilcheck.CheckNilErr(err)
		b, err:=strconv.ParseInt(markers[5],10,64)
		nilcheck.CheckNilErr(err)
		a-- 
		b--
		inTransit:= stacks[a][len(stacks[a])-int(c):]
		// slices.Reverse(inTransit)
		stacks[a]=stacks[a][:len(stacks[a])-int(c)]
		stacks[b] = append(stacks[b], inTransit...)
	}
	for _, line:= range(stacks){
		fmt.Print(line[len(line)-1])
	}
}

func main(){
	part1()
	fmt.Println()
	part2()
}