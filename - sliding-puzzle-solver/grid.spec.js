const grid = require("./grid")

describe("Sliding Puzzle Solver", () => {
    let input;

    beforeEach(() => {
        input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 0],
        ]
    })

    it("grid should rotate correctly", () => {
        const output = [
            [1, 4, 7],
            [2, 5, 8],
            [3, 6, 0],
        ]

        expect(grid.rotate(input)).toEqual(output)
        expect(grid.rotate(output)).toEqual(input)
    });

    it("should move a tile down", () => {
        const input2 = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 0],
        ]

        const output2 = [
            [1, 2, 3],
            [4, 5, 0],
            [7, 8, 6]
        ]

        expect(grid.moveDown(input2)).toEqual(output2)
    });

    it("shouldn't change input if cant move a tile down", () => {
        const input1 = [
            [1, 0, 3],
            [2, 4, 5],
            [6, 7, 8]
        ];

        const output1 = [
            [1, 0, 3],
            [2, 4, 5],
            [6, 7, 8]
        ]

        expect(grid.moveDown(input1)).toEqual(output1)
    });

    it("should move a tile to the left", () => {
        const input = [
            [1, 0, 3],
        ]

        const output = [
            [1, 3, 0]
        ];

        expect(grid.moveLeft(input)).toEqual(output)
    });

    it("shouldn't move move tile if cant", () => {
        const input = [
            [1, 3, 0]
        ];

        const output = [
            [1, 3, 0]
        ]

        expect(grid.moveLeft(input)).toEqual(output)
    });

    it("should move a tile to the right", () => {
        const input = [
            [1, 0, 3]
        ]

        const output = [
            [0, 1, 3]
        ]

        expect(grid.moveRight(input)).toEqual(output)
    });

    it("shouldn't move if cant", () => {
        const input = [
            [0, 1, 3]
        ];

        const output = [
            [0, 1, 3]
        ]

        expect(grid.moveRight(input)).toEqual(output)
    });

    it("should move a tile to the up", () => {
        const input = [
            [1, 2, 3],
            [4, 0, 5],
            [6, 7, 8]
        ]

        const output = [
            [1, 2, 3],
            [4, 7, 5],
            [6, 0, 8]
        ]

        expect(grid.moveUp(input)).toEqual(output)
    });

    it("shouldn't move if cant", () => {
        const input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 0, 8]
        ];

        expect(grid.moveUp(input)).toEqual(input)
    });
})
