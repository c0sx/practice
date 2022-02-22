const compare = (aScore, bScore) => {
    const result = [0, 0];

    for (let i = 0; i < 3; ++i) {
        const a = aScore[i];
        const b = bScore[i];

        if (a > b) {
            result[0] += 1;
        } else if (b > a) {
            result[1] += 1;
        }
    }

    return result;
}

//     let mut result = [0, 0];
//
//     let mut i = 0;
//     while i < 3 {
//         let a = a_score[i];
//         let b = b_score[i];
//
//         if a > b {
//             result[0] += 1
//         } else if b > a {
//             result[1] += 1
//         }
//
//         i += 1;
//     }
//
//     result
