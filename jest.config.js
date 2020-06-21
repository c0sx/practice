module.exports = {
    clearMocks: true,
    collectCoverage: true,
    collectCoverageFrom: ["**/*.js"],
    coverageReporters: ["lcov", "html", "text"],
    coverageDirectory: "coverage",
    coveragePathIgnorePatterns: [
        "/node_modules/",
        "/coverage",
        "/jest.config.js"
    ],
    moduleDirectories: [
        "node_modules"
    ],
    testRegex: "(/__tests__/.*|(\\.|/)(test|spec))\\.(js?)$",
    moduleFileExtensions: ["js"],
    testEnvironment: "node",
    transformIgnorePatterns: [
        "/node_modules/",
        "/coverage"
    ],
};
