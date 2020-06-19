const Calculator = require("./index")

describe("Calculator", () => {
    let calculator;

    beforeEach(() => {
        calculator = new Calculator();
    })

    it("should find root operator correctly", () => {

    });

    it("tests", () => {
        expect(calculator.evaluate("127")).toEqual(127);
        expect(calculator.evaluate("2 + 3")).toEqual(5);
        expect(calculator.evaluate("2 - 3 - 4")).toEqual(-5);
        expect(calculator.evaluate("10 * 5 / 2")).toEqual(25);
        expect(calculator.evaluate("2 + 2 * 2")).toEqual(6)
        expect(calculator.evaluate("2 * 2 + 2")).toEqual(6)
        expect(calculator.evaluate("2 * (2 + 3)")).toEqual(10);
        expect(calculator.evaluate("2 + 2 * 3")).toEqual(8)
    });
});
