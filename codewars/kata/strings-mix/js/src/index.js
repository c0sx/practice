function mix(s1, s2) {
    const map1 = collect(normalize(s1));
    const map2 = collect(normalize(s2));

    const acc = new Map();
    process(acc, 1, map1);
    process(acc, 2, map2);

    return show(acc);
}

const normalize = s => {
    return s.replace(/[^a-z]/g, "");
}

const collect = (s) => {
    return s.split("").reduce((acc, char) => {
        const counter = acc.get(char) || 0;
        return acc.set(char, counter + 1);
    }, new Map());
};

const process = (acc, source, map) => {
    return [...map.entries()].reduce((acc, [char, counter]) => {
        if (counter <= 1) {
            return acc;
        }

        const item = acc.get(char);
        if (!item || item.counter < counter) {
            return acc.set(char, {
                char,
                counter,
                source,
            });
        }

        if (item && item.counter === counter) {
            return acc.set(char, {
                char,
                counter,
                source: 3,
            });
        }

        return acc;
    }, acc)
};

const show = (acc) => {
    return [...acc.values()]
        .sort(order)
        .map(data => {
            const source = data.source === 3 ? "=" : String(data.source);
            return `${source}:${data.char.repeat(data.counter)}`;
        })
        .join("/")
}

const order = (a, b) => {
    const byLength = b.counter - a.counter;
    if (byLength !== 0) {
        return byLength;
    }

    const bySource = a.source - b.source;
    if (bySource !== 0) {
        return bySource;
    }

    return a.char.localeCompare(b.char);
}

module.exports = mix;
