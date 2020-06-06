const logic = require("./logic");

describe("Sliding Puzzle Solver", () => {
    it("logic should find next row", () => {
        const input = [
            [1, 2, 3],
            [4, 0, 5],
            [6, 7, 8]
        ];

        const output = [4, 5, 6]
        expect(logic(input).findNextRow()).toEqual(output)
    });
})
