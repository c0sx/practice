const moveLeft = require("../moveLeft");

describe("Sliding Puzzle Solver", () => {
    it("should move a tile to the left", () => {
        const input = [
            [1, 9, 3],
        ]

        const output = [
            [1, 3, 9]
        ];

        expect(moveLeft(input)).toEqual(output)
    });

    it("shouldn't move move tile if cant", () => {
        const input = [
            [1, 3, 9]
        ];

        const output = [
            [1, 3, 9]
        ]

        expect(moveLeft(input)).toEqual(output)
    });
})
