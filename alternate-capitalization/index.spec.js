const capitalize = require("./index");

describe("Alternate Capitalization", function() {
    it("tests", () => {
        expect(capitalize("abcdef")).toEqual(["AbCdEf", "aBcDeF"]);
        expect(capitalize("codewars")).toEqual(["CoDeWaRs", "cOdEwArS"]);
        expect(capitalize("abracadabra")).toEqual(["AbRaCaDaBrA", "aBrAcAdAbRa"]);
        expect(capitalize("codewarriors")).toEqual(["CoDeWaRrIoRs", "cOdEwArRiOrS"]);
    });
});
