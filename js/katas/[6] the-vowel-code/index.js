const map = new Map([
    ["a", 1],
    ["e", 2],
    ["i", 3],
    ["o", 4],
    ["u", 5]
]);

const reverse = new Map([
    ["1", "a"],
    ["2", "e"],
    ["3", "i"],
    ["4", "o"],
    ["5", "u"]
])

function encode(string) {
    return string.split("").map(one => map.get(one) || one).join("");
}

function decode(string) {
    return string.split("").map(one => reverse.get(one) || one).join("");
}

module.exports = {
    encode,
    decode
}
