const josephus = require("./")

describe("Josephus Permutation", () => {
    it("tests", () => {
        expect(josephus([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 1)).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])
        expect(josephus([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 2)).toEqual([2, 4, 6, 8, 10, 3, 7, 1, 9, 5])
        expect(josephus(["C", "o", "d", "e", "W", "a", "r", "s"], 4)).toEqual(["e", "s", "W", "o", "C", "d", "r", "a"])
        expect(josephus([1, 2, 3, 4, 5, 6, 7], 3)).toEqual([3, 6, 2, 7, 5, 1, 4])
        expect(josephus([], 3)).toEqual([])
    });
})
