const moveEmptyTileByNegativeVerticalOffset = require("./moveEmptyTileByNegativeVerticalOffset");
const moveEmptyTileByPositiveVerticalOffset = require("./moveEmptyTileByPositiveVerticalOffset");
const copy = require("./copy");

module.exports = (input, offset) => {
    if (offset === 0) {
        return copy(input)
    }

    if (offset < 0) {
        return moveEmptyTileByNegativeVerticalOffset(input, offset)
    }

    return moveEmptyTileByPositiveVerticalOffset(input, offset)
}
