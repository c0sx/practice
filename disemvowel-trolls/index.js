function disemvowel(str) {
    const excludes = ["a", "e", "u", "o", "i"];
    return str.split("").filter(char => !excludes.includes(char.toLowerCase())).join("");
}

console.log(disemvowel("This website is for losers LOL!")); // Ths wbst s fr lsrs LL!"
