package main

import (
	nilcheck "advent_of_code/commonUtils"
	"fmt"
	"os"
	"strings"
)

func part1(){
	file, err := os.ReadFile("input.txt")
	nilcheck.CheckNilErr(err)
	fileContent:=strings.TrimSpace(string(file))
	hashMap := map[byte]int{}
	for j:=0;j<len(fileContent);j++{
		if(j<14){
			hashMap[fileContent[j]]++;
		} else{
			hashMap[fileContent[j]]++;
			hashMap[fileContent[j-14]]--;
			if(hashMap[fileContent[j-14]] == 0){
				delete(hashMap,fileContent[j-14])
			}
		}
		if(len(hashMap)==14){
			fmt.Println(j+1)
			break
		}
	}
}

func main(){
	part1()
}