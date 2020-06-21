const filter_list = require("./index");

describe("List Filtering", () => {
    it("tests", () => {
        expect(filter_list([1, 2, "a", "b"])).toEqual([1, 2])
        expect(filter_list([1, "a", "b", 0, 15])).toEqual([1, 0, 15])
        expect(filter_list([1, 2, "aasf", "1", "123", 123])).toEqual([1, 2, 123])
        expect(filter_list(["a", "b", "1"])).toEqual([])
        expect(filter_list([1, 2, "a", "b"])).toEqual([1, 2])
    });
})
