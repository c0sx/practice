const findTileCoordinates = require("./findTileCoordinates");
const findCoordinatesOfEmptyCell = require("./findCoordinatesOfEmptyCell");
const findActualRow = require("./findActualRow");
const findCoordinatesOfActualTile = require("./findCoordinatesOfActualTile");
const findActualTile = require("./findActualTile");
const findTile = require("./findTile");

const rotate = require("./rotate");

const moveTile = require("./moveTile")

module.exports = {
    findCoordinatesOfEmptyCell,
    findCoordinatesOfActualTile,
    findTileCoordinates,
    findActualRow,
    findActualTile,
    findTile,
    rotate,
    moveTile
}
