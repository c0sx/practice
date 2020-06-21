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

module.exports = splitAndAdd;
