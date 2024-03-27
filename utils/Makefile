CONFIG_PATH = ${HOME}/.prolog/
.PHONY: init 
init:
	mkdir -p ${CONFIG_PATH}
.PHONY: gencert 
gencert:
	cfssl gencert -initca certificates/ca-csr.json | cfssljson -bare ca
	cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=certificates/ca-config.json -profile=server certificates/server-csr.json | cfssljson -bare server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=certificates/ca-config.json \
		-profile=client \
		certificates/client-csr.json | cfssljson -bare client
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=certificates/ca-config.json \
		-profile=peer \
		certificates/peer-csr.json | cfssljson -bare peer
	mv *.pem *.csr ${CONFIG_PATH}
.PHONY: test 
test:
	go test -race ./...
.PHONY: compile
compile:
	protoc api/v1/*.proto \ 
		--go_out=. \
		--go-grpc_out = . \
		--go_opt = paths = source_relative \
		--go-grpc_opt = paths = source_relative\
		--proto_path = .

