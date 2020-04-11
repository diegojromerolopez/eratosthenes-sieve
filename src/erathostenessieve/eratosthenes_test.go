package eratosthenessieve

import (
	"math/big"
	"testing"
	"time"
)

func TestSieveOfEratosthenes(t *testing.T) {
	n := big.NewInt(31)

	var primes []*big.Int
	primeStorer := func(prime *big.Int) {
		primes = append(primes, prime)
	}

	startTime := time.Now()

	sieve := newSieveOfEratosthenes(n, primeStorer)
	sieve.Run()

	elapsedTime := time.Since(startTime)

	expectedPrimes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	if len(expectedPrimes) != len(primes) {
		t.Errorf("Prime count failed on: %d primes expected, %d found.",
			len(expectedPrimes), len(primes))
	}

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
		t.Logf("Elapsed time: %s", elapsedTime)
	} else {
		t.Errorf("Prime count failed on: %d primes", badPrimesCount)
	}

}

func TestSieveOfEratosthenesFor2(t *testing.T) {
	n := big.NewInt(2)

	var primes []*big.Int
	primeStorer := func(prime *big.Int) {
		primes = append(primes, prime)
	}

	startTime := time.Now()

	sieve := newSieveOfEratosthenes(n, primeStorer)
	sieve.Run()

	elapsedTime := time.Since(startTime)

	expectedPrimes := []int64{2}
	if len(expectedPrimes) != len(primes) {
		t.Errorf("Prime count failed on: %d primes expected, %d found.",
			len(expectedPrimes), len(primes))
	}

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
		t.Logf("Elapsed time: %s", elapsedTime)
	} else {
		t.Errorf("Prime count failed on: %d primes", badPrimesCount)
	}

}

func TestSieveOfEratosthenesPerformance(t *testing.T) {
	n, ok := new(big.Int).SetString("10000000", 10)
	if !ok {
		t.Errorf("Invalid 10-base number found.")
		return
	}

	var primes []*big.Int
	primeStorer := func(prime *big.Int) {
		primes = append(primes, prime)
	}

	startTime := time.Now()

	sieve := newSieveOfEratosthenes(n, primeStorer)
	sieve.Run()

	elapsedTime := time.Since(startTime)

	t.Logf("Elapsed time: %s", elapsedTime)
}
