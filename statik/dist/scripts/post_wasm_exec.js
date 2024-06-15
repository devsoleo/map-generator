function runGo(argv) {
  const go = new Go()

  go.argv = argv

  WebAssembly.instantiateStreaming(fetch("./build/main.wasm"), go.importObject).then((result) => {
    go.run(result.instance)
  })
}