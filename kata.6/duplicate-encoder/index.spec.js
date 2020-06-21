const duplicateEncode = require("./index");

describe("Duplicate Encoder", () => {
    it("tests", () => {
        expect(duplicateEncode("din")).toEqual("(((");
        expect(duplicateEncode("recede")).toEqual("()()()");
        expect(duplicateEncode("Success")).toEqual(")())())")//,"should ignore case");
        expect(duplicateEncode("CodeWarrior")).toEqual("()(((())())");
        expect(duplicateEncode("Supralapsarian")).toEqual(")()))()))))()(")//,"should ignore case");
        expect(duplicateEncode("iiiiii")).toEqual("))))))")//,"duplicate-only-string")
        expect(duplicateEncode("(( @")).toEqual("))((");
    });
})
