module.exports = {
  productionSourceMap: false,
  css: { extract: false, sourceMap: false },
  configureWebpack: {
    output: {
      filename: "[name].js",
      chunkFilename: "[name].js",
    },
    optimization: { splitChunks: false },
  },
  chainWebpack: (config) => {
    config.resolve.extensions.add(".yml").add(".yaml");
    config.module
      .rule("yaml")
      .test(/\.ya?ml$/)
      .use("json-loader")
      .loader("json-loader")
      .end()
      .use("yaml-loader")
      .loader("yaml-loader")
      .end();
  },
};
