const Token = require("./Token")

const mapper = [
    {
        condition: token => token === "+",
        action: () => Token.Plus()
    },
    {
        condition: token => token === "-",
        action: () => Token.Minus()
    },
    {
        condition: token => token === "*",
        action: () => Token.Multiply()
    },
    {
        condition: token => token === "/",
        action: () => Token.Division()
    },
    {
        condition: token => token === "(",
        action: () => Token.LParen()
    },
    {
        condition: token => token === ")",
        action: () => Token.RParen()
    },
    {
        condition: token => Number.isInteger(Number(token)),
        action: token => Token.Integer(token)
    },
    {
        condition: token => token === undefined,
        action: () => Token.EOF()
    }
]

class Lexer {
    constructor(text) {
        this.position = 0;
        this.text = text.split("").filter(char => char !== " ")
    }

    advance() {
        this.position += 1;
    }

    int() {
        let result = ""
        while (Number.isInteger(Number(this.text[this.position]))) {
            result += this.text[this.position];
            this.advance();
        }

        return Number(result)
    }

    token() {
        if (Number.isInteger(Number(this.text[this.position]))) {
            return Token.Integer(this.int());
        }

        const char = this.text[this.position];
        const action = mapper.find(one => one.condition(char));
        this.advance();
        return action.action(char);
    }
}

module.exports = Lexer;
