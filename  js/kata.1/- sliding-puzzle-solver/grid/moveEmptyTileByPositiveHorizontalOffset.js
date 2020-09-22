const copy = require("./copy");

module.exports = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);
    for (let i = 0; i < length; ++i) {
        temp = moveLeft(temp)
    }

    return temp;
}
