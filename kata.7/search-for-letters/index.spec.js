const change = require("./index")

describe("Search for Letters", () => {
    it("tests", () => {
        expect(change("a **&  bZ")).toEqual("11000000000000000000000001")
        expect(change("abcdefghijklmnopqrstuvwxyz")).toEqual("11111111111111111111111111")
        expect(change("1234567890")).toEqual("00000000000000000000000000")
        expect(change("qQwWeEабва")).toEqual("00001000000000001000001000")
    })
})
