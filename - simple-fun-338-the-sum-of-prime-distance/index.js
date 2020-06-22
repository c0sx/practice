function sumOfPrimeDistance(arr) {
    const pairs = getPairs2(arr);
    let primed = calculatePrimes(pairs);

    let total = sumRow(primed);
    let step = 1;

    while (step < primed.length) {
        for (let i = 0; i < primed.length - step; ++i) {
            const [left] = primed[i];
            const [_, right] = primed[i + step];

            const slice = merge2(primed, left, right);
            total += sumRow(slice)
        }

        step++;
    }

    return total;
}

const sumRow = row => row.reduce((acc, [l, r, p]) => acc + p, 0)

const merge2 = (row, left, right) => {
    const start = row.findIndex(([a]) => a === left);
    const end = row.findIndex(([a, b]) => b === right);
    return row.slice(start, end + 1);
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

function isPrime(x) {
    const size = Math.trunc(Math.pow(x, 1 / 2));
    for (let i = 2; i <= size; ++i) {
        if (x % i === 0) {
            return false;
        }
    }

    return true;
}

module.exports = sumOfPrimeDistance
