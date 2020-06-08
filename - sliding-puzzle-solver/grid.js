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
//
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
    moveUp
}
