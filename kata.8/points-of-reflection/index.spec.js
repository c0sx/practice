const symmetricPoint = require("./index")

describe("Points of Reflection", function() {
    it("tests", () => {
        expect(symmetricPoint([0, 0], [1, 1])).toEqual([2, 2]);
        expect(symmetricPoint([2, 6], [-2, -6])).toEqual([-6, -18]);
        expect(symmetricPoint([10, -10], [-10, 10])).toEqual([-30, 30]);
        expect(symmetricPoint([1, -35], [-12, 1])).toEqual([-25, 37]);
        expect(symmetricPoint([1000, 15], [-7, -214])).toEqual([-1014, -443]);
        expect(symmetricPoint([0, 0], [0, 0])).toEqual( [0, 0]);
    });
});

