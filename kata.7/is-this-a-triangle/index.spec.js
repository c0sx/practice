const isTriangle = require("./index")

describe("Is This a Triangle", () => {
    it("tests", () => {
        expect(isTriangle(1, 2, 2)).toEqual(true);
        expect(isTriangle(7, 2, 2)).toEqual(false);
    });
});
