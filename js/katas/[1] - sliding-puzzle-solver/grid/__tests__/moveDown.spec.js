const moveDown = require("../moveDown")

describe("Sliding Puzzle Solver", () => {

    it("should move a tile down", () => {
        const input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
        ]

        const output = [
            [1, 2, 3],
            [4, 5, 9],
            [7, 8, 6]
        ]

        expect(moveDown(input)).toEqual(output)
    });

    it("shouldn't change input if cant move a tile down", () => {
        const input = [
            [1, 9, 3],
            [2, 4, 5],
            [6, 7, 8]
        ];

        const output = [
            [1, 9, 3],
            [2, 4, 5],
            [6, 7, 8]
        ]

        expect(moveDown(input)).toEqual(output)
    });
})
