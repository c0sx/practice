const findCoordinatesOfEmptyCell = require("./findCoordinatesOfEmptyCell");
const findTile = require("./findTile");
const copy = require("./copy")

module.exports = input => {
    const [spaceRowIndex, spaceCellIndex] = findCoordinatesOfEmptyCell(input);
    if (spaceCellIndex === 0) {
        return copy(input)
    }

    const tile = findTile(input, spaceRowIndex, spaceCellIndex - 1);
    const copied = copy(input);
    copied[spaceRowIndex][spaceCellIndex - 1] = 0;
    copied[spaceRowIndex][spaceCellIndex] = tile;

    return copied;
}
