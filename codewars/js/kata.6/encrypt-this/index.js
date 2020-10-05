function encryptThis(str) {
    return str.split(" ").map(word => {
        let encryption = word.split("");
        encryption[0] = encryption[0].charCodeAt(0);

        if (encryption.length > 1) {
            const t = encryption[1];
            encryption[1] = encryption[encryption.length - 1];
            encryption[encryption.length - 1] = t;
        }

        return encryption.join("")
    }).join(" ")
}

module.exports = encryptThis;
