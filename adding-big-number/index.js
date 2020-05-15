function add(a, b) {
    const size = Math.max(a.length, b.length);
    const reversedA = a.split("").reverse().map(Number);
    const reversedB = b.split("").reverse().map(Number);

    const sum = [];
    let remainder = 0;
    for (let i = 0; i < size; ++i) {
        const x = reversedA[i] || 0
        const y = reversedB[i] || 0;
        const z = x + y + remainder;

        if (z < 10) {
            sum.push(z);
            remainder = 0;
            continue
        }

        const [whole, part = 0] = splitNumber(z);
        remainder = whole
        sum.push(part)
    }

    const total = sum.reverse().join("");
    if (remainder > 0) {
        return `${remainder}${total}`;
    }

    return total;
}

function splitNumber(n) {
    return String(n / 10).split(".").map(Number);
}

module.exports = add;
