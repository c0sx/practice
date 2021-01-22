const map = new Map([
    ["jabroni", "Patron Tequila"],
    ["school counselor", "Anything with Alcohol"],
    ["programmer", "Hipster Craft Beer"],
    ["bike gang member", "Moonshine"],
    ["politician", "Your tax dollars"],
    ["rapper", "Cristal"]
])

function getDrinkByProfession(param) {
    return map.get(param.toLowerCase()) || "Beer"
}

module.exports = getDrinkByProfession
