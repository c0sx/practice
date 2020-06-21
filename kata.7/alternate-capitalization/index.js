function capitalize(s) {
    const [left, right] = s.split("").reduce(([left, right], current, index) => {
        if (index % 2 === 0) {
            left.push(current.toUpperCase())
            right.push(current)
        }
        else {
            left.push(current);
            right.push(current.toUpperCase())
        }

        return [left, right];
    }, [[], []]);

    return [left.join(""), right.join("")];
}

module.exports = capitalize;
