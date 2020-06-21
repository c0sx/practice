function solve(a, b) {
    const as = a.split("");
    const bs = b.split("");

    const first = as.filter(one => !bs.includes(one))
    const second = bs.filter(one => !as.includes(one))

    return [...first, ...second].join("")
}

module.exports = solve;
