function domainName(url) {
    const [domain] = url.replace(/(http(s)?:\/\/)?(www\.)?/, "").split(".");
    return domain;
}

module.exports = domainName;
