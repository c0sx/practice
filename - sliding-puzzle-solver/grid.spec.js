const grid = require("./grid")

describe("Sliding Puzzle Solver", () => {
    it("grid should rotate correctly", () => {
        const input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 0],
        ]

        const output = [
            [1, 4, 7],
            [2, 5, 8],
            [3, 6, 0],
        ]

        expect(grid.rotate(input)).toEqual(output)
        expect(grid.rotate(output)).toEqual(input)
    });
})
