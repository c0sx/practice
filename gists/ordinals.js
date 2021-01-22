const numbers = new Map();
numbers.set("I", 1);
numbers.set("II", 2);
numbers.set("III", 3);
numbers.set("IV", 4);
numbers.set("V", 5);
numbers.set("VI", 6);
numbers.set("VII", 7);
numbers.set("VIII", 8);
numbers.set("IX", 9);
numbers.set("X", 10);
numbers.set("XX", 20);
numbers.set("XXX", 30);
numbers.set("XL", 40);
numbers.set("L", 50);

const names = ["George VI", "William II", "Elizabeth I", "William I", "William IX", "William XI", "William L", "William XL", "Admiral XXVIII"];

const ordinalToDecimal = ordinal => {
    const symbols = [ ... ordinal.split("") ];
    let decimal = 0;
    let stack = "";
    while (symbols.length) {
        const nextDigit = stack + symbols[0];
        if (!numbers.has(nextDigit)) {
            decimal += numbers.get(stack);
            stack = "";
        }

        stack += symbols.shift();
    }

    decimal += numbers.get(stack);
    if (stack.length > 0) {
        decimal += numbers.get(stack);
    }

    return decimal;
};

const getSortedList = (names) => {
    const grouped = names.reduce((map, royalName) => {
        const [name] = royalName.split(" ");
        let names = map.get(name);
        if (!names) {
            names = [];
            map.set(name, names);
        }

        names.push(royalName);
        return map;
    }, new Map());

    let resultNames = [];
    Array.from(grouped.keys()).sort().forEach(name => {
        const royalNames = grouped.get(name);
        const sortedRoyalNames = royalNames.sort((a, b) => {
            const [aName, aOrdinal] = a.split(" ");
            const [bName, bOrdinal] = b.split(" ");

            const aDecimal = ordinalToDecimal(aOrdinal);
            const bDecimal = ordinalToDecimal(bOrdinal);

            return aDecimal - bDecimal;
        });

        resultNames = [].concat(resultNames, sortedRoyalNames);
    });

    return resultNames;
}

console.log(getSortedList(names));
