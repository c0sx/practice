const { slidePuzzle, isFinished } = require("./index")

const validate = (input, sequence) => {
    sequence.forEach(tile => {
        const [spaceRow, spaceCell] = findTile(input, 0);
        const [tileRow, tileCell] = findTile(input, tile);

        input[spaceRow][spaceCell] = tile;
        input[tileRow][tileCell] = 0;
    })

    return isFinished(input)
}

const findTile = (input, tile) => {
    const rowIndex = input.findIndex(row => row.includes(tile));
    const row = input[rowIndex];
    const cellIndex = row.findIndex(cell => cell === tile);
    return [rowIndex, cellIndex];
}

describe("Sliding Puzzle Solver", () => {
    it("tests", () => {
        // let puzzle0 = [
        //     [0]
        // ];
        //
        const puzzle1 = [
            [4, 1, 3],
            [2, 8, 0],
            [7, 6, 5]
        ];

        const puzzle2 = [
            [10, 3, 6, 4],
            [ 1, 5, 8, 0],
            [ 2, 13, 7, 15],
            [14, 9, 12, 11]
        ];

        const puzzle3 = [
            [ 3, 7, 14, 15, 10],
            [ 1, 0, 5, 9, 4],
            [16, 2, 11, 12, 8],
            [17, 6, 13, 18, 20],
            [21, 22, 23, 19, 24]
        ];

        const simpleExample = [
            [ 1, 2, 3, 4],
            [ 5, 0, 6, 8],
            [ 9, 10, 7, 11],
            [13, 14, 15, 12]
        ];

        const simpleInput = simpleExample.map(r => r.slice());
        const simpleSequence = slidePuzzle(simpleInput)
        expect(simpleSequence).toEqual([6, 7, 11, 12]);
        expect(validate(simpleExample, simpleSequence)).toBeTruthy();

        const input1 = puzzle1.map(r => r.slice())
        const sequence1 = slidePuzzle(input1)
        expect(sequence1).toEqual([8, 2, 4, 1, 2, 8, 5, 6, 8, 5, 6]);
        expect(validate(puzzle1, sequence1)).toBeTruthy();

        const input2 = puzzle2.map(r => r.slice());
        const sequence2 = slidePuzzle(input2);
        expect(validate(puzzle2, sequence2)).toBeTruthy()

        const input3 = puzzle3.map(r => r.slice());
        const sequence3 = slidePuzzle(input3)
        expect(validate(puzzle3, sequence3)).toBeTruthy()
    });
})

