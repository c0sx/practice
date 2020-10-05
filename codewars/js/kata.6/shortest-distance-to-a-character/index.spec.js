const shortesttoChar = require("./index");

describe("Shortest distance to a character", () => {
    it("tests", () => {
        expect(shortesttoChar("lovecodewars", "e")).toEqual([3, 2, 1, 0, 1, 2, 1, 0, 1, 2, 3, 4])
        expect(shortesttoChar("aaaabbbb", "b")).toEqual([4, 3, 2, 1, 0, 0, 0, 0])
        expect(shortesttoChar("", "b")).toEqual([])
        expect(shortesttoChar("abcde", "")).toEqual([])
    });
})
