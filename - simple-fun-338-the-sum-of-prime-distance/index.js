function sumOfPrimeDistance(arr) {
    console.time("pairs")
    const pairs = getPairs(arr);
    console.timeEnd("pairs");

    console.time("primed");
    let primed = calculatePrimes(pairs);
    console.timeEnd("primed")

    console.time("counter")
    let step = 0;
    const row = [];

    // if arr.length % 2 === 0 -- middle === primed.length / 2
    // if arr.length % 2 === 1 -- middle === primed.length / 2, но middle 2 значения
    while (step < primed.length / 2) { // каждый ряд прямоугольника
        let rowSum = 0;
        for (let i = 0; i < primed.length - step; ++i) { // скашиваем прямоугольник до треугольника
            for (let k = i; k < i + step + 1; ++k) { // считаем каждый отрезок в строке треугольника
                rowSum += primed[k]
            }
        }

        row.push(rowSum)
        step++;
    }

    console.timeEnd("counter")

    console.time("result")
    if (arr.length % 2 === 0) {
        const last = row.pop();
        const total = last + row.reduce((acc, curr) => acc + curr, 0) * 2
        console.timeEnd("result");
        return total;
    }

    const last = row.pop();
    row.pop();

    const total = last * 2 + row.reduce((acc, curr) => acc + curr, 0) * 2
    console.timeEnd("result");
    return total;
}

const sumRow = row => {
    let sum = 0;
    for (let i = 0; i < row.length; ++i) {
        sum += row[i]
    }

    return sum;
}

const getPairs = arr => {
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
        return primes.length;
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
