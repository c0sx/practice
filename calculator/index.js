const Calculator = function() {
    this.operators = ["+", "-", "*", "/"]
    this.root = undefined;

    this.evaluate = string => {
        const splitted = string.split("").filter(one => one !== " ");
        const ast = this.buildAST(splitted);
        if (ast instanceof Operator) {
            return ast.evaluate()
        }
        
        return ast.evaluate();
    }

    this.buildAST = (arr) => {
        const indexOfOperator = arr.findIndex(symbol => this.operators.includes(symbol));
        if (indexOfOperator === -1) {
            return new Value(arr.join(""))
        }

        const operator = arr[indexOfOperator];
        const left = arr.slice(0, indexOfOperator);
        const right = arr.slice(indexOfOperator + 1)

        return new Operator(operator, this.buildAST(left), this.buildAST(right));
    }

    //    Operator
    //      / \
    // Value   Operator
    //          /
    //        Value
};

// построить ast
// исполнить ast

class Value {
    constructor(value) {
        this.value = value;
    }

    evaluate() {
        return this.value;
    }
}

class Operator {
    constructor(name, left, right) {
        this.value = name;
        this.left = left;
        this.right = right;
    }

    evaluate() {
        const left = this.left.evaluate();
        const right = this.right.evaluate();

        return eval(`${left}${this.value}${right}`);
    }
}

module.exports = Calculator;
