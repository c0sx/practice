const findCoordinatesOfEmptyCell = require("./findCoordinatesOfEmptyCell");
const moveEmptyTileByHorizontalOffset = require("./moveEmptyTileByHorizontalOffset");
const moveEmptyTileByVerticalOffset = require("./moveEmptyTileByVerticalOffset");

module.exports = (input, [destRowIndex, destCellIndex]) => {
    const [spaceRowIndex, spaceCellIndex] = findCoordinatesOfEmptyCell(input);

    const vertical = destRowIndex - spaceRowIndex;
    const horizontal = destCellIndex - spaceCellIndex;

    const horizontalMoved = moveEmptyTileByHorizontalOffset(input, horizontal);
    return moveEmptyTileByVerticalOffset(horizontalMoved, vertical)
}
