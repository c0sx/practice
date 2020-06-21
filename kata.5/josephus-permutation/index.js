function josephus(items, k) {
    if (items.length === 0) {
        return [];
    }

    const results = [];
    let step = 1;
    while (items.length > 0) {
        for (let i = 0; i < items.length; ++i) {
            if (step % k === 0) {
                results.push(items[i])
                items[i] = undefined;
            }

            step++;
        }

        items = items.filter(item => item !== undefined)
    }

    return results;
}

module.exports = josephus;
