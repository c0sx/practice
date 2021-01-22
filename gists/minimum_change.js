// посчитать минимальное количество монет для сдачи

const getPreviousValue = (table, rowIndex, cellIndex) => getValueFromCell(table, rowIndex - 1, cellIndex);

const getValueFromCell = (table, rowIndex, cellIndex) => {
    if (rowIndex < 0 || cellIndex < 0) {
        return [];
    }

    return table[rowIndex][cellIndex];
}

const calculateCurrentValue = (accumulatedRow, value,  capacity) => {
    const difference = capacity - value;
    if (difference === 0) {
        return [value]
    }

    const checked = accumulatedRow[difference - 1] || [];
    if (checked.length === 0) {
        return []
    }

    return [...checked, value];
}

const compareMinimumCoins = (current, previous) => {
    if (current.length === 0 && previous.length > 0) {
        return previous;
    }

    if (current.length > 0 && previous.length > 0) {
        return current.length <= previous.length ? current : previous
    }

    return current;
}

const calculateMinimumCoins = (values, maxCapacity) => {
    const table = Array(values.length).fill([]).map(() => Array(maxCapacity).fill([]));
    for (let rowIndex = 0; rowIndex < table.length; ++rowIndex) {
        const value = values[rowIndex];
        const row = table[rowIndex];

        for (let cellIndex = 0; cellIndex < row.length; ++cellIndex) {
            const currentCapacity = cellIndex + 1;
            const previous = getPreviousValue(table, rowIndex, cellIndex);
            const current = calculateCurrentValue(table[rowIndex], value, currentCapacity);

            table[rowIndex][cellIndex] = compareMinimumCoins(current, previous);
        }
    }

    return table;
}

/**
 * const change = 92;
 * const values = [3, 5, 10]
 * @param {number} values номинал валюты
 * @param {number[]} change сдача
 * @return {number[]} массив валюты
 */
const getCoins = (values, change) => {
    const table = calculateMinimumCoins(values, change);
    const row = table[table.length - 1] || [];
    return row[row.length - 1] || [];
}
