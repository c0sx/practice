const staircase = n => {
    const result = [];

    for (let i = 0; i < n; ++i) {
        const amountOfSpaces = n - i - 1;
        const amountOfSymbols = n - amountOfSpaces;
        const spaces = " ".repeat(amountOfSpaces);
        const symbols = "#".repeat(amountOfSymbols);

        result.push(spaces + symbols);
    }

    return result;
}
