function splitAndAdd(arr, n) {
    if (n === 0) {
        return arr
    }

    const [left, right] = split(arr);
    const sum = add(left, right)
    return splitAndAdd(sum, n - 1);
}

function split(arr) {
    const index = Math.floor(arr.length / 2);
    const left = arr.slice(0, index);
    const right = arr.slice(index);
    return [left, right]
}

function add(left, right) {
    const [max, other] = getMaxAndOtherArrays(left, right);

    const diff = max.length - other.length;
    const padded = Array.from({ length: diff }, () => 0).concat(other);
    return max.map((one, index) => {
        const value = padded[index] || 0;
        return one + value;
    })
}

function getMaxAndOtherArrays(left, right) {
    if (left.length > right.length) {
        return [left, right]
    }

    return [right, left]
}

console.log(splitAndAdd([1,2,3,4,5], 2)) //, [5,10])
console.log(splitAndAdd([1,2,3,4,5], 3)) //, [15])
console.log(splitAndAdd([15], 3)) //, [15])
console.log(splitAndAdd([32,45,43,23,54,23,54,34], 2)) //, [183, 125])
console.log(splitAndAdd([32,45,43,23,54,23,54,34], 0)) //, [32,45,43,23,54,23,54,34])
console.log(splitAndAdd([3,234,25,345,45,34,234,235,345], 3)) //, [305, 1195])
console.log(splitAndAdd([3,234,25,345,45,34,234,235,345,34,534,45,645,645,645,4656,45,3], 4)) //, [1040, 7712])
console.log(splitAndAdd([23,345,345,345,34536,567,568,6,34536,54,7546,456], 20)) //, [79327])
