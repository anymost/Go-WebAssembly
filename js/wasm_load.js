if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

(async function () {
    const go = new Go();
    const value = await WebAssembly.instantiateStreaming(fetch("/wasm"), go.importObject);
    let {
        module,
        instance
    } = value;
    await go.run(instance);
    instance = await WebAssembly.instantiate(module, go.importObject);
})();