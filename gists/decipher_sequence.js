const getAlphabet = () => {
    let alphabet = "";
    let i = "а";
    while (i !== "я") {
        if (i === "е") {
            alphabet += "е";
            i = "ё";
        }
        else if (i === "ё") {
            alphabet += "ё";
            i = "ж";
        }
        else {
            alphabet += i;
            i = String.fromCharCode(i.charCodeAt() + 1);
        }
    }

    return alphabet + "я";
};

const switchAlphabet = (alphabet, step) => {
    let carouseled = "";
    for (let index = 0; index < alphabet.length; ++index) {
        let nextIndex = index + step;
        if (nextIndex >= alphabet.length) {
            nextIndex = nextIndex % alphabet.length;
        }

        carouseled += alphabet[nextIndex];
    }

    return carouseled;
};

const getRow = (keyChar) => {
    const alphabet = getAlphabet();
    const step = alphabet.indexOf(keyChar);
    return switchAlphabet(alphabet, step);
};

const getDecipherChar = (row, cipherChar) => {
    const index = row.indexOf(cipherChar);
    const alphabet = getAlphabet();
    const column = switchAlphabet(alphabet, index);
    return column[0];
};

const getSequence = (str1, str2) => {
    const arr1 = str1.split(",");
    const arr2 = str2.split(",");

    let i = 0;
    let arr3 = arr2.map(n => i = arr2[i]);
    let result = "%s-%s-%s-%s-%s-%s";

    arr3.forEach(n => {
        result = result.replace("%s", arr1[n]);
    });

    return result;
};

const cipherText = "чмщъы ьччфя рдже йаиёсбпз ксспыф сж гхвчи ъцэлз, эч жд оюфзаззе жчклуг ёсш сёюнуцпн гпцвдд: юмэц,сас,лнскщщ,мбп,шуыы,эчае 56,25,21,57,26,54. оюхщцющ сй хдбыфхз ш кёыясе ц уну";
// const cipherText = "пщйссфужчп бсёыдьд";
const key = "ПИФИЯ".toLowerCase();
const alphabet = getAlphabet();

let result = "";
let j = 0;
for (let index = 0; index < cipherText.length; ++index) {
    const cipherChar = cipherText[index];

    if (j >= key.length) {
        j = 0;
    }

    if (alphabet.indexOf(cipherChar) === -1) {
        result += cipherChar;
        continue;
    }

    const keyChar = key[j];
    const row = getRow(keyChar);
    result += getDecipherChar(row, cipherChar);
    j++;
}

console.log(result);
console.log(getSequence("56,25,21,57,26,54", "1,3,4,2,5,0"));
