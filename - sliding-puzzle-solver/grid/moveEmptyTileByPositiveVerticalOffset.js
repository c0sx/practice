const copy = require("./copy");
const moveUp = require("./moveUp");

module.exports = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);

    for (let i = 0; i < length; ++i) {
        temp = moveUp(temp);
    }

    return temp;
}
