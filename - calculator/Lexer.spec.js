const Lexer = require("./Lexer");

describe("Calculator.Lexer", () => {
    it("should tokenize correctly", () => {
        const expression = "2 + 2 * 2";
        const lexer = new Lexer(expression);

        const tokens = collectTokens(lexer);
        expect(tokens).toEqual([2, "+", 2, "*", 2])
    });

    it("should tokenize correctly with parenthesis", () => {
        const expression = "2 + (2 + 3) * 2";
        const lexer = new Lexer(expression);

        const tokens = collectTokens(lexer);
        expect(tokens).toEqual([2, "+", "(", 2, "+", 3, ")", "*", 2])
    });
})

const collectTokens = (lexer) => {
    const tokens = [];
    while (true) {
        const token = lexer.token();
        if (token === undefined) {
            break;
        }

        tokens.push(token);
    }

    return tokens.map(token => token.value)
}
