export ROOT_MOD=github.com/777continue/gomall
.PHONY: server  #target
server:  # command 
	@cd app/frontend && air

.PHONY: rpc-user  #target
user:  # command 
	@cd app/user && air

.PHONY: rpc-product  #target
product:  # command 
	@cd app/product && air

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server  --type HTTP --service frontend --module github.com/777continue/gomall/app/frontend --idl ../../idl/frontend/category_page.proto -I ../../idl

.PHONY: gen-user
gen-user-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/777continue/gomall/rpc_gen --I ../idl --idl ../idl/user.proto
gen-user-server:		# .eg   --pass "-use client module path"
	@cd app/user && cwgo server --type RPC --service user --module github.com/777continue/gomall/app/user --pass "-use github.com/777continue/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
wcnm: 
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/product.proto


