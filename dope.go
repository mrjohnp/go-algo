package dope

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
)

// Least Common Multiple of two numbers
// Dependency: PrimeFactorization()
func lcm(a int, b int) int {
	var fa map[int]int = PrimeFactorization(int(a))
	var fb map[int]int = PrimeFactorization(int(b))
	var res map[int]int = make(map[int]int)
	var lcm int = 1

	for k, v := range fa {
		if fb[k] == 0 {
			res[k] = v
		} else {
			if v > fb[k] {
				res[k] = v
			} else {
				res[k] = fb[k]
			}
		}
	}

	for k, v := range fb {
		if res[k] == 0 {
			res[k] = v
		}
	}

	for k, v := range res {
		lcm *= int(math.Pow(float64(k), float64(v)))
	}

	return lcm
}

// Greatest Common Divisor of two numbers
// Dependency: PrimeFactorization()
func gcd(a int, b int) int {
	var fa map[int]int = PrimeFactorization(int(a))
	var fb map[int]int = PrimeFactorization(int(b))
	var res map[int]int = make(map[int]int)
	var gcd int = 1

	for k, v := range fa {
		if fb[k] > 0 {
			if v < fb[k] {
				res[k] = v
			} else {
				res[k] = fb[k]
			}
		}
	}

	for k, v := range fb {
		if res[k] > 0 {
			if v < res[k] {
				res[k] = v
			}
		}
	}

	for k, v := range res {
		gcd *= int(math.Pow(float64(k), float64(v)))
	}

	return gcd
}

// Prime Factorization of a number n
// @return Map => {2: 3} => [2 == base] :: [3 == exponential]
// Dependency: PrimeNumbers()
// @param > n int: number to factorize
func PrimeFactorization(n int) map[int]int {
	const limit int = 1000
	var primes []int = PrimeNumbers(20)
	var counter int = 0
	var res map[int]int = make(map[int]int)

	if big.NewInt(int64(n)).ProbablyPrime(0) {
		res[n] = 1
	} else {
		for n != 1 && counter < limit {
			counter += 1

			if big.NewInt(int64(n)).ProbablyPrime(0) {
				if res[n] == 0 {
					res[n] = 1
				} else {
					res[n] += 1
				}
				n = 1
				break
			} else {
				for i := range primes {
					if n%primes[i] == 0 {
						if res[primes[i]] == 0 {
							res[primes[i]] = 1
						} else {
							res[primes[i]] += 1
						}
						n = n / primes[i]
						break
					}
				}
			}
		}
	}

	return res
}

// NOTES: Consider using a closure
// to avoid a static number of primes.
// So you can just re-call the function to add the next prime number in the array.
// Print the first n Prime Numbers
// @param > n int: how many prime numbers
func PrimeNumbers(n int) []int {
	var primes []int
	primes = append(primes, 2)

	for i := 3; len(primes) < n; i++ {
		var isPrime bool = big.NewInt(int64(i)).ProbablyPrime(0)
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

// This is to check the number and types of the function's parameters
// And the number an types of it's output
// @param > m interface: function/method to analyse
func FuncAnalyse(m interface{}) {
	x := reflect.TypeOf(m)

	nInPar := x.NumIn()   // Number of input (inbound) parameters
	nOutPar := x.NumOut() // Number of output/return-value (outbound) parameters

	fmt.Println("Method:", x.String())
	fmt.Println("Variadic:", x.IsVariadic())
	fmt.Println("Package:", x.PkgPath())

	for i := 0; i < nInPar; i++ {
		inV := x.In(i)
		inKind := x.Kind()
		fmt.Println("Parameter IN:", strconv.Itoa(i))
		fmt.Println("Kind:", inKind)
		fmt.Println("Name:", inV.Name())
	}

	for o := 0; o < nOutPar; o++ {
		returnV := x.Out(0)
		returnKind := returnV.Kind()
		fmt.Println("Parameter OUT:", strconv.Itoa(o))
		fmt.Println("Kind:", returnKind)
		fmt.Println("Name:", returnV.Name())
	}
}
