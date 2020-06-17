const findActualRow = require("../findActualRow");

describe("Sliding Puzzle Solver", () => {
    it("should return actual row", () => {
        expect(findActualRow([
            [1, 7, 2],
            [3, 4, 5],
            [6, 8, 9]
        ])).toEqual(0)

        expect(findActualRow([
            [1, 2, 3],
            [4, 7, 6],
            [5, 8, 9]
        ])).toEqual(1)

        expect(findActualRow([
            [1, 2, 3],
            [4, 5, 6],
            [8, 7, 9]
        ])).toEqual(2)

        expect(findActualRow([
            [9, 1, 2],
            [3, 4, 5],
            [6, 7, 8]
        ])).toEqual(0)

        expect(findActualRow([
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
        ])).toEqual(-1)
    });
})
