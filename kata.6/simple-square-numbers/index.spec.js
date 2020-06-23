const solve = require("./index");

describe("Simple square numbers", () => {

    it("tests", () => {
        // expect(solve(1)).toEqual(-1)
        // expect(solve(2)).toEqual(-1)
        // expect(solve(3)).toEqual(1)
        // expect(solve(4)).toEqual(-1)
        // expect(solve(5)).toEqual(4)
        expect(solve(7)).toEqual(9)
        // expect(solve(8)).toEqual(1)
        // expect(solve(9)).toEqual(16)
        // expect(solve(10)).toEqual(-1)
        // expect(solve(11)).toEqual(25)
        // expect(solve(13)).toEqual(36)
        // expect(solve(17)).toEqual(64)
        // expect(solve(88901)).toEqual(5428900)
        // expect(solve(290101)).toEqual(429235524)
    });
});
