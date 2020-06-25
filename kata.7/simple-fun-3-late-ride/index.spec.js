const lateRide = require("./index");

describe("Simple fun #3: Late Ride", function() {
    it("tests", function() {
        expect(lateRide(240)).toEqual(4)
        expect(lateRide(808)).toEqual(14)
        expect(lateRide(1439)).toEqual(19)
        expect(lateRide(0)).toEqual(0)
        expect(lateRide(23)).toEqual(5)
        expect(lateRide(8)).toEqual(8)
    });
});
