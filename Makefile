export ROOT_MOD=github.com/777continue/gomall
.PHONY: server rpc-user rpc-product rpc-cart all-services

server:  # command 
	@cd app/frontend && air

rpc-user:  # command 
	@cd app/user && air

rpc-product:  # command 
	@cd app/product && air

rpc-cart:  # command
	@cd app/cart && air

rpc-checkout:  # command
	@cd app/checkout && air

rpc-payment:  # command
	@cd app/payment && air


all-services: server rpc-user rpc-product rpc-cart
	@echo "All services are started."

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server  --type HTTP --service frontend --module github.com/777continue/gomall/app/frontend --idl ../../idl/frontend/checkout_page.proto -I ../../idl

.PHONY: gen-user
gen-user-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/777continue/gomall/rpc_gen --I ../idl --idl ../idl/user.proto
gen-user-server:		# .eg   --pass "-use client module path"
	@cd app/user && cwgo server --type RPC --service user --module github.com/777continue/gomall/app/user --pass "-use github.com/777continue/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart: 
	@cd rpc_gen && cwgo client --type RPC --service cart --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module ${ROOT_MOD}/app/cart --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment: 
	@cd rpc_gen && cwgo client --type RPC --service payment --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/payment.proto
	@mkdir app/payment && cd app/payment && cwgo server --type RPC --service payment --module ${ROOT_MOD}/app/payment --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout: 
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@mkdir app/checkout && cd app/checkout && cwgo server --type RPC --service checkout --module ${ROOT_MOD}/app/checkout --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/checkout.proto


