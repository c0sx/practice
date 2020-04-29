function accum(s) {
    return s.split("").map((char, index) => {
        const string = char.repeat(index)
        return char.toUpperCase() + string.toLowerCase();
    }).join("-");
}

console.log(accum("abcd")) // "A-Bb-Ccc-Dddd"
console.log(accum("RqaEzty")) // "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
console.log(accum("cwAt")) // "C-Ww-Aaa-Tttt"
