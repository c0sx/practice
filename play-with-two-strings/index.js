function workOnStrings(a, b) {
    let asOriginal = a.split("");
    let bsOriginal = b.split("");

    const bs = transform(asOriginal, bsOriginal);
    const as = transform(bs, asOriginal);

    return as.concat(bs).join("")
}

function transform(a, b) {
    a.forEach(letter => {
        b = b.map(one => compare(one, letter))
    })

    return b;
}

function compare(one, letter) {
    if (one.toLowerCase() === letter.toLowerCase()) {
        if (one.toUpperCase() === one) {
            return one.toLowerCase()
        }
        else {
            return one.toUpperCase()
        }
    }

    return one;
}

module.exports = workOnStrings
