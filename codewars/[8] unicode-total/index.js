function uniTotal(str) {
    let total = 0;
    for (let i = 0; i < str.length; ++i) {
        total += str.charCodeAt(i)
    }

    return total;
}

module.exports = uniTotal
