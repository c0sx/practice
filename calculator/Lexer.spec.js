const Lexer = require("./Lexer");

describe("Calculator.Lexer", () => {
    it("should tokenize correctly", () => {
        const expression = "2 + 2 * 2";
        const lexer = new Lexer(expression);

        const tokens = [];
        while (true) {
            const token = lexer.token();
            if (token === undefined) {
                break;
            }

            tokens.push(token);
        }

        expect(tokens).toEqual(["2", "+", "2", "*", "2"])
    });
})
