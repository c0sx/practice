const chain = require("./index");

function sum(x, y) {

    return x + y;
}

function double(x) {
    return sum(x, x);
}

function minus (x, y) {
    return x - y;
}

function addOne(x) {
    return sum(x, 1);
}

describe("Born to be Chained", function() {
    let c;
    let c1;

    beforeEach(() => {
        c = chain({ sum, minus, double, addOne });
        c1 = c.sum(1, 2)
    })

    it("should sum correctly1", function() {
        const c = chain({ sum, minus, double, addOne })
        expect(c.sum(3, 4).execute()).toEqual(7);
        expect(c.sum(1, 2).execute()).toEqual(3);
    });

    it("should chain correctly1", () => {
        expect(c1.execute()).toEqual(3)
        expect(c1.double().execute()).toEqual(6);
        expect(c1.addOne().execute()).toEqual(4)
        expect(c1.execute()).toEqual(3);
    });

    it("should chain correctly2", () => {
        const c2 = c1.sum(5);
        expect(c2.addOne().execute()).toEqual(9);
        expect(c2.sum(3).execute()).toEqual(11);
        expect(c2.execute()).toEqual(8);
        expect(c1.execute()).toEqual(3)
    });
});
