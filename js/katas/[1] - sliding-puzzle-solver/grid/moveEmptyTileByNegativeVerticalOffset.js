const copy = require("./copy");
const moveDown = require("./moveDown");

module.exports = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);

    for (let i = 0; i < length; ++i) {
        temp = moveDown(temp);
    }

    return temp;
}
