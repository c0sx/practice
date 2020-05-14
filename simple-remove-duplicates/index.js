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

console.log(solve([3,4,4,3,6,3])) // [4,6,3]
console.log(solve([1,2,1,2,1,2,3])) // [1,2,3]
console.log(solve([1,2,3,4])) // [1,2,3,4]
console.log(solve([1,1,4,5,1,2,1])) // [4,5,2,1]
console.log(solve([1,2,1,2,1,1,3])) // [2,1,3]
