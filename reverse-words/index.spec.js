const reverseWords = require("./");

describe("reverse-words", () => {
    it("should reverse-words in sentences", () => {
        expect(reverseWords("The quick brown fox jumps over the lazy dog."))
            .toEqual("ehT kciuq nworb xof spmuj revo eht yzal .god")
        expect(reverseWords("apple")).toEqual("elppa")
        expect(reverseWords("a b c d")).toEqual("a b c d")
        expect(reverseWords("double  spaced  words")).toEqual("elbuod  decaps  sdrow")
    });
});

