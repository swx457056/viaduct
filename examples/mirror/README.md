# How to Run
#test the viaduct
#Run the viaduct
# Testing bot
# testing
# bottttttttttttttttttt Test
# testinggggggggggggggg bottttttttttttttttt
#test test
# testing sample
#test botttttttttttttttttttttttttttttttttt
#sample testing
#bot TEsting
    cd examples/mirror
    go build

## quic
- start server

	./mirror --cmd-type=server --type=quic --addr=localhost:9890
- start client

	./mirror --cmd-type=client --type=quic --addr=localhost:9890

## websocket
- start server

	./mirror --cmd-type=server --type=websocket --addr=localhost:9890
	
- start client

	./mirror --cmd-type=client --type=websocket --addr=wss://localhost:9890/test
