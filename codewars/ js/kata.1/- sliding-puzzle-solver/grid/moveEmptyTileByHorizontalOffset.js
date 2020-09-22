const moveEmptyTileByNegativeHorizontalOffset = require("./moveEmptyTileByNegativeHorizontalOffset");
const moveEmptyTileByPositiveHorizontalOffset = require("./moveEmptyTileByPositiveHorizontalOffset");
const copy = require("./copy");

module.exports = (input, offset) => {
    if (offset === 0) {
        return copy(input)
    }

    if (offset < 0) {
        return moveEmptyTileByNegativeHorizontalOffset(input, offset)
    }

    return moveEmptyTileByPositiveHorizontalOffset(input, offset);
}
