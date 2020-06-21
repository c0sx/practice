function solve(arr) {
    return arr.reduce((acc, curr, index) => {
        const slice = arr.slice(index + 1)
        if (slice.includes(curr)) {
            return acc;
        }

        acc.push(curr)
        return acc;
    }, [])
}

module.exports = solve;
