const findCellCoordinates = (input, needle) => {
    const rowIndex = input.findIndex(row => row.includes(needle));
    const row = input[rowIndex];

    const cellIndex = row.findIndex(cell => cell === needle);
    return [rowIndex, cellIndex]
}

const findCoordinatesOfEmptyCell = input => findCellCoordinates(input, 0)

const findTile = (input, rowIndex, cellIndex) => input[rowIndex][cellIndex];

const rotate = input => {
    const rotated = [];
    const rowLength = input[0].length;
    for (let i = 0; i < rowLength; ++i) {
        const tiles = [];
        for (let j = 0; j < rowLength; ++j) {
            const tile = findTile(input, j, i)
            tiles.push(tile);
        }

        rotated.push(tiles);
    }

    return rotated;
}

const moveDown = (input) => {
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

const moveLeft = (input) => {
    const [spaceRowIndex, spaceCellIndex] = findCoordinatesOfEmptyCell(input);
    const row = input[spaceRowIndex];
    if (spaceCellIndex === row.length - 1) {
        return copy(input)
    }

    const tile = findTile(input, spaceRowIndex, spaceCellIndex + 1)
    const copied = copy(input);
    copied[spaceRowIndex][spaceCellIndex + 1] = 0;
    copied[spaceRowIndex][spaceCellIndex] = tile;

    return copied;
}

const moveRight = (input) => {
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

const moveUp = (input) => {
    const [spaceRowIndex, spaceCellIndex] = findCoordinatesOfEmptyCell(input);
    if (spaceRowIndex === input.length - 1) {
        return copy(input)
    }

    const tile = findTile(input, spaceRowIndex + 1, spaceCellIndex);
    const copied = copy(input);
    copied[spaceRowIndex + 1][spaceCellIndex] = 0;
    copied[spaceRowIndex][spaceCellIndex] = tile;

    return copied;
}

const moveEmptyTile = (input, [destRowIndex, destCellIndex]) => {
    const [spaceRowIndex, spaceCellIndex] = findCoordinatesOfEmptyCell(input);

    const vertical = destRowIndex - spaceRowIndex;
    const horizontal = destCellIndex - spaceCellIndex;

    const horizontalMoved = moveEmptyTileByHorizontalOffset(input, horizontal);
    return moveEmptyTileByVerticalOffset(horizontalMoved, vertical)
}

const moveEmptyTileByHorizontalOffset = (input, offset) => {
    if (offset === 0) {
        return copy(input)
    }

    if (offset < 0) {
        return moveEmptyTileByNegativeHorizontalOffset(input, offset)
    }

    return moveEmptyTileByPositiveHorizontalOffset(input, offset);
}

const moveEmptyTileByNegativeHorizontalOffset = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);
    for (let i = 0; i < length; ++i) {
        temp = moveRight(temp);
    }

    return temp;
}

const moveEmptyTileByPositiveHorizontalOffset = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);
    for (let i = 0; i < length; ++i) {
        temp = moveLeft(temp)
    }

    return temp;
}

const moveEmptyTileByVerticalOffset = (input, offset) => {
    if (offset === 0) {
        return copy(input)
    }

    if (offset < 0) {
        return moveEmptyTileByNegativeVerticalOffset(input, offset)
    }

    return moveEmptyTileByPositiveVerticalOffset(input, offset)
}

const moveEmptyTileByNegativeVerticalOffset = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);

    for (let i = 0; i < length; ++i) {
        temp = moveDown(temp);
    }

    return temp;
}

const moveEmptyTileByPositiveVerticalOffset = (input, offset) => {
    const length = Math.abs(offset);
    let temp = copy(input);

    for (let i = 0; i < length; ++i) {
        temp = moveUp(temp);
    }

    return temp;
}

const copy = (input) => {
    return input.slice().map(row => row.slice())
}

module.exports = {
    findCellCoordinates,
    findCoordinatesOfEmptyCell,
    rotate,
    moveDown,
    moveLeft,
    moveRight,
    moveUp,
    moveEmptyTile
}
