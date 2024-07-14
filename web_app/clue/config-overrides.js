module.exports = function override (config, env) {
    console.log('override')
    let loaders = config.resolve
    loaders.fallback = {
        // existing configs...
        "buffer": require.resolve("buffer/"),
        "http2": require.resolve("http2/"),
        "os": require.resolve("os-browserify/browser"),
        "stream": require.resolve("stream-browserify"),
        "zlib": require.resolve("browserify-zlib"),
        "http": require.resolve("stream-http"),
        "url": require.resolve("url/"),
        "process": require.resolve("process/browser"),
        "tls": require.resolve("tls"),
        "net": require.resolve("net"),
        "dns": require.resolve("dns"),
        "util": require.resolve("util/"),
        "fs": false,
        "path": require.resolve("path-browserify"),
        "assert": require.resolve("assert/"),
        "https": require.resolve("https-browserify"),
        "dgram": false,
   }
    
    return config
}