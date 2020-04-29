const regex = /^(\d{4}|\d{6})$/;
function validatePIN (pin) {
    return regex.test(pin);
}

console.log(validatePIN("1234")); // true
console.log(validatePIN("12345")); // false
console.log(validatePIN("a234")); // false
