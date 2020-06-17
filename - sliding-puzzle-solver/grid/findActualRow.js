module.exports = input => {
    return input.findIndex((row, rowIndex) => {
        const factor = row.length * rowIndex + 1;

        return row.some((cell, cellIndex) => {
            return cell !== cellIndex + factor;
        })
    })
}
