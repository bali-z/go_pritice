package main

import (
	"fmt"
)

// map 相关 
func main(){
	sMap := make(map[int]string)
	sMap[1] = "111"
	sMap[2] = "222"

	for k,v := range sMap {
		fmt.Printf("k is %v, v is %v \n",k,v)
	}
}