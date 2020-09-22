function change(string) {
    const upper = string.toUpperCase();
    const template = "0".repeat(26).split("")

    for (let i = 0; i < string.length; ++i) {
        const code = upper.charCodeAt(i) - 65;
        template[code] = "1"
    }

    return template.slice(0, 26).join("")
}

module.exports = change;
