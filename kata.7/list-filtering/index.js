function filter_list(l) {
    return l.filter(item => Number.isFinite(item))
}

module.exports = filter_list;

