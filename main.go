package main

import (
	"fmt"
	"sort"
)

func Fibonacci(n int) {
	f1 := 0
	f2 := 1
	var f3 int

	fmt.Print("\n Fibonacci serie is: ", f1, f2)

	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3

		fmt.Print(" ", f3)
	}
	fmt.Println()
}

func SortArrayWithSortPackage(arr []int) {
	sort.Ints(arr)
	fmt.Println("The sorted array is: ", arr)
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("The sorted array with SelectionSort is: ", arr)
}

func MaxElement(arr []int) int {
	maxim := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxim {
			maxim = arr[i]
		}
	}
	return maxim
}

func FrequencyArray(arr []int, helper []int) {
	for i := 0; i < len(arr); i++ {
		helper[arr[i]]++
	}
}

func FindDublicateElements(arr []int) {

	maximArray := MaxElement(arr)

	//helper array is like frequency array in c++
	helper := make([]int, maximArray+1)

	FrequencyArray(arr, helper)

	for i := 0; i < len(helper); i++ {
		if helper[i] >= 2 {
			fmt.Printf("The element %v appears %v times.\n", i, helper[i])
		}
	}
}

func DeletionWithNewArray(arr []int, helper []int) {
	nr := 0
	for j := 0; j < len(helper); j++ {
		if helper[j] == 1 {
			nr++
		}
	}
	newArray := make([]int, nr)
	k := 0
	for i := 0; i < len(arr); i++ {
		if helper[arr[i]] == 1 {
			newArray[k] = arr[i]
			k++
		}
	}
	fmt.Println("The new array after the deletions: ", newArray)
}

func DeletionFromArray(arr []int, helper []int) {
	//O(n^2)
	n := len(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		if helper[arr[i]] >= 2 {
			for j := i; j < len(arr)-1; j++ {
				arr[j] = arr[j+1]
			}
			n--
		}
	}

	var newArray []int
	newArray = append(newArray, arr[:4]...)

	fmt.Println("The new array after the deletions: ", newArray)
}

func DeleteDublicateElements(arr []int, method int) {
	maximArray := MaxElement(arr)

	//helper array is like frequency array in c++
	helper := make([]int, maximArray+1)

	FrequencyArray(arr, helper)

	switch method {
	case 1:
		//The variant with adding the elements that interest me in a new array to avoid O(n^2)
		DeletionWithNewArray(arr, helper)
	case 2:
		DeletionFromArray(arr, helper)
	default:
		fmt.Println("There is no other implementation for method ", method)
	}
}

func FindLongestPalindromicSubstring(str string) {
	//a palindrome can have even length or odd length. For a certain fixed position i, it is checked if there is a
	//palindrome of odd length centered in position i and then a palindrome of even length centered in positions i and i+1
	b := []byte(str)

	lgmax := 1

	var newString []byte

	var copie_p int
	var copie_q int

	for i := 1; i < len(b); i++ {
		//search for palindrome of odd length  centered in position i
		p := i - 1
		q := i + 1

		for p >= 0 && q < len(b) && b[p] == b[q] {
			p--
			q++
		}

		p++
		q--

		if lgmax < q-p+1 {
			lgmax = q - p + 1
			copie_p = p
			copie_q = q
		}

		//search for palindrome of even length centered in position i and i+1
		p = i
		q = i + 1
		for p >= 0 && q < len(b) && b[p] == b[q] {
			p--
			q++
		}

		q--

		if lgmax < q-p+1 {
			lgmax = q - p + 1
			copie_p = p
			copie_q = q
		}
	}

	newString = append(newString, b[copie_p:copie_q+1]...)
	if lgmax == 1 {
		fmt.Println("There are no palindromes")
	} else {
		fmt.Println("The longest palindromic substring is: ", string(newString[:]))
	}
}

func FindElementMissing(arr []int) {
	maximArray := MaxElement(arr)
	fmt.Printf("\nThe array has elements in range [ 0 , %v ].", maximArray)

	helper := make([]int, maximArray+1)
	FrequencyArray(arr, helper)

	fmt.Print("\nThe element that are missing is: ")
	for i := 0; i < len(helper); i++ {
		if helper[i] == 0 {
			fmt.Printf("%v ", i)
			break //because is just one element that are missing.
		}
	}
	fmt.Println()
}

func main() {

	var ex int
	fmt.Print("Enter the exercise number: ")
	fmt.Scan(&ex)
	fmt.Println()

	switch ex {

	case 1:
		fmt.Println("----------- Exercise 1 -----------")
		fmt.Println()
		var n int
		fmt.Print("Enter how many elements of the Fibonacci serie you want to generate: ")
		fmt.Scan(&n)
		Fibonacci(n)

	case 2:
		fmt.Println("----------- Exercise 2 -----------")
		array := []int{20, 30, 10, 50, 90, 70, 60, 80, 40}
		fmt.Println("Unsorted array: ", array)
		var method int
		fmt.Print("\nEnter the method with which you want to sort the array: \n 1 for method with package sort \n 2 for method with Selection Sort ")
		fmt.Scan(&method)
		fmt.Println()
		switch method {
		case 1:
			SortArrayWithSortPackage(array)
		case 2:
			SelectionSort(array)
		default:
			fmt.Println("There is no method with number ", method)
		}

	case 3:
		fmt.Println("----------- Exercise 3 -----------")
		fmt.Println()
		array := []int{5, 4, 5, 5, 3, 1, 3}
		fmt.Println("The array is: ", array)
		FindDublicateElements(array)

	case 4:
		fmt.Println("----------- Exercise 4 -----------")
		fmt.Println()
		array := []int{5, 10, 10, 10, 2, 1, 1, 2, 3, 4, 8, 8, 7}
		fmt.Println("The array is: ", array)
		var method int
		fmt.Print("\nEnter the method with which you want to delete the duplicate elements: \n 1 for method with a new array \n 2 for method with delete, complexity O(n^2) ")
		fmt.Scan(&method)
		fmt.Println()
		DeleteDublicateElements(array, method)

	case 5:
		fmt.Println("----------- Exercise 5 -----------")
		fmt.Println()
		str := "asdanaminim"
		fmt.Println("The string is: ", str)
		FindLongestPalindromicSubstring(str)
		str1 := "capacrotitormagnific"
		fmt.Println()
		fmt.Println("Another string is: ", str1)
		FindLongestPalindromicSubstring(str1)

	case 6:
		fmt.Println("----------- Exercise 6 -----------")
		fmt.Println()
		array := []int{0, 5, 2, 3, 9, 1, 4, 6, 8}
		fmt.Println("The array is: ", array, "In this array, just one element is missing.")
		FindElementMissing(array)

	default:
		fmt.Println("There is no exercise with number: ", ex)
	}

}
