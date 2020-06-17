const True = T => F => T;
const False = T => F => F;

const Not = A => A(False)(True)
const And = A => B => A(B)(A)
const Or = A => B => A(A)(B)
const Xor = A => B => A(Not(B))(B)

module.exports = {
    Not,
    And,
    Or,
    Xor
}
