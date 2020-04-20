const path = require("path");
const merge = require("webpack-merge");
const common = require("./webpack.common.js");


module.exports = merge(common, {
    mode: "development",

    entry: {
        "vws": path.join(__dirname, "./src/launcher.ts")
    },

    devtool: "source-map"

});