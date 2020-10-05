function chain(fns) {
    return chained(0, fns);
}

const chained = (value, fns) => {
    return Object.keys(fns).reduce((acc, key) => {
        const f = fns[key];
        acc[key] = (...args) => {
            if (value === 0) {
                return chained(f(...args), fns)
            }

            return chained(f(value, ...args), fns)
        };

        return acc;
    }, {
        execute:() => {
            return value;
        }
    });
}

module.exports = chain;
