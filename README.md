# Eratosthenes' Sieve

An implemententation of the
[Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
in Golang.

# Use

Use eratosthenessieve.SieveOfErathostenes

```go
// Use an auxiliar primeStorer function that will
// store each one of the prime numbers found
var primes []*big.Int
primeStorer := func(prime *big.Int) {
    primes = append(primes, prime)
}

// Create a new Sieve of Eratosthenes and run it
sieve := newSieveOfEratosthenes(n, primeStorer)
sieve.Run()
```

## Contact me

Contact me at diegojromerolopez at gmail dot com

## License

[MIT](LICENSE)
