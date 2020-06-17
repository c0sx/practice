const copy = require("./copy");
const moveRight = require("./moveRight");

module.exports = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);
    for (let i = 0; i < length; ++i) {
        temp = moveRight(temp);
    }

    return temp;
}
