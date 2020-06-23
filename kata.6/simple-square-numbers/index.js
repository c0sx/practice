const solve = n => {
    return getAddition(n)
}

const getAddition = n => {
    let i = 1;

    while (true) {
        const k = n + i;

        const root = Math.sqrt(k);
        if (Math.trunc(root) === root) {
            return i;
        }

        i++;
    }
}

module.exports = solve;
