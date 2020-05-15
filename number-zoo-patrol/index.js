function findNumber(arr) {
    const totalLength = arr.length + 1
    const totalSum = (totalLength * (totalLength + 1)) / 2
    const currentSum = arr.reduce((acc, num) => acc + num, 0)
    return totalSum - currentSum
}

module.exports = findNumber
