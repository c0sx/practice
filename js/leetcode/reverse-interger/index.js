const reverse = function(x) {
    const sign = Math.sign(x);
    const abs = Math.abs(x);
    const reversed = String(abs).split("").reverse().join("");
    const result = Number(reversed) * sign;

    if (result < Math.pow(-2, 31) || result > Math.pow(2, 31) - 1) {
        return 0;
    }

    return result;
};
