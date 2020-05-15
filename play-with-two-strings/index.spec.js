const workOnStrings = require("./");

describe("Play with Two Strings", function() {
    it("tests", function() {
        expect(workOnStrings("abc", "cde")).toEqual("abCCde");
        expect(workOnStrings("abcdeFgtrzw", "defgGgfhjkwqe")).toEqual("abcDeFGtrzWDEFGgGFhjkWqE");
        expect(workOnStrings("abcdeFg", "defgG")).toEqual("abcDEfgDEFGg");
        expect(workOnStrings("abab", "bababa")).toEqual("ABABbababa");
    });
});
