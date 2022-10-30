const Lexer = require("./Lexer");
const Parser = require("./Parser");

describe("Calculator.Parser", () => {
    it("should parse correctly", () => {
        const lexer = new Lexer("14 + 2 * 3 - 6 / 2")
        const parser = new Parser(lexer);

        expect(parser.parse()).toMatchSnapshot()
    });
})
