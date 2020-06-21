function sumOfPrimeDistance(arr) {
    const pairs = getPairs2(arr);
    const primed = calculatePrimes(pairs);

    const length = arr.length - 1;
    const total = length * length / 2

    let step = 1;
    while (step < total) {
        for (let i = 0; i < primed.length - 1; ++i) {
            const left = primed[i];
            const right = primed[i + step];

            const merged = merge(left, right);
            primed.push(merged)
        }

        step++;
    }

    return sum;
    //
    // for (let i = 0; i < arr.length - 1; i++) {
    //     acc.push([current, original[step]])
    // }


    // return primed;

    //     const length = arr.length - 1;
    //     const total = length * length / 2

    // const pairs = getPairs(arr);
    // return pairs.reduce((acc, [left, right]) => {
    //     const primes = getPrimesInInterval(left, right);
    //     return acc + primes.length;
    // }, 0)
}

const merge = (left, right) => {
    const [ll, lr, lp] = left;
    const [rl, rr, rp] = right;

    return [ll, rr, lp + rp];
}

const getPairs2 = arr => {
    const pairs = [];
    for (let i = 0; i < arr.length - 1; ++i) {
        const current = arr[i];
        const next = arr[i + 1];

        pairs.push([current, next])
    }

    return pairs;
}

const calculatePrimes = pairs => {
    return pairs.map(([left, right]) => {
        const primes = getPrimesInInterval(left, right);
        return [left, right, primes.length]
    });
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
