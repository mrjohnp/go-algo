package hackerrank

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// path >> text file path
// m >> function/method
// fn >> CHANGE THIS TO TEST OTHER FUNCTIONS/METHODS
func Test(path string, m interface{}, fn func(int32, int32, int32) int32) {
	// Open the file
	inf, err := os.Open(path)
	delErr := os.Remove("output.txt")
	outf, outfErr := os.Create("output.txt")
	resf, resfErr := os.Open("results.txt")

	// Check if there is any error
	if err != nil {
		log.Fatal(err)
	}

	if delErr != nil {
		log.Fatal(delErr)
	}

	if outfErr != nil {
		log.Fatal(outfErr)
	}

	if resfErr != nil {
		log.Fatal(resfErr)
	}

	// Close when the function ends
	defer inf.Close()
	defer outf.Close()
	defer resf.Close()

	// Get type of function parameters
	fncPtr := reflect.ValueOf(m)
	var fncParams []string = getArgTypes(m)

	// Scan every line of the file
	scanner := bufio.NewScanner(inf)
	var inScanLine int32 = 0
	// resScanner := bufio.NewScanner(resf)
	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), " ")
		inScanLine++
		if len(line) == len(fncParams) {
			fncArgs := makeArgs(line, fncParams)
			res := fncPtr.Call(fncArgs)[0]
			// Print the output into the output file
			fmt.Fprintln(outf, res)
			// fmt.Println(res, inScanLine)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR", err)
	}
}

func getArgTypes(fnc interface{}) []string {
	var argTypes []string = []string{}
	fncType := reflect.TypeOf(fnc)
	fncInputNum := fncType.NumIn()
	for i := 0; i < fncInputNum; i++ {
		a := fncType.In(i)
		argTypes = append(argTypes, a.Kind().String())
	}
	return argTypes
}

func makeArgs(argsInput []string, argsTypes []string) []reflect.Value {
	args := make([]reflect.Value, len(argsInput))
	for i, arg := range argsInput {
		for _, argType := range argsTypes {
			switch argType {
			case "int32":
				convArg, convErr := strconv.Atoi(arg)
				if convErr == nil {
					args[i] = reflect.ValueOf(int32(convArg))
				}
				break
			case "int64":
				convArg, convErr := strconv.Atoi(arg)
				if convErr == nil {
					args[i] = reflect.ValueOf(int64(convArg))
				}
				break
			case "int":
				convArg, convErr := strconv.Atoi(arg)
				if convErr == nil {
					args[i] = reflect.ValueOf(convArg)
				}
				break
			default:
				args[i] = reflect.ValueOf(arg)
			}
		}
	}
	return args
}

// Convert 12-hour AM/PM format to military 24-hour time
// in >> "07:05:45PM"
// out >> "19:05:45"
func timeConversion(s string) string {
	var getTime func(c rune) bool = func(c rune) bool {
		return (unicode.IsLetter(c) && !unicode.IsNumber(c)) || unicode.IsPunct(c)
	}

	var time []string = strings.FieldsFunc(s, getTime)
	var period string = s[len(s)-2:]
	var converted string = ""
	var h, _ = strconv.Atoi(time[0])

	if period == "AM" {
		if h == 12 {
			time[0] = "00"
		}
	}

	if period == "PM" {
		if h < 12 {
			time[0] = strconv.Itoa((h + 12))
		}
	}

	converted = strings.Join(time, ":")
	return converted
}

// Check if the two kanagaroos can reach the same location
// in >> x1 = 0, v1 = 2, x2 = 5, v2 = 3
// out >> YES/NO
func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	const max int = 10000
	var res = "NO"

	for i := 0; i < max; i++ {
		x1 = x1 + v1
		x2 = x2 + v2

		if x1 == x2 {
			res = "YES"
			break
		}
	}

	return res
}

// ---- TO-DO ----
// Between Two Sets
func GetTotalX(a []int32, b []int32) int32 {
	// Example: a = [2, 6] b = [24, 36]
	// Numbers Beteween > 6, 12
	// Explanation condition.1: 6%a[0] = 0 :: 6%a[1] = 0
	// Explanation condition.2: b[0]%6 = 0 :: b[1]%6 = 0
	const limit int = 100
	var res []int

	for i := 1; i <= limit; i++ {
		if i%int(a[0]) == 0 && i%int(a[1]) == 0 {
			// fmt.Println("found-b[0]", i, int(b[0])%i)
			// fmt.Println("found-b[1]", i, int(b[1])%i)
			if int(b[0])%i == 0 && int(b[1])%i == 0 {
				res = append(res, i)
			}
		}
	}

	return 0
}

// Apple and Orange
// in >> apples = []int32{2, 3, -4}, oranges = []int32{3, -2, -4}, s = 7, t = 10, a = 4, b = 12
// out >> n-apples && n-oranges
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var resA int32
	var resB int32

	for i := range apples {
		if (apples[i]+a) >= s && (apples[i]+a) <= t {
			resA += 1
		}
	}

	for k := range oranges {
		if (oranges[k]+b) >= s && (oranges[k]+b) <= t {
			resB += 1
		}
	}

	fmt.Println(resA)
	fmt.Println(resB)
}

