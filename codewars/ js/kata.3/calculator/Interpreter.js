class Interpreter {
    constructor(parser) {
        this.parser = parser;
    }

    interpret() {
        const tree = this.parser.parse();
        return tree.evaluate()
    }
}

module.exports = Interpreter;
