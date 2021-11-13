function sortAnimal(animal) {
    if (!animal) {
        return null;
    }

    return [...animal].sort((a, b) => {
        return a.numberOfLegs - b.numberOfLegs || a.name.localeCompare(b.name);
    });
}

module.exports = sortAnimal;
