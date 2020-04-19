const path = require("path");
const merge = require("webpack-merge");
const common = require("./webpack.common.js");
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');

module.exports = merge(common, {
    mode: "production",

    entry: {
        "bundle.min": path.join(__dirname, "./src/launcher.ts")
    },

    optimization: {
        minimize: true,
        minimizer: [
            new UglifyJsPlugin({
                include: /\.min\.js$/,
                exclude: /node_modules/,
            })
        ]
    },

});