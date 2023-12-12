WASMTIME_BACKTRACE_DETAILS = 1
WASMOPT_PATH = /tinygo/lib/binaryen/bin/wasm-opt

run:
	@tinygo version
	@$(WASMOPT_PATH) --version
	@WASMOPT="$(WASMOPT_PATH)" tinygo run -target=wasi main.go