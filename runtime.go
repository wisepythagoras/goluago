package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/wisepythagoras/goluago/native"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type Runtime struct {
	L    *lua.LState
	Path string
	Dir  fs.DirEntry
	Main fs.DirEntry
}

func (r *Runtime) GetPath(withMain bool) string {
	scriptPath := r.Path

	if withMain {
		scriptPath = filepath.Join(scriptPath, r.Main.Name())
	}

	return scriptPath
}

func (r *Runtime) Init() {
	r.L = lua.NewState()

	r.L.SetGlobal("fetch", luar.New(r.L, native.Fetch))
	r.L.SetGlobal("len", luar.New(r.L, native.Len))
	r.L.SetGlobal("runAsync", luar.New(r.L, native.RunAsync))

	// Allow requiring lua files from the plugin's directory.
	pkg := r.L.GetGlobal("package")
	newPath := fmt.Sprintf("%s/?.lua;%s", r.GetPath(false), pkg.String())
	r.L.SetField(pkg, "path", luar.New(r.L, newPath))

	// Run the extension's main file.
	if err := r.L.DoFile(r.GetPath(true)); err != nil {
		panic(err)
	}
}
