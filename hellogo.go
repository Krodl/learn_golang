package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	passingArraysUsingPointers()
}

func ioExamples() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}

func loop() {
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}

func timeExample1() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}

func forLoop() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50 :")
		pl("Randon Number is :", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Lower")
		} else if iGuess < randNum {
			pl("Higher")
		} else {
			pl("You guessed it")
			break
		}
	}
}

func forLoopRange() {
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}
}

func arraysExample1() {
	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0 :", arr2[0])
	pl("Arr Length :", len(arr2))
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d: %d\n", i, v)
	}
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}
}

func arraysExample2() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)
}

// Slices are like arrays but they can grow (similar to Java ArrayList)
func sliceExample() {
	// var name []datatype
	sl1 := make([]string, 5)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Psycho"
	sl1[4] = "Helmet"
	pl("Slice Size :", len(sl1))

	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}

	// Slices are parts of an array. Making changes in the array will reflect in the slice and vice versa.
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl("1st 3 :", sArr[:3])
	pl("Last 3 :", sArr[2:])
	sArr[0] = 10
	pl("sl3 :", sl3)
	sl3[0] = 1
	pl("sArr :", sArr)

	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])

}

func passingPointers() {
	f3 := 1

	// Comment out one or the other, when we call changeVal before changeVal2 go sets the
	// Type of f3 specifically to int and won't interpret it as an *int for changeVal
	// TODO: Don't know why

	// changeVal(f3)
	// pl("Value doesn't change due to function call :", f3)

	pl("Value before function call :", f3)
	changeVal2(&f3) // & is "Address of" operator.
	pl("Value DOES change due to void function call :", f3)
}

func storingPointers() {
	f4 := 10
	var f4Ptr *int = &f4

	pl("f4 Address :", f4Ptr)
	pl("f4 Value :", *f4Ptr)

	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	pl("f4 before func :", f4)
	changeVal2(&f4)
	pl("f4 after func :", f4)

}

func passingArraysUsingPointers() {
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)
}

func passingSliceUsingPointers() {
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}

// Takes in unspecified size of an array. This makes it a slice.
func getAverage(nums ...float64) float64 {
	pl(reflect.TypeOf(nums))
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, v := range nums {
		sum += v
	}
	return (sum / numSize)
}

func dblArrVals(arr *[4]int) {
	pl(reflect.TypeOf(arr))
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

// Doesn't use pointer, won't change value outside of function
func changeVal(f3 int) {
	f3 = 10
}

// Uses pointer. WILL change value outside of function
func changeVal2(f3 *int) {
	*f3 = 10
}
