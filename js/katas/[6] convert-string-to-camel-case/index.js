function toCamelCase(str) {
    const [first, ...other] = str.replace(/_/g, "-").split("-");
    const converted = other.map(one => {
        return one.charAt(0).toUpperCase() + one.slice(1)
    })

    return [first, ...converted].join("")
}

module.exports = toCamelCase;