// Migratory Birds
// in >> sightings = []int32{1, 4, 4, 4, 5, 3}
// out >> n-birds
func migratoryBirds(arr []int32) int32 {
	var m map[int32]int32 = make(map[int32]int32)
	var res int32 = 0
	var ans int32 = 0

	for i := range arr {
		if m[arr[i]] == 0 {
			m[arr[i]] = 1
		} else {
			m[arr[i]] += 1
		}
	}

	for k, v := range m {
		if v > res {
			res = v
			ans = k
		}
	}

	// REFACTOR
	// This solution works only with sorted arrays
	// Dependecy: insertionSort()
	insertionSort(arr)
	var curr int32 = arr[0]
	var poss int32 = 0
	var currCounter int32 = 1
	var possCounter int32 = 0

	for i := 1; i < len(arr); i++ {
		if curr == arr[i] {
			currCounter += 1
		} else {
			if curr != arr[i] && poss != arr[i] {
				poss = arr[i]
				possCounter = 1
			} else if poss == arr[i] {
				possCounter += 1
			}

			if possCounter > currCounter {
				curr = poss
				currCounter = possCounter
			}
		}
	}

	fmt.Println("ANSWER", curr)
	return ans
}

// Insertion Sort Implementation
func insertionSort(arr []int32) []int32 {
	for j := 1; j < len(arr); j++ {
		var curr int32 = arr[j]
		var i = j - 1

		for i > -1 && arr[i] > curr {
			arr[i+1] = arr[i]
			i = i - 1
		}

		arr[i+1] = curr
	}

	return arr
}

// Day of the Programmer
// in >> year = 1984
// out >> date
func dayOfProgrammer(year int32) string {
	var isLeap bool = false
	var m []int32 = []int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var t int32 = 0
	var res string = ""

	if year < 1918 {
		isLeap = year%4 == 0
	} else {
		isLeap = (year%400 == 0) || ((year%4 == 0) && (year%100 != 0))
	}

	if isLeap {
		m[1] = 29
	}

	if year == 1918 {
		m[1] = 15
	}

	for i := range m {
		var tmp int32 = t + m[i]
		if tmp < 256 {
			t += m[i]

			if (256 - t) < m[i+1] {
				var month int32 = int32(i + 2)
				var day int32 = 256 - t
				if len(strconv.Itoa(int(day))) < 2 {
					res += "0"
				}
				res += strconv.Itoa(int(day)) + "."
				if len(strconv.Itoa(int(month))) < 2 {
					res += "0"
				}
				res += strconv.Itoa(int(month)) + "."
				res += strconv.Itoa(int(year))
			}
		}
	}

	return res
}

// Bon Appétit - Bill division
// in >> bill = []int32{3, 10, 2, 9}, k = 1, b = 12
// out >> n-charge/Bon-Appetit
func bonAppetit(bill []int32, k int32, b int32) {
	var total int32 = 0
	for i := range bill {
		if i != int(k) {
			total += bill[i]
		}
	}

	var charge int32 = total / 2
	if charge == b {
		fmt.Println("Bon Appetit")
	} else if b > charge {
		fmt.Println(b - charge)
	}
}

// Drawing Book
// in >> n-pages = 6, target-p = 2
// out >> minimum n-turns
func pageCount(n int32, p int32) int32 {
	var forwardTurns int = 0
	var backwardTurns int = 0
	var res int32 = 0

	for i := 1; i <= int(n); i++ {
		if i == 1 && i == int(p) {
			break
		}

		if (i%2) == 0 && i > 1 {
			forwardTurns += 1
			if i == int(p) || i+1 == int(p) {
				break
			}
		}
	}

	for k := n; k > 0; k-- {
		if k == n && (((k%2) == 0 && k == p) || (k%2) == 1 && (k == p || k-1 == p)) {
			break
		}

		if (k%2) == 1 && k < n {
			backwardTurns += 1
			if k == p || k-1 == p {
				break
			}
		}
	}

	if forwardTurns > backwardTurns {
		res = int32(backwardTurns)
	} else {
		res = int32(forwardTurns)
	}

	return res
}

// Electronics Shop
// in >> budget = 60, keyboardsPrices = []int32{40, 50, 60}, drivesPrices = []int32{5, 8, 12}
// @return int: the maximum that can be spent, or -1 if it is not possible to buy both items
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var res int32 = -1

	for i := range keyboards {
		for k := range drives {
			var price int32 = keyboards[i] + drives[k]

			if price <= b && price > res {
				res = price
			}
		}
	}

	return res
}

// Cats and a Mouse
// in >> x = 2, y = 5, z = 4
// @return string: Either 'Cat A', 'Cat B', or 'Mouse C'
func catAndMouse(x int32, y int32, z int32) string {
	var distX = math.Abs(float64(z - x))
	var distY = math.Abs(float64(z - y))
	var res = ""

	if distX < distY {
		res = "Cat A"
	} else if distY < distX {
		res = "Cat B"
	} else {
		res = "Mouse C"
	}

	return res
}

