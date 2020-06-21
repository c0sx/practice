const validParentheses = require("./index");

describe("Valid Parentheses", () => {
    it("tests", () => {
        expect(validParentheses("()")).toBeTruthy(); // true
        expect(validParentheses(")(()))")).toBeFalsy(); // false
        expect(validParentheses("(")).toBeFalsy(); // false
        expect(validParentheses("(())((()())())")).toBeTruthy(); // true
    });
})
