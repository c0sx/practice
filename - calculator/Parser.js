const Num = require("./Num");
const MultiplyOperation = require("./MultiplyOperation");
const PlusOperation = require("./PlusOperation");
const MinusOperation = require("./MinusOperation");
const DivisionOperation = require("./DivisionOperation")
const TokenType = require("./TokenType");

class Parser {
    constructor(lexer) {
        this.lexer = lexer;
        this.token = this.lexer.token();
    }

    factor() {
        const token = this.token
        if (token.type === TokenType.Integer) {
            this.token = this.lexer.token()
            return new Num(token)
        }
        else if (this.token.type === TokenType.LParen) {
            this.token = this.lexer.token()
            const node = this.expr();
            this.token = this.lexer.token()
            return node;
        }
    }

    term() {
        let node = this.factor();
        const operations = new Map([
            [TokenType.Multiply, (left, right) => new MultiplyOperation(left, right)],
            [TokenType.Division, (left, right) => new DivisionOperation(left, right)],
        ]);

        while ([TokenType.Multiply, TokenType.Division].includes(this.token.type)) {
            const operation = operations.get(this.token.type);
            this.token = this.lexer.token()
            node = operation(node, this.factor());
        }

        return node;

    }

    expr() {
        let node = this.term()
        const operations = new Map([
            [TokenType.Plus, (left, right) => new PlusOperation(left, right)],
            [TokenType.Minus, (left, right) => new MinusOperation(left, right)]
        ]);

        while ([TokenType.Plus, TokenType.Minus].includes(this.token.type)) {
            const operation = operations.get(this.token.type);
            this.token = this.lexer.token()
            node = operation(node, this.term())
        }

        return node;
    }

    parse() {
        return this.expr();
    }
}

module.exports = Parser;
