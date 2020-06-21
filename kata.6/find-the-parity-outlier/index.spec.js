const findOutlier = require("./index");

describe("Find the Parity Outlier", () => {
    it("tests", () => {
        expect(findOutlier([2, 4, 0, 100, 4, 11, 2602, 36])).toEqual(11); // Should return: 11 (the only odd number)
        expect(findOutlier([160, 3, 1719, 19, 11, 13, -21])).toEqual(160) // Should return: 160 (the only even number)
        expect(findOutlier([0, 1, 2])).toEqual(1)
        expect(findOutlier([1, 2, 3])).toEqual(2)
        expect(findOutlier([2, 6, 8, 10, 3])).toEqual(3)
        expect(findOutlier([0, 0, 3, 0, 0])).toEqual(3)
        expect(findOutlier([1, 1, 0, 1, 1])).toEqual(0)
    });
})
