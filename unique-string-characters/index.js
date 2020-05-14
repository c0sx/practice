function solve(a, b) {
    const as = a.split("");
    const bs = b.split("");

    const first = as.filter(one => !bs.includes(one))
    const second = bs.filter(one => !as.includes(one))

    return [...first, ...second].join("")
}

console.log(solve("xyab","xzca")) // "ybzc"
console.log(solve("xyabb","xzca")) // "ybbzc"
console.log(solve("abcd","xyz")) // "abcdxyz"
console.log(solve("xxx","xzca")) // "zca"
