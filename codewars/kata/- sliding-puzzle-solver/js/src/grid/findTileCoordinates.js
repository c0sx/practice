module.exports = (input, needle) => {
    const rowIndex = input.findIndex(row => row.includes(needle));
    const row = input[rowIndex] || [];

    const cellIndex = row.findIndex(cell => cell === needle);
    return [rowIndex, cellIndex]
}
