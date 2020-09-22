function reverseWords(str) {
    return str.split(" ")
        .map(one => one.split("").reverse().join(""))
        .join(" ")
}

module.exports = reverseWords



