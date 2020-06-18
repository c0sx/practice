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
            return this.int();
        }

        const char = this.char;
        this.advance();
        return char;
    }
}

module.exports = Lexer;
