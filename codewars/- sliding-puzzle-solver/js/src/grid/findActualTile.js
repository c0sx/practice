const findActualRow = require("./findActualRow");

module.exports = input => {
    const rowIndex = findActualRow(input);
    const row = input[rowIndex];

    const factor = row.length * rowIndex + 1;
    const index = row.findIndex((cell, cellIndex) => cell !== cellIndex + factor);
    return index + factor;
}
