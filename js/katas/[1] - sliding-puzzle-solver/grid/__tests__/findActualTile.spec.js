const findActualTile = require("../findActualTile");

describe("Sliding Puzzle Solver", () => {
    it("should return next tile", () => {
        expect(findActualTile([
            [1, 4, 2],
            [3, 5, 6],
            [7, 8, 9]
        ])).toEqual(2)

        expect(findActualTile([
            [1, 2, 3],
            [9, 4, 5],
            [6, 7, 8]
        ])).toEqual(4)

        expect(findActualTile([
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9]
        ])).toEqual(9)
    });
})
