function accum(s) {
    return s.split("").map((char, index) => {
        const string = char.repeat(index)
        return char.toUpperCase() + string.toLowerCase();
    }).join("-");
}

module.exports = accum;
