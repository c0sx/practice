const findCellCoordinates = (input, needle) => {
    const rowIndex = input.findIndex(row => row.includes(needle));
    const row = input[rowIndex];

    const cellIndex = row.findIndex(cell => cell === needle);
    return [rowIndex, cellIndex]
}

const findCoordinatesOfEmptyCell = input => findCellCoordinates(input, 0)

const rotate = input => {
    const rotated = [];
    const rowLength = input[0].length;
    for (let i = 0; i < rowLength; ++i) {
        const cells = [];
        for (let j = 0; j < rowLength; ++j) {
            const cell = input[j][i];
            cells.push(cell);
        }

        rotated.push(cells);
    }

    return rotated;
}

module.exports = {
    findCellCoordinates,
    findCoordinatesOfEmptyCell,
    rotate,
}
