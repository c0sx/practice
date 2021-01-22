const fromCamelCase = string => {
    const parts = [];
    const arr = string.split("");

    while (arr.length > 0) {
        const index = arr.findIndex(one => one !== one.toLowerCase());
        if (index === -1) {
            parts.push(arr.join(""));
            return parts
        }

        parts.push(arr.splice(0, index).join(""));
        arr[0] = arr[0].toLowerCase();
    }

    return parts;
}

const fromKebabCase = string => string.split("-")
const fromSnakeCase = string => string.split("_");

const toCamelCase = value => {
    const [first, ...other] = value;
    const converted = other.map(one => {
        return one.charAt(0).toUpperCase() + one.slice(1)
    })

    return [first, ...converted].join("")
}

const toSnakeCase = value => value.join("_")
const toKebabCase = value => value.join("-")

const targets = new Map([
    ["snake", toSnakeCase],
    ["camel", toCamelCase],
    ["kebab", toKebabCase]
])

const parsers = new Map([
    ["camel", fromCamelCase],
    ["kebab", fromKebabCase],
    ["snake", fromSnakeCase],
])

function changeCase(identifier, targetCase) {
    if (identifier.length === 0) {
        return "";
    }

    const target = targets.get(targetCase);
    if (!target) {
        return;
    }

    const detectedCase = detectCase(identifier);
    if (!detectedCase) {
        return;
    }

    const parser = parsers.get(detectedCase);
    const common = parser(identifier);
    return target(common);
}

const detectCase = value => {
    if (isCamelCase(value)) {
        return "camel"
    }

    if (isKebabCase(value)) {
        return "kebab"
    }

    if (isSnakeCase(value)) {
        return "snake"
    }
}

const hasUpperCaseLetter = value => value !== value.toLowerCase();

const isCamelCase = value => {
    const split = value.split("");
    return ["-", "_"].every(one => !split.includes(one));
}

const isKebabCase = value => {
    const split = value.split("");
    return split.includes("-") && !split.includes("_") && !hasUpperCaseLetter(value)
}

const isSnakeCase = value => {
    const split = value.split("");
    return split.includes("_") && !split.includes("-") && !hasUpperCaseLetter(value)
}

module.exports = changeCase;
