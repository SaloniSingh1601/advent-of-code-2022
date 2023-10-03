package main

import (
	"advent_of_code/commonUtils"
	"fmt"
	"os"
	"strings"
)



func part1(){
	file,err:=os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent := string(file)
	priority := 0
	rudderstacks:= strings.Split(fileContent, "\n")
	for _, rudderstack := range(rudderstacks){
		hashMap := map[byte] int{}
		size := len(rudderstack)
		for i:=0;i<size/2;i++{
			hashMap[rudderstack[i]]++;
		}
		for i:=size/2; i<size;i++{
			if(hashMap[rudderstack[i]] != 0){
				if(rudderstack[i]>96){
					
					priority += (int(rudderstack[i])-96)
				} else{
					priority += (int(rudderstack[i]) -64 + 26)
				}
				break;
			}
		}
	}
	fmt.Println(priority)
}

func part2(){
	file, err:= os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent := string(file)
	priority := 0
	rudderstacks := strings.Split(fileContent, "\n")
	size := len(rudderstacks)
	for i:=0;i<size;{
		hashMap := map[byte] int{}
		for _,ele:= range(rudderstacks[i]){
			if(hashMap[byte(ele)]==0){
				hashMap[byte(ele)]++;
			}
		}
		i++
		for _,ele:= range(rudderstacks[i]){
			if(hashMap[byte(ele)]==1){
				hashMap[byte(ele)]++;
			}
		}
		i++;
		for _,ele:= range(rudderstacks[i]){
			if(hashMap[byte(ele)]==2){
				if(byte(ele)>96){
					
					priority += (int(byte(ele))-96)
				} else{
					priority += (int(byte(ele)) -64 + 26)
				}
				break;
			}
		}
		i++;
	}
	fmt.Println(priority)
}

func main(){
	part1()
	part2()
}