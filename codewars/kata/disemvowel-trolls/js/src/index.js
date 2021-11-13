function disemvowel(str) {
    const excludes = ["a", "e", "u", "o", "i"];
    return str.split("").filter(char => !excludes.includes(char.toLowerCase())).join("");
}

module.exports = disemvowel;
