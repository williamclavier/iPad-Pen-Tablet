# Wireless iPad Pen Tablet

This runs a webserver that can be connected to on the local network via the browser of the iPad. With the Apple Pencil, you will be able to control the mouse on your computer. This can be used for playing games or for photo/video editing.

## Running
```batch
# To run without configuration
go run main.go

# To change the host address and port, specify with --addr flag
go run main.go --addr localhost:8000

# To change the senitivity, specify with --sensitivity flag
go run main.go --addr localhost:8000 --sensitivity 75
```