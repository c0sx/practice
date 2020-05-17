function sumOfDivided(lst) {
    const primes = lst.reduce((acc, current) => {
        const primes = getPrimes(current);
        acc.push(...primes);
        return acc;
    }, []);

    const normalizedPrimes = [...new Set(primes)].sort((a, b) => a - b);

    return normalizedPrimes.reduce((acc, current) => {
        const total = lst.reduce((sum, num) => {
            const divide = Math.trunc(num / current);
            if (divide * current === num) {
                return sum + num;
            }

            return sum;
        }, 0)

        acc.push([current, total])
        return acc;
    }, []);
}

function getPrimes(x) {
    let n = Math.abs(x);
    let divider = 2;

    const primes = [];
    while (n > 1) {
        divider = getDividerForX(n, divider);
        primes.push(divider)
        n = n / divider;
    }

    return primes;
}

function getDividerForX(x, divider) {
    while (x % divider !== 0) {
        divider = nextPrime(divider)
    }

    return divider;
}

function nextPrime(x) {
    while (!isPrime(x += 1)) {
        //
    }

    return x;
}

function isPrime(x) {
    const size = Math.trunc(Math.pow(x, 1 / 2));
    for (let i = 2; i <= size; ++i) {
        if (x % i === 0) {
            return false;
        }
    }

    return true;
}

module.exports = sumOfDivided;



