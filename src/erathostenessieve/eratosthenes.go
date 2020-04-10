package eratosthenessieve

import (
	"fmt"
	"math/big"
)

// SieveOfErathostenes : a big int version of the sieve of Erathostenes
type SieveOfErathostenes struct {
	limit          *big.Int
	primeStorer func(*big.Int)
}

func newSieveOfErathostenes(limit *big.Int, primeStorer func(*big.Int)) *SieveOfErathostenes {
	sieve := new(SieveOfErathostenes)
	sieve.limit = limit
	sieve.primeStorer = primeStorer
	return sieve
}

// Run the sieve of erathostenes get all primes lower than n (BigInt version)
func (sieve SieveOfErathostenes) Run() {
	one := big.NewInt(int64(1))
	two := big.NewInt(int64(2))
	isPrime := make(map[string]bool)
	for i := two; i.Cmp(sieve.limit) < 0; i = new(big.Int).Add(i, one) {
		fmt.Printf("Number %s\n", i.String())
		iIsPrime, iPrimalityIsComputed := isPrime[i.String()]
		if iIsPrime || !iPrimalityIsComputed {
			jStart := new(big.Int).Exp(i, two, nil)
			for j := jStart; j.Cmp(sieve.limit) < 0; j = new(big.Int).Add(j, i) {
				isPrime[j.String()] = false
			}
			sieve.primeStorer(i)
			delete(isPrime, i.String())
		}
	}
}
