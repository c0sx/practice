const elevator = require("./index");

describe("Closest elevator", () => {
    it("tests", () => {
        expect(elevator(0, 1, 0)).toEqual("left"); // => "left"
        expect(elevator(0, 1, 1)).toEqual("right"); // => "right"
        expect(elevator(0, 1, 2)).toEqual("right") // => "right"
        expect(elevator(0, 0, 0)).toEqual("right") // => "right"
        expect(elevator(0, 2, 1)).toEqual("right") // => "right"
    });
})
