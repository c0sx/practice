const uniTotal = require("./")

describe("Unicode Total", () => {
    it("no chars should return null ", () => {
        expect(uniTotal("")).toEqual(0)
    });

    it("should work with Single Letters", () => {
        expect(uniTotal("a")).toEqual(97);
        expect(uniTotal("b")).toEqual(98);
        expect(uniTotal("c")).toEqual(99);
        expect(uniTotal("d")).toEqual(100);
        expect(uniTotal("e")).toEqual(101);
    });

    it("should work with multiple letters ", () => {
        expect(uniTotal("aaa")).toEqual(291);
    });

    it("should work with sentences and spaces", () => {
        expect(uniTotal("Mary Had A Little Lamb")).toEqual(1873);
    });
})
