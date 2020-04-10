package eratosthenessieve

import (
	"math"
	"math/big"
)

// SieveOfErathostenesInt : get all primes lower than n (Int64 version)
func SieveOfErathostenesInt(n int64) []int64 {
	one := int64(1)
	two := int64(2)
	isPrime := make(map[int64]bool)
	var primes []int64
	for i := two; i < n; i = i + one {
		iIsPrime, iPrimalityIsComputed := isPrime[i]
		if iIsPrime || !iPrimalityIsComputed {
			for j := int64(math.Pow(float64(i), float64(two))); j < n; j = j + i {
				isPrime[j] = false
			}
			primes = append(primes, i)
		}
	}
	return primes
}

// SieveOfErathostenes : get all primes lower than n (BigInt version)
func SieveOfErathostenes(n *big.Int) []*big.Int {
	one := big.NewInt(int64(1))
	two := big.NewInt(int64(2))
	isPrime := make(map[string]bool)
	var primes []*big.Int
	for i := two; i.Cmp(n) < 0; i = new(big.Int).Add(i, one) {
		iIsPrime, iPrimalityIsComputed := isPrime[i.String()]
		if iIsPrime || !iPrimalityIsComputed {
			jStart := new(big.Int).Exp(i, two, nil)
			for j := jStart; j.Cmp(n) < 0; j = new(big.Int).Add(j, i) {
				isPrime[j.String()] = false
			}
			primes = append(primes, i)
		}
	}
	return primes
}
