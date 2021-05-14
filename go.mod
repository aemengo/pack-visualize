module github.com/aemengo/pack-visualize

go 1.15

require (
	github.com/buildpacks/pack v0.18.1
	github.com/docker/docker v20.10.0-beta1.0.20201110211921-af34b94a78a1+incompatible
	github.com/fatih/color v1.10.0
	github.com/gdamore/tcell/v2 v2.2.0
	github.com/google/go-containerregistry v0.5.0 // indirect
	github.com/rivo/tview v0.0.0-20210312174852-ae9464cc3598
)

replace github.com/buildpacks/pack => ../pack
