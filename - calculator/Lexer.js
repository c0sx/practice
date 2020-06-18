const Token = require("./Token")

class Lexer {
    constructor(text) {
        this.position = 0;
        this.text = text.split("").filter(char => char !== " ")
        this.char = this.text[this.position]
    }

    advance() {
        this.position += 1;
        this.char = this.text[this.position]
    }

    int() {
        let result = ""
        while (Number.isInteger(Number(this.char))) {
            result += this.char;
            this.advance();
        }

        return Number(result)
    }

    token() {
        if (Number.isInteger(Number(this.char))) {
            return new Token("INTEGER", this.int());
        }

        const char = this.char;
        if (char === "+") {
            this.advance()
            return new Token("PLUS", "+")
        }

        if (char === "-") {
            this.advance();
            return new Token("MINUS", "-")
        }

        if (char === "*") {
            this.advance()
            return new Token("MULTIPLY", "*")
        }

        if (char === "/") {
            this.advance();
            return new Token("DIVISION", "/")
        }

        if (char === "(") {
            this.advance()
            return new Token("LPAREN", "(")
        }

        if (char === ")") {
            this.advance()
            return new Token("RPAREN", ")")
        }
    }
}

module.exports = Lexer;