// ---- TO-DO ----
// Forming a Magic Square
// @return int: the minimal total cost of converting the input square to a magic square
// Note: magic constant == 15
func formingMagicSquare(s [][]int32) int32 {
	for i := range s {
		var rowSum int32 = 0
		row := s[i]

		for k := range row {
			rowSum += row[k]
			fmt.Printf("%d \t", row[k])
			if k == len(row)-1 {
				fmt.Println("Sum", i, rowSum)
			}
		}
	}

	return 0
}

// in >> grades = []int32{73, 67, 38, 33}
// out >> rounded grandes
func gradingStudents(grades []int32) []int32 {
	var res []int32 = []int32{}
	for n := range grades {
		var grade int32 = grades[n]
		if grade < 38 {
			res = append(res, grade)
			continue
		}

		var nextMult int32 = ((grade / 5) + 1) * 5
		var diff int32 = nextMult - grade
		if diff < 3 {
			res = append(res, nextMult)
		} else {
			res = append(res, grade)
		}
	}

	return res
}

// in >> s = []int32{1, 2, 1, 3, 2}, d = 3, m = 2
func birthday(s []int32, d int32, m int32) int32 {
	var res int32 = 0
	for n := range s {
		var curr int32 = s[n]
		var sum int32 = curr

		if m > 1 {
			for i := 1; i < int(m); i++ {
				if n+i < len(s) {
					sum += s[n+i]
				}
			}
		}

		if sum == d {
			res++
		}
	}

	return res
}

// in >> ar-len = 6, k = 3, ar = []int32{1, 3, 2, 6, 1, 2}
// out >> pairs where i < j and ar[i] + ar[j] is evenly divisible by k
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var res int32 = 0
	for i := range ar {
		for j := i + 1; j < int(n); j++ {
			if i < j {
				var sum int32 = ar[i] + ar[j]
				if sum%k == 0 {
					res++
				}
			}
		}
	}
	return res
}

// Longest subarray
// in >> a = []int32{ 4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98 }
// out >> length of the longest subarray where the absolute
// difference between any two elements is less than or equal to 1.
func pickingNumbers(a []int32) int32 {
	var res int32 = 0
	var cand []int32 = []int32{}

	for i := range a {
		var tmp1 []int32 = []int32{}
		var tmp2 []int32 = []int32{}
		var nUp int32 = a[i] + 1
		var nDown int32 = a[i] - 1
		tmp1 = append(tmp1, a[i])
		tmp2 = append(tmp2, a[i])
		for j := i + 1; j < len(a); j++ {
			if i+1 < len(a) {
				if a[j] == nUp {
					tmp1 = append(tmp1, a[j])
				} else if a[j] == nDown {
					tmp2 = append(tmp2, a[j])
				} else if a[j] == a[i] {
					tmp1 = append(tmp1, a[j])
					tmp2 = append(tmp2, a[j])
				}
			}
		}

		var candLen = len(cand)
		var tmp1Len = len(tmp1)
		var tmp2Len = len(tmp2)
		if tmp1Len >= tmp2Len && tmp1Len > candLen {
			cand = make([]int32, tmp1Len)
			copy(cand, tmp1)
		} else if tmp2Len >= tmp1Len && tmp2Len > candLen {
			cand = make([]int32, tmp2Len)
			copy(cand, tmp2)
		}
	}

	res = int32(len(cand))
	return res
}

// in >> prisoner-n = 7, sweets-m = 19, start-s = 2
// out >> the chair number of the prisoner to warn
func saveThePrisoner(n int32, m int32, s int32) int32 {
	var res int32 = s
	if m >= n {
		if m%n == 0 {
			if n-(s-1) == n {
				res = n
			} else if n-(s-1) < n {
				res = s - 1
			}
		} else {
			var pos int32 = s + ((m % n) - 1)
			if pos > n {
				if s+((m%n)-1) > n {
					diff := n - (s - 1)
					res = m%n - diff
				} else {
					res = m % n
				}
			} else {
				res = pos
			}
		}
	} else {
		if m+s > n {
			diff := ((m + s) - n) - 1
			if diff == 0 {
				res = n
			} else {
				res = ((m + s) - n) - 1
			}
		} else {
			res = m - 1 + s
		}
	}

	return res
}

// int a[n]: the array to rotate
// int k: the rotation count
// int queries[1]: the indices to report
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	for i := 0; i < int(k); i++ {
		var x int32 = a[len(a)-1]    // Get the last element from the array
		a = a[:len(a)-1]             // Remove the last element from the array
		a = append([]int32{x}, a...) // Prepend an element to the array
		fmt.Println(a)
	}

	for q := range queries {
		fmt.Println(a[queries[q]])
	}

	return a
}

func saveThePrisonerV2(n int32, m int32, s int32) int32 {
	var res int32 = s
	if m > n {
		var k int32 = m - (n - (s - 1))
		if k < n {
			res = k
		} else if k == n {
			res = n
		} else if k > n {
			res = k % n
		}
	}
	return res
}
