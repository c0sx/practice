const moveEmptyTile = require("../moveEmptyTile")

describe("Sliding Puzzle Solver", () => {
    it("should move empty tile", () => {
        const input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 9, 8]
        ];

        const target = [0, 0];

        const output = [
            [9, 2, 3],
            [1, 5, 6],
            [4, 7, 8]
        ]

        expect(moveEmptyTile(input, target)).toEqual(output)
    });
})
