const grid = require("./grid")

const slidePuzzle = (input, snapshots) => {
    const sequence = [];

    let i = 0;
    while (!isFinished(input) && i < 100) {
        i++;

        const currentMove = getMove(input, sequence)
        if (!currentMove) {
            return sequence;
        }

        sequence.push(currentMove);
        move(input, currentMove);
        snapshots && snapshots.push([currentMove, input.map(row => row.slice())]);
    }

    return sequence;
}

// 1. Найти самую верхнюю строку, в которой нужно произвести изменения findActualRow
// 2. Найти в этом ряду тайл который нужно установить findActualTile
// 3. Найти в сетке тайл который нужно перенести findCellCoordinates
// 4. Проверка алгоритма.

// найти самую верхнюю строку в которой нужно произвести изменения
// найти следующий тайл, который нужно поместить в верхнуюю строку
// если это не последняя цифра строки, то просто разместить в нужном поле не трогая уже зафиксированные части ранее
// для того что бы сдвинуть число в нужном направлении, нужно поставить в это место ноль путем прокрутки цифр
// если это последняя цифра в строке, то поместите ее на позицию прямо под ее правильным местом и пробел под нее
// далее хак который игнорирует фиксацию элементов: Вниз, вниз, право, вверх, лево, вверх,право, вниз, лево, вверх

// не трогать части которые уже были собраны -- помечать собранную строку перед поворотом как законченную
// собрать первую строку
// научиться перемещать цифру в нужную позицию с сохранением истории перемещения
// повернуть паззл

const findAvailableMoves = (input) => {
    const space = grid.findCoordinatesOfEmptyCell(input);

    const higher = getHigherTile(input, space);
    const leftHand = getLeftHandTile(input, space);
    const rightHand = getRightHandTile(input, space);
    const lower = getLowerTile(input, space);

    return [higher, leftHand, rightHand, lower]
}

const getHigherTile = (input, [rowIndex, cellIndex]) => {
    if (rowIndex <= 0 || isRowLinedUp(input, rowIndex - 1)) {
        return;
    }

    return getTile(input, rowIndex - 1, cellIndex)
}

const getLeftHandTile = (input, [rowIndex, cellIndex]) => {
    if (cellIndex <= 0 || isCellLinedUp(input, rowIndex, cellIndex - 1)) {
        return;
    }

    return getTile(input, rowIndex, cellIndex - 1)
}

const getRightHandTile = (input, [rowIndex, cellIndex]) => {
    const rowLength = input[rowIndex].length;

    if (cellIndex >= rowLength) {
        return;
    }

    return getTile(input, rowIndex, cellIndex + 1)
}

const getLowerTile = (input, [rowIndex, cellIndex]) => {
    if (rowIndex >= input.length - 1) {
        return;
    }

    return input[rowIndex + 1][cellIndex]
}

const getTile = (input, rowIndex, cellIndex) => input[rowIndex][cellIndex];

const isRowLinedUp = (input, rowIndex) => {
    // если rowIndex последние два ряда
    if (rowIndex === input.length - 1 || rowIndex === input.length - 2) {
        return false;
    }

    const rowLength = input[rowIndex].length;

    const expected = Array.from({ length: input[rowIndex].length }).map((one, index) => {
        return index + 1 + rowLength * rowIndex;
    });

    return input[rowIndex].every((one, index) => one === expected[index]);
}

const isCellLinedUp = (input, rowIndex, cellIndex) => {
    const rowLength = input[rowIndex].length;
    if (cellIndex === rowLength - 1 || cellIndex === rowLength - 2) {
        return false
    }

    const expected = rowLength * rowIndex + cellIndex + 1;
    return input[rowIndex][cellIndex] === expected
}

const getMove = (input, sequence) => {
    const moves = findAvailableMoves(input);
    const expected = findTargetTile(input)
    const prevMove = sequence[sequence.length - 1];
    if (moves.includes(expected) && prevMove !== expected) {
        return expected;
    }

    // выбрать число которое нужно поднять наверх
    return findTileToTop(input, moves, sequence);
}

// найти тайл который должен занять место пустого
const findTargetTile = (input) => {
    const [rowIndex, cellIndex] = grid.findCoordinatesOfEmptyCell(input);
    const rowLength = input[rowIndex].length;
    return rowIndex * rowLength + cellIndex + 1;
}

const findTileToTop = (input, [top, left, right, bottom], sequence) => {
    const [rowIndex] = grid.findCoordinatesOfEmptyCell(input);
    const expectedRow = input[rowIndex].map((one, index) => {
        return index + 1 + rowIndex * input[rowIndex].length;
    })

    const prevMove = sequence[sequence.length - 1];

    if (left && !expectedRow.includes(left) && left !== prevMove) {
        return left;
    }

    if (top && top !== prevMove) {
        return top;
    }

    if (right && right !== prevMove) {
        return right;
    }

    if (bottom && bottom !== prevMove) {
        return bottom;
    }
}

const move = (input, target) => {
    const [spaceRow, spaceCell] = grid.findCoordinatesOfEmptyCell(input);
    const [targetRow, targetCell] = grid.findTileCoordinates(input, target);

    input[spaceRow][spaceCell] = target;
    input[targetRow][targetCell] = 0
}

const isFinished = (input) => {
    return input.every((row, rowIndex) => {
        const multiplier = row.length * rowIndex;
        return row.every((cell, cellIndex) => {
            let expected = cellIndex + 1 + multiplier;
            if (expected === input.length * row.length) {
                expected = 0;
            }

            return cell === expected;
        });
    })
}

module.exports = {
    slidePuzzle,
    isFinished
}
