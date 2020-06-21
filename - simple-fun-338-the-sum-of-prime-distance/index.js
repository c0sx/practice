function sumOfPrimeDistance(arr) {
    const pairs = getPairs(arr);

    return pairs.reduce((acc, [left, right]) => {
        const primes = getPrimesInInterval(left, right);
        return acc + primes.length;
    }, 0)
}

const getPairs = arr => {
    return arr.reduce((acc, current, index, original) => {
        for (let step = index + 1; step < original.length; step++) {
            acc.push([current, original[step]])
        }

        return acc;
    }, []);
}

const getPrimesInInterval = (left, right) => {
    const primes = []

    let i = left + 1
    while (i <= right) {
        if (isPrime(i)) {
            primes.push(i)
        }

        i++;
    }

    return primes;
}

const cache = new Map();
function isPrime(x) {
    const cached = cache.get(x);
    if (cached !== undefined) {
        return cached
    }

    const size = Math.trunc(Math.pow(x, 1 / 2));
    for (let i = 2; i <= size; ++i) {
        if (x % i === 0) {
            cache.set(x, false);
            return false;
        }
    }

    cache.set(x, true);
    return true;
}

module.exports = sumOfPrimeDistance
