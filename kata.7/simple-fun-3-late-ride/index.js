const lateRide = n => {
    const hours = Math.trunc(n / 60);
    const minutes = n - hours * 60;

    return [hours, minutes].map(sum).reduce((s, a) => s + a);
};

const sum = a => {
    let s = 0;

    while (a > 0) {
        s += a % 10;
        a = Math.trunc(a / 10);
    }

    return s;
}

module.exports = lateRide;
