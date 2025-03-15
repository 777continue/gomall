export ROOT_MOD=github.com/777continue/gomall
.PHONY: server user rproduct cart checkout payment order email

server:  # command 
	@cd app/gateway && air

user:  # command 
	@cd app/user && air

product:  # command 
	@cd app/product && air			

cart:  # command
	@cd app/cart && air

checkout:  # command
	@cd app/checkout && air

payment:  # command
	@cd app/payment && air

order:  # command
	@cd app/order && air

email:  # command
	@cd app/email && air


.PHONY: gen-gateway
gen-gateway:
	@cd app/gateway && cwgo server  --type HTTP --service frontend --module github.com/777continue/gomall/app/frontend --idl ../../idl/frontend/checkout_page.proto -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/777continue/gomall/rpc_gen --I ../idl --idl ../idl/user.proto
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

.PHONY: gen-order
gen-order: 
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/order.proto
	@mkdir app/order&& cd app/order && cwgo server --type RPC --service order --module ${ROOT_MOD}/app/order --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-email
gen-email: 
	@cd rpc_gen && cwgo client --type RPC --service email --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/email.proto
	@mkdir app/email&& cd app/email && cwgo server --type RPC --service email --module ${ROOT_MOD}/app/email --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/email.proto

.PHONY: vue
vue:
	@cd app/vue && npm start
