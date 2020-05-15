const nodes = [];
const finished = new Set();
const keyMap = [8, 2, 4, 1, 2, 8, 5, 6, 8, 5, 6];

//  [4,1,3],
//  [2,8,0],
//  [7,6,5]

const slidePuzzle = (input) => {
    const target = findTargetState(input);
    nodes.push({
        node: input,
        length: 0,
    });

    //   return [8, 2, 4, 1, 2, 8, 5, 6, 8, 5, 6];
    generateGraph(target);
    return [];
};

const generateGraph = (target) => {
    let i = 0;
    while (!finished.has(JSON.stringify(target))) {
        console.log(i);
        const inputs = []
            .concat(nodes.find(row => row.length === i))
            .filter(Boolean);

        inputs.forEach(row => processGraph(row.node, i, target));
        i++;
    }

    console.log("target found", finished.has(JSON.stringify(target)), finished.size);
};

const processGraph = (input, length, target) => {
    const availableMoves = findAvailableMoves(input, target);
    // .filter(move => keyMap[length] === move);

    // console.log(input, availableMoves);

    const nextNodes = availableMoves
        .map(move => changeTile(input, move))
        .filter(node => !finished.has(JSON.stringify(node)));

    finished.add(JSON.stringify(input));
    console.log("finished", input);
    nextNodes.forEach(nextNode => {
        const index = nodes.findIndex(storedNode => {
            return JSON.stringify(nextNode) === JSON.stringify(storedNode);
        });

        if (index === -1) {
            return nodes.push({
                node: nextNode,
                length: length + 1,
            });
        }

        if (nodes[index].length > length + 1) {
            nodes[index].length = length + 1;
        }
    });
};

const changeTile = (input, tile) => {
    const [zeroRow, zeroCell] = findEmptyCell(input);
    const [targetRow, targetCell] = findCell(input, tile);

    const copy = input.map(row => [].concat(...row));
    const temp = copy[zeroRow][zeroCell];
    copy[zeroRow][zeroCell] = copy[targetRow][targetCell];
    copy[targetRow][targetCell] = temp;
    return copy;
};

const findAvailableMoves = (input, target) => {
    const space = findEmptyCell(input);

    const top = getTopMove(input, space, target);
    const left = getLeftMove(input, space, target);
    const right = getRightMove(input, space);
    const bottom = getBottomMove(input, space);

    return [top, left, right, bottom].filter(Boolean)
};

const getTopMove = (input, [row, cell], [firstRow]) => {
    if (row <= 0) {
        return undefined;
    }

    const topRowIndex = row - 1;
    if (topRowIndex === 0) {
        if (JSON.stringify(input[topRowIndex]) === JSON.stringify(firstRow)) {
            return undefined;
        }
    }

    return input[row - 1][cell];
};

const getLeftMove = (input, [row, cell], target) => {
    if (cell <= 0) {
        return undefined;
    }

    const columnIndex = cell - 1;
    if (columnIndex === 0) {
        const inputColumn = input.map(row => row[0]);
        const targetColumn = target.map(row => row[0]);
        if (JSON.stringify(inputColumn) === JSON.stringify(targetColumn)) {
            return undefined;
        }
    }

    return input[row][cell - 1];
};

const getRightMove = (input, [row, cell]) => {
    const rowSize = input[0].length;
    return cell >= rowSize ? undefined : input[row][cell + 1];
};

const getBottomMove = (input, [row, cell]) => {
    return row === input.length - 1 ? undefined : input[row + 1][cell];
};

const findCell = (input, needle) => {
    const row = input.find(row => row.includes(needle));
    const cell = row.findIndex(cell => cell === needle);
    return [input.findIndex(item => item[0] === row[0]), cell];
};

const findEmptyCell = input => findCell(input, 0);

const findTargetState = input => {
    const reduced = input.reduce((accumulator, current) => {
        accumulator = accumulator.concat(current);
        return accumulator;
    }, []);

    const sorted = reduced.sort((a, b) => a - b);
    const zero = sorted.shift();
    sorted.push(zero);

    return input.reduce((accumulator, current) => {
        accumulator.push(sorted.splice(0, current.length));
        return accumulator;
    }, []);
};

const test = [
    [4, 1, 3],
    [2, 8, 0],
    [7, 6, 5]
];
const result = slidePuzzle(test);

console.log(result, [8, 2, 4, 1, 2, 8, 5, 6, 8, 5, 6]);
