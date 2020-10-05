const decompose = require("./index");

describe("Square into Squares", () => {
    it("tests", () => {
        expect(decompose(85986)).toEqual([1, 3, 6, 23, 414, 85985]);
        expect(decompose(50)).toEqual([1, 3, 5, 8, 49]); //  1 + 9 + 25 + 64 + 2401 = 2500
        expect(decompose(11)).toEqual([1, 2, 4, 10]); // 1 + 4 + 16 + 100 = 121
        expect(decompose(7)).toEqual([2, 3, 6]);
        expect(decompose(4)).toEqual("Nothing");
        expect(decompose(2)).toEqual("Nothing");
    });
})
