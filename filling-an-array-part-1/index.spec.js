const arr = require("./")

describe("#arr creates a new array with the numbers 0 to N-1", () => {
    it("should return an array", () => {
        expect(arr()).toBeInstanceOf(Array)
    });

    it("should return a blank array when called with no argument", () => {
        expect(arr()).toEqual([]);
    });

    it("should return 0 to 3 when called with 4", () => {
        expect(arr(4)).toEqual([0, 1, 2, 3])
    });
});

