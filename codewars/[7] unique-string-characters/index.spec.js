const solve = require("./index");

describe("Unique string character", () => {
    it("tests", () => {
        expect(solve("xyab", "xzca")).toEqual("ybzc");
        expect(solve("xyabb", "xzca")).toEqual("ybbzc");
        expect(solve("abcd", "xyz")).toEqual("abcdxyz");
        expect(solve("xxx", "xzca")).toEqual("zca");
    });
})
