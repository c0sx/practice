const Lexer = require("./Lexer");
const Parser = require("./Parser");
const Interpreter = require("./Interpreter");

const Calculator = function() {
    this.evaluate = string => {
        const lexer = new Lexer(string);
        const parser = new Parser(lexer);
        const interpreter = new Interpreter(parser);
        return interpreter.interpret();
    }
};

module.exports = Calculator;
