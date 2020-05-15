const domainName = require("./")

describe("Extract the Domain name from a URL", () => {
    it("tests", () => {
        expect(domainName("http://google.com")).toEqual("google");
        expect(domainName("http://google.co.jp")).toEqual( "google");
        expect(domainName("www.xakep.ru")).toEqual( "xakep");
        expect(domainName("https://youtube.com")).toEqual( "youtube");
    });
});
