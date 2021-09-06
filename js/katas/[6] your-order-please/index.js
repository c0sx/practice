const regexp = /(\d)/;

const order = words => {
    if (words.length === 0) {
        return "";
    }

    return words.split(" ").sort((a, b) => {
        const a1 = a.match(regexp)[0]
        const b1 = b.match(regexp)[0]

        return a1 - b1;
    }).join(" ")
}

module.exports = {
    order
};
