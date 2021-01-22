function sumOfPrimeDistance(arr) {
    const primes = getPrimes(arr);

    let row = 0;
    let sum = 0;
    while (row < primes.length / 2 - 1) {
        sum += iteration(primes, row);
        row++;
    }

    const peak = iteration(primes, row)
    if (arr.length % 2 === 0) {
        return sum * 2 + peak;
    }

    return sum * 2 + peak * 2
}

const iteration = (arr, step) => {
    let sum = 0;
    for (let i = 0; i < arr.length - step; ++i) {
        for (let j = i; j < i + step + 1; ++j) {
            sum += arr[j]
        }
    }

    return sum;
}

const getPrimes = arr => {
    const primes = [];
    for (let i = 0; i < arr.length - 1; ++i) {
        const current = arr[i];
        const next = arr[i + 1];

        const length = getPrimesInInterval(current, next)
        primes.push(length);
    }

    return primes;
}

const getPrimesInInterval = (left, right) => {
    let primes = 0;

    let i = left + 1
    while (i < right) {
        if (isPrime(i)) {
            primes++;
        }

        i++;
    }

    return primes;
}

const isPrime = x => {
    const size = Math.trunc(Math.pow(x, 1 / 2));
    for (let i = 2; i <= size; ++i) {
        if (x % i === 0) {
            return false;
        }
    }

    return true;
}

module.exports = sumOfPrimeDistance
