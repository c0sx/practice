class DivisionOperation {
    constructor(left, right) {
        this.left = left;
        this.right = right;
    }

    evaluate() {
        return this.left.evaluate() / this.right.evaluate();
    }
}

module.exports = DivisionOperation;
