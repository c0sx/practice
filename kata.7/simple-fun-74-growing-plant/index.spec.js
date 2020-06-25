const growingPlant = require("./index");

describe("Simple fund #74: Growing plant", function() {
    it("tests", function() {
        expect(growingPlant(100, 10, 910)).toEqual(10)
        expect(growingPlant(10, 9, 4)).toEqual(1)
    });
});
