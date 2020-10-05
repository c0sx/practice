const toCamelCase = require("./index");

describe("Convert string to CamelCase", () => {
    it("tests", () => {
        expect(toCamelCase("")).toEqual("") // , "An empty string was provided but not returned")
        expect(toCamelCase("the_stealth_warrior")).toEqual("theStealthWarrior") // , "toCamelCase('the_stealth_warrior') did not return correct value")
        expect(toCamelCase("The-Stealth-Warrior")).toEqual("TheStealthWarrior") // , "toCamelCase('The-Stealth-Warrior') did not return correct value")
        expect(toCamelCase("A-B-C")).toEqual("ABC") // "toCamelCase('A-B-C') did not return correct value")
    });
})
