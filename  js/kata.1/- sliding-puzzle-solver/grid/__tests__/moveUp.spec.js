const moveUp = require("../moveUp")

describe("Sliding Puzzle Solver", () => {

    it("should move a tile to the up", () => {
        const input = [
            [1, 2, 3],
            [4, 9, 5],
            [6, 7, 8]
        ]

        const output = [
            [1, 2, 3],
            [4, 7, 5],
            [6, 9, 8]
        ]

        expect(moveUp(input)).toEqual(output)
    });

    it("shouldn't move if cant", () => {
        const input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 9, 8]
        ];

        expect(moveUp(input)).toEqual(input)
    });
});
