const multiply = require("./index");

describe("Multiply", function() {
    it("tests", function() {
        expect(multiply(5, 4)).toEqual(20);
        expect(multiply(10, 3)).toEqual(30)
    });
});
