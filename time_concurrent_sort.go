package main

import(
	"fmt"
	"math/rand"
	"time"
	"sync"
)

const(
	totalNumbers = 100000
)

var wg sync.WaitGroup

//**************************
//--Main--
//**************************
func main(){

	//Make initial slices
	sliceOne := buildSlice()
	time.Sleep(2000)
	sliceTwo := buildSlice()

	//Make two copies
	sliceCopyOne := make([]int, len(sliceOne))
	copy(sliceCopyOne, sliceOne)
	sliceCopyTwo := make([]int, len(sliceTwo))
	copy(sliceCopyTwo, sliceTwo)
	
	
	//**************DEBUG**************
	//fmt.Println("PRE SLICE ONE: ", sliceOne)
	//fmt.Println("PRE SLICE TWO: ", sliceTwo)
	
	
	//Sort two standard
	start := time.Now()
	wg.Add(2)
	sortSlice(sliceOne)
	sortSlice(sliceTwo)
	endOne := time.Since(start)
	
	
	//**************DEBUG**************
	//fmt.Println("POST SLICE ONE: ", sliceOne)
	//fmt.Println("POST SLICE TWO: ", sliceTwo)
	//fmt.Println("COPY SLICE ONE: ", sliceCopyOne)
	//fmt.Println("COPY SLICE TWO: ", sliceCopyTwo)
			
			
	//Sort two concurrent
	startTwo := time.Now()	
	wg.Add(2)
	go sortSlice(sliceCopyOne)
	go sortSlice(sliceCopyTwo)
	wg.Wait()
	endTwo := time.Since(startTwo)
	
	
	//**************DEBUG**************
	//fmt.Println("POST COPY SLICE ONE: ", sliceCopyOne)
	//fmt.Println("POST COPY SLICE TWO: ", sliceCopyTwo)
	
	delta := (endOne - endTwo)
	
	fmt.Println("Sort task complete.")
	fmt.Println("\nBasic routine run time: ", endOne, "\nGo routine run time: ", endTwo)
	fmt.Println("Time saved using concurrency: ", delta)
}

//**************************
//Create slice of random numbers
//**************************
func buildSlice() []int{
    
	resetRand := rand.NewSource(time.Now().UnixNano())
    rNum := rand.New(resetRand)
	
	unorderedSlice := make([]int, 0, totalNumbers)
	
	i := 0	
	for i < totalNumbers{
		unorderedSlice = append(unorderedSlice, rNum.Intn(100))
		i++
	}
	return unorderedSlice
}

//**************************
//Bubble sort the slice
//**************************
func sortSlice(unorderedSlice[]int){

	defer wg.Done()
	i := 0
	
	for i < totalNumbers{
		indexOne := 0
		indexTwo := 1
		
		for(indexTwo < totalNumbers){
		
			varOne := unorderedSlice[indexOne]
			varTwo := unorderedSlice[indexTwo]
			
			if(varTwo < varOne){
				unorderedSlice[indexOne] = varTwo
				unorderedSlice[indexTwo] = varOne
			}			
			indexOne++
			indexTwo++
		}
		i++
	}
}