const ulamSequence = require("./")

describe("Ulam Sequence (performance edition)", () => {
    it("tests", () => {
        // expect(ulamSequence(1, 2, 5)).toEqual([1, 2, 3, 4, 6]);
        // expect(ulamSequence(3, 4, 5)).toEqual([3, 4, 7, 10, 11]);
        // expect(ulamSequence(5, 6, 8)).toEqual([5, 6, 11, 16, 17, 21, 23, 26]);
        // expect(ulamSequence(3, 4, 5)).toEqual([3, 4, 7, 10, 11]);
        expect(ulamSequence(1, 2, 20)).toEqual([1, 2, 3, 4, 6, 8, 11, 13, 16, 18, 26, 28, 36, 38, 47, 48, 53, 57, 62, 69]);
        // expect(ulamSequence(1, 3, 60)).toEqual([1, 3, 4, 5, 6, 8, 10, 12, 17, 21, 23, 28, 32, 34, 39, 43, 48, 52, 54, 59, 63, 68, 72, 74, 79, 83, 98, 99, 101, 110, 114, 121, 125, 132, 136, 139, 143, 145, 152, 161, 165, 172, 176, 187, 192, 196, 201, 205, 212, 216, 223, 227, 232, 234, 236, 243, 247, 252, 256, 258]);
    });
});

// [1, 2, 3, 4, 6]
// left = [1, 2, 3]
// right = [4, 5]

// 1 + 3 = 4
// 2 + 3 = 5

// 1 + 3 = 4
// 2 + 3 = 5
// 1 + 4 = 5
// 2 + 4 = 6

// [3, 4, 7, 10, 11]
// left queue = [3, 4]
// right queue = [7, 10, 11]

// 3 + 7 = 10
// 4 + 7 = 11
