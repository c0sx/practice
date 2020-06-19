const roundToNext5 = require("./index")

describe("Round up to the Next Multiple of 5", () => {
    it("tests", () => {
        [
            [0, 0],
            [1, 5],
            [3, 5],
            [5, 5],
            [7, 10],
            [39, 40],
            [-2, 0],
            [-5, -5]
        ].forEach(
            // toEqual not working
            ([x, out]) => expect(roundToNext5(x) === out).toBeTruthy()
        );

    });
})
