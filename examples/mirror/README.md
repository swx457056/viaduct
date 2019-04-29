# How to Run
#test the viaduct
#how to run viaduct
#Run the viaduct
#testinggggggggggggggggggggggggg
# to test botttttttttttttttt
# Testing bot
# testing
# bottttttttttttttttttt Test
# testinggggggggggggggg bottttttttttttttttt
#test test
# testing sample
#test botttttttttttttttttttttttttttttttttt
#TEST01
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
