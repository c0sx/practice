function duplicateEncode(word) {
    return word.split("").reduce((accumulator, char, index, origin) => {
        const newChar = origin.filter(symbol => char.toLowerCase() === symbol.toLowerCase()).length > 1 ? ")" : "(";
        return accumulator + newChar;
    }, "");
}

module.exports = duplicateEncode;

