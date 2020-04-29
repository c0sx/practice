function duplicateEncode(word){
    return word.split("").reduce((accumulator, char, index, origin) => {
        const newChar = origin.filter(symbol => char.toLowerCase() === symbol.toLowerCase()).length > 1 ? ")" : "(";
        return accumulator + newChar;
    }, "");
}

console.log(duplicateEncode("din")) // "((("
console.log(duplicateEncode("recede")) // "()()()"
console.log(duplicateEncode("Success")) // ")())())"
console.log(duplicateEncode("(( @")) // "))(("
