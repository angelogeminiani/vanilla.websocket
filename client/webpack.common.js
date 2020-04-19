const path = require("path");

module.exports = {

    entry: {
        // popup: path.join(__dirname, "src/popup/index.tsx"),
        // "background": path.join(__dirname, "src/launcher.ts"),
        // "background.min": path.join(__dirname, "src/launcher.ts")
    },

    output: {
        path: path.join(__dirname, "build/"),
        filename: (context) => {
            return "[name].js";
        }
    },

    plugins: [
        //new MiniCssExtractPlugin({
            // Options similar to the same options in webpackOptions.output
            // both options are optional
            //filename: '[name].css',
            //chunkFilename: '[id].css',
        //}),
    ],

    module: {
        rules: [
            {
                exclude: /node_modules/,
                // test: /\.tsx?$/,
                use: "ts-loader"
            },
            {
                exclude: /node_modules/,
                test: /\.s[ac]ss$/i,
                use: []
            }
        ]
    },
    resolve: {
        extensions: [".ts", ".js"]
    },


};