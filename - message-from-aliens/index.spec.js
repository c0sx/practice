const decode = require("./index");

describe("Message from Aliens", () => {
    it("tests", () => {
        expect(decode("]()]|_]|_]][-]|-|]")).toEqual("hello");
        expect(decode("{|^{|{{|_{]3{")).toEqual("blip");
        expect(decode("..[-.|_.|^....().[-.|^.__..|)...|.|^.|_|..~|~._\\~.__...[-..|.|)..")).toEqual("die stupid people");
        expect(decode("'''_\\~'|_|'()'|''('|'|_'[-'|)''__'_\\~'/<'()'()'|_'''__'|\\|'|''/\\'/?']3'__''/?'|_|''()'`/''")).toEqual( "your brain looks delicious");
        expect(decode("}/\\}~|~}/\\}/<}__}|)}}}[-}~|~}/\\}(}|}|_}|^}|_|}|)}__}|)}}}|\\|}|}/=}__}()}}}~|~}__}`/}/?}}~|~}")).toEqual("try to find duplicated kata");
    });
})
