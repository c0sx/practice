const grid = require("../grid");

module.exports = input => {
    const [expectedRowIndex, expectedCellIndex] = grid.findCoordinatesOfActualTile(input);
    const expectedTile = grid.findTile(input, expectedRowIndex, expectedCellIndex);

    const rowIndex = grid.findActualRow(input);
    const row = input[rowIndex];
    const factor = rowIndex * row.length + 1;
    const expected = row.map((cell, cellIndex) => cellIndex + factor);
    const lastOfRow = expected.pop();
    if (expectedTile !== lastOfRow) {
        // разместить ее в нужном месте не трогая зафиксированные элементы
        grid.moveTile(input, expectedTile, [expectedRowIndex, expectedCellIndex]);
    }


    // поместить ее на позицию под ее правильным местом и пустой тайл под нее
    // хак, игнорирует фиксацию элементов

    // собрать строку -- повернуть пазл
}
