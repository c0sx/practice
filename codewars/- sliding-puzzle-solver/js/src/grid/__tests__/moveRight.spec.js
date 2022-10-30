const moveRight = require("../moveRight");

describe("Sliding Puzzle Solver", () => {

    it("should move a tile to the right", () => {
        const input = [
            [1, 9, 3]
        ]

        const output = [
            [9, 1, 3]
        ]

        expect(moveRight(input)).toEqual(output)
    });

    it("shouldn't move if cant", () => {
        const input = [
            [9, 1, 3]
        ];

        const output = [
            [9, 1, 3]
        ]

        expect(moveRight(input)).toEqual(output)
    });
})
