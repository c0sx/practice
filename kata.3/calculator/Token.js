const TokenType = require("./TokenType");

class Token {
    constructor(type, value) {
        this.type = type;
        this.value = value;
    }
}

Token.Plus = () => new Token(TokenType.Plus, "+");
Token.Minus = () => new Token(TokenType.Minus, "-");
Token.Multiply = () => new Token(TokenType.Multiply, "*");
Token.Division = () => new Token(TokenType.Division, "/");
Token.LParen = () => new Token(TokenType.LParen, "(");
Token.RParen = () => new Token(TokenType.RParen, ")")
Token.EOF = () => new Token(TokenType.EOF, undefined);
Token.Integer = value => new Token(TokenType.Integer, Number(value));

module.exports = Token;
