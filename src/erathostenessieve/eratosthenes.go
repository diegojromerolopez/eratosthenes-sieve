package eratosthenessieve

import (
	"math/big"
)

// SieveOfErathostenes : a big int version of the sieve of Erathostenes
type SieveOfErathostenes struct {
	limit       *big.Int
	primeStorer func(*big.Int)
}

func newSieveOfEratosthenes(limit *big.Int, primeStorer func(*big.Int)) *SieveOfErathostenes {
	sieve := new(SieveOfErathostenes)
	sieve.limit = limit
	sieve.primeStorer = primeStorer
	return sieve
}

// Run the sieve of erathostenes get all primes lower than n (BigInt version)
func (sieve SieveOfErathostenes) Run() {
	two := big.NewInt(int64(2))
	three := big.NewInt(int64(3))
	isPrime := make(map[string]bool)
	// Two is the only even number that can be prime
	if sieve.limit.Cmp(two) >= 0 {
		sieve.primeStorer(two)
	}

	// Odd numbers: starting in 3, loop increasing by two
	for i := three; i.Cmp(sieve.limit) <= 0; i = new(big.Int).Add(i, two) {
		iString := i.String()
		iIsPrime, iPrimalityIsComputed := isPrime[iString]
		if iIsPrime || !iPrimalityIsComputed {
			jStart := new(big.Int).Exp(i, two, nil)
			for j := jStart; j.Cmp(sieve.limit) <= 0; j = new(big.Int).Add(j, i) {
				isPrime[j.String()] = false
			}
			sieve.primeStorer(i)
		}
		delete(isPrime, iString)
	}
}
