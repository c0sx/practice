const findCoordinatesOfEmptyCell = require("./findCoordinatesOfEmptyCell");
const findTile = require("./findTile");
const copy = require("./copy")

module.exports = input => {
    const [spaceRowIndex, spaceCellIndex] = findCoordinatesOfEmptyCell(input);
    if (spaceRowIndex === 0) {
        return copy(input)
    }

    const tile = findTile(input, spaceRowIndex - 1, spaceCellIndex);
    const copied = copy(input)
    copied[spaceRowIndex - 1][spaceCellIndex] = 0
    copied[spaceRowIndex][spaceCellIndex] = tile;

    return copied
}
