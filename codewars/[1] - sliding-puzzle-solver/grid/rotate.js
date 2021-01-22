const findTile = require("./findTile")

module.exports = input => {
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
