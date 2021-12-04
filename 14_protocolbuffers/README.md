# Protocol buffers

In order to use protocol buffers, first you need to install the `protoc` 
to be available in your command line. Follow the instructions from the official
documentation for your OS [here](https://developers.google.com/protocol-buffers/docs/downloads).

The `.go` corresponding to the `.proto` examples are already available in the
folder. If you want to compile them you can run:

`protoc --go_out=$GOPATH/src *.proto`

Remember than depending on where the generated `.go` files are stored you 
need your imports to be set accordingly.
