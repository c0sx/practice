const changeCase = require("./index")

describe("Convert All Cases", () => {
    it("tests", () => {
        expect(changeCase("snakeCase", "snake")).toEqual("snake_case")
        expect(changeCase("some-lisp-name", "camel")).toEqual("someLispName");
        expect(changeCase("map_to_all", "kebab")).toEqual("map-to-all");
        expect(changeCase("invalid-inPut_bad", "kebab")).toEqual(undefined);
        expect(changeCase("valid-input", "huh???")).toEqual(undefined);
        expect(changeCase("", "camel")).toEqual("");
        expect(changeCase("tyhjgwxclc", "camel")).toEqual("tyhjgwxclc")
        expect(changeCase("kekLul", "camel")).toEqual("kekLul")
    });
})
