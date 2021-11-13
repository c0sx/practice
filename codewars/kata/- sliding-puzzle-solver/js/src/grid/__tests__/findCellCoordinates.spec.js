const findCellCoordinates = require("../findTileCoordinates")

describe("Sliding Puzzle Solver", () => {
    it("should find cell coordinates by value", () => {
        const input = [
            [1, 2, 3],
            [4, 5, 6],
        ];

        expect(findCellCoordinates(input, 5)).toEqual([1, 1]);
        expect(findCellCoordinates(input, 1)).toEqual([0, 0]);
        expect(findCellCoordinates(input, 6)).toEqual([1, 2]);
        expect(findCellCoordinates(input, 10)).toEqual([-1, -1]);
    });
})
