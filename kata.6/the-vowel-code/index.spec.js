const { encode, decode } = require("./index");

describe("The Vowel Code", () => {
    it("tests", () => {
        expect(encode("hello")).toEqual("h2ll4");
        expect(encode("How are you today?")).toEqual("H4w 1r2 y45 t4d1y?");
        expect(encode("This is an encoding test.")).toEqual("Th3s 3s 1n 2nc4d3ng t2st.");
        expect(decode("h2ll4")).toEqual("hello");
    });
})
