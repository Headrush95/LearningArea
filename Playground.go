package main

// Не удалять, алгоритм в разработке!!!!
import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
)

// Set - моя имплементация hashSet
type Set map[string]struct{}

// Add добавляет элемент в Set
func (s Set) Add(n string) {
	s[n] = struct{}{}
}

// Contains проверяет наличие элемента в Set
func (s Set) Contains(n string) bool {
	_, ok := s[n]
	return ok
}

// Remove удаляет элемент из Set
func (s Set) Remove(n string) {
	delete(s, n)
}

// Print печатает содержимое Set
func (s Set) Print() {
	for k := range s {
		fmt.Print(k, " ")
	}
}

// NextPrime Выдает следующее простое число
func NextPrime(arr *[]bool, I int) int {
	for i := I; i < len(*arr); i++ {
		if !(*arr)[i] {
			return i
		}
	}
	return -1
}

// CollectPrimes возвращает список простых чисел в формате Set
func CollectPrimes(arr *[]bool) Set {
	Primes := make(Set)
	for i, isPrime := range *arr {
		if !isPrime && i > 10 {
			Primes.Add(strconv.Itoa(i))
		}
	}
	return Primes
}

// GetPrimes ищет простые числа с помощью решета Эратосфена
func GetPrimes(n int) Set {
	sieve := make([]bool, n+1, n+1)
	sieve[0], sieve[1] = true, true

	for i := 4; i <= n; {
		sieve[i] = true
		i += 2
	} // помечаем кратные 2-м

	i := NextPrime(&sieve, 3)
	if i == -1 {
		log.Fatal(fmt.Errorf("something went wrong"))
	} //error

	sqrtN := int(math.Sqrt(float64(n)))

	for i <= sqrtN {
		for j := i * 2; j <= n; {
			if !sieve[j] {
				sieve[j] = true
			}
			j += i
		}

		i = NextPrime(&sieve, i+1)
		if i == -1 {
			log.Fatal(fmt.Errorf("something went wrong"))
		} // error
	}

	//PrintArr(&sieve)

	return CollectPrimes(&sieve)
}

// Reverse разворачивает строку
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SearchEmirp(s Set, n int) []int {
	var revNum string
	Emirp := make([]int, 0, len(s)/10)
	for num := range s {
		revNum = Reverse(num)
		if s.Contains(revNum) && num != revNum {
			fmt.Printf("[SearchEmirp] Num = %s\n", num)
			fmt.Printf("[SearchEmirp] RevNum = %s\n", revNum)
			tmp, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(fmt.Errorf("something went wrong while converting: ", err))
			} // error
			Emirp = append(Emirp, tmp)
			tmp, _ = strconv.Atoi(revNum)
			if tmp <= n {
				Emirp = append(Emirp, tmp)
			}
		}

	}
	sort.Ints(Emirp)
	return Emirp
}

func main() {
	n := 50
	Primes := GetPrimes(n)
	Primes.Print()
	fmt.Println()
	fmt.Println(len(Primes))
	fmt.Println()
	fmt.Println(SearchEmirp(Primes, n))

}
