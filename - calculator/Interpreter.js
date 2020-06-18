const Num = require("./Num");
const BinOp = require("./BinOp");

class Interpreter {
    constructor(parser) {
        this.parser = parser;
    }

    visitBinOp(node) {
        if (node.op.type === "PLUS") {
            return node.left.value + node.right.value;
        }

        if (node.op.type === "MINUS") {
            return node.left.value - node.right.value;
        }

        if (node.op.type === "MULTIPLY") {
            return node.left.value * node.right.value;
        }

        if (node.op.type === "DIVISION") {
            return node.left.value / node.right.value;
        }
    }

    visitNum(node) {
        return node.value;
    }

    interpret() {
        const tree = this.parser.parse();
        if (tree instanceof BinOp) {
            return this.visitBinOp(tree)
        }

        return this.visitNum(tree)
    }
}

module.exports = Interpreter;
