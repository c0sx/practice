const rotate = require("../rotate")

describe("Sliding Puzzle Solver", () => {
    it("grid should rotate correctly", () => {
        const input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
        ];

        const output = [
            [1, 4, 7],
            [2, 5, 8],
            [3, 6, 9],
        ];

        expect(rotate(input)).toEqual(output)
        expect(rotate(output)).toEqual(input)
    });
})
