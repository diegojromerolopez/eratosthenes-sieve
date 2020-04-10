package eratosthenessieve

import (
	"math/big"
	"testing"
)

func TestSieveOfErathostenes(t *testing.T) {
	n := big.NewInt(31)

	var primes []*big.Int
	primeStorer := func(prime *big.Int) {
		primes = append(primes, prime)
	}

	sieve := newSieveOfErathostenes(n, primeStorer)
	sieve.Run()

	expectedPrimes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	if len(expectedPrimes) != len(primes) {
		t.Errorf("Prime count failed on: %d primes expected, %d found.",
			len(expectedPrimes), len(primes))
	} else {
		badPrimesCount := 0
		for i := 0; i < len(primes); i++ {
			expectedPrimeI := big.NewInt(int64(expectedPrimes[i]))
			if expectedPrimeI.Cmp(primes[i]) != 0 {
				t.Errorf("Prime %d not found. Found %s instead.",
					expectedPrimes[i], primes[i].String())
				badPrimesCount++
			}
		}
		if badPrimesCount == 0 {
			t.Logf("Prime OK")
		} else {
			t.Errorf("Prime count failed on: %d primes", badPrimesCount)
		}
	}
}
