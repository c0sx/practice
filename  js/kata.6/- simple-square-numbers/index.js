const solve = n => {
    if (n === 1 || n % 2 === 0) {
        return -1;
    }

    return getAddition(n);
}

const getAddition = n => {
    let i = 1;

    while (true) {
        const k = n + i * i;

        const root = Math.sqrt(k);
        if (Math.trunc(root) === root) {
            return i * i
        }

        i++;
    }
}

module.exports = solve;
