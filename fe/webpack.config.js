
let webpack = require("webpack");
let path = require("path");
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');

module.exports = {
	entry: {
		app: "./resources/assets/js/app.js",
		vendor: ["vue", "axios"],
	},
	mode: "development",

	output: {
		path: path.resolve(__dirname, "public/js"),
		filename: "[name].js",
		publicPath: "./public",
	},

	module: {
		rules: [
			{
				test: /\.js$/,
				exclude: /node_modules/,
				loader: "babel-loader",
			}
		],
	},

	resolve: {
		alias: {
			"vue$" : "vue/dist/vue.common.js"
		}
	},

	plugins: [],

	optimization: {
		splitChunks: false,
	}

	// optimization: {
	// 	splitChunks: {
	// 		// include all types of chunks
	// 		chunks: ["app", 'vendor'],
	// 	}
	// }

};

if (process.env.NODE_ENV === "production") {
	console.log("in production mode");

	module.exports.plugins.push(
		new UglifyJsPlugin()
	);

	module.exports.plugins.push(
		new webpack.DefinePlugin({
			"process.env": {
				NODE_ENV: "production",
			},
		})
	)
}

// if (process.env.NODE_ENV === "production") {
// 	module.exports.plugins.push(
// 		// new config.optimization.minimize();
// 		new webpack.optimize.UglifyJsPlugin({
// 			sourcemap: true,
// 			compress: {
// 				warnings: false,
// 			}
// 		})
// 	);
// }













