const validatePIN = require("./index")

describe("Regex validate PIN code", () => {
    it("tests", () => {
        expect(validatePIN("1234")).toBeTruthy(); // true
        expect(validatePIN("12345")).toBeFalsy(); // false
        expect(validatePIN("a234")).toBeFalsy(); // false
    });
})
