package eratosthenessieve

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestSieveOfErathostenesInt(t *testing.T) {
	primes := SieveOfErathostenesInt(int64(33))
	expectedPrimes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	if reflect.DeepEqual(primes, expectedPrimes) {
		t.Logf("Prime count OK")
	} else {
		t.Errorf("Prime fucked: %s", fmt.Sprint(primes))
	}
}

func TestSieveOfErathostenes(t *testing.T) {
	primes := SieveOfErathostenes(big.NewInt(33))
	expectedPrimes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	badPrimesCount := 0
	for i := 0; i < len(primes); i++ {
		expectedPrimeI := big.NewInt(int64(expectedPrimes[i]))
		if expectedPrimeI.Cmp(primes[i]) != 0 {
			t.Errorf("Prime %d not found. Found %s instead",
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
