const Num = require("./Num");
const BinOp = require("./BinOp")

class Parser {
    constructor(lexer) {
        this.lexer = lexer;
        this.token = this.lexer.token();
    }

    eat(type) {
        if (this.token.type === type) {
            this.token = this.lexer.token();
        }
    }

    factor() {
        const token = this.token;
        if (token.type === "INTEGER") {
            this.eat("INTEGER")
            return new Num(token)
        }
        else if (token.type === "LPAREN") {
            this.eat("LPAREN")
            const node = this.expression();
            this.eat("RPAREN")
            return node;
        }
    }

    term() {
        const nodeFactor = this.factor();
        let node
        while (["MULTIPLY", "DIVISION"].includes(this.token.type)) {
            const token = this.token;
            this.eat(token.type);

            node = new BinOp(nodeFactor, token, this.factor())
        }

        return node;

    }

    expression() {
        const nodeTerm = this.term()

        let node;
        while (["PLUS", "MINUS"].includes(this.token.type)) {
            const token = this.token;
            this.eat(token.type)

            node = new BinOp(nodeTerm, token, this.term())
        }

        return node;
    }

    parse() {
        return this.expression();
    }
}

module.exports = Parser;
