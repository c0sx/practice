const isSquare = require("./")

describe("isSquare", function() {
    it("should work for some examples", function(){
        expect(isSquare(-1)).toEqual(false) // , "-1: Negative numbers cannot be square numbers");
        expect(isSquare( 0)).toEqual(true) // , "0 is a square number (0 * 0)");
        expect(isSquare( 3)).toEqual(false) // , "3 is not a square number");
        expect(isSquare( 4)).toEqual(true) // , "4 is a square number (2 * 2)");
        expect(isSquare(25)).toEqual(true) // , "25 is a square number (5 * 5)");
        expect(isSquare(26)).toEqual(false) // , "26 is not a square number");
        expect(isSquare(42)).toEqual(false)
        expect(isSquare(23230573032496482275618623685)).toEqual(true)
    });
});
