
.PHONY: server  #target
svr:  # command 
	@cd app/frontend && air

.PHONY: rpc-server  #target
rpc-svr:  # command 
	@cd app/user && air

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server  --type HTTP --service frontend --module github.com/777continue/gomall/app/frontend --idl ../../idl/frontend/auth_page.proto -I ../../idl

.PHONY: gen-user-client
gen-user-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/777continue/gomall/rpc_gen --I ../idl --idl ../idl/user.proto

.PHONY: gen-user-server
gen-user-server:		# .eg   --pass "-use client module path"
	@cd app/user && cwgo server --type RPC --service user --module github.com/777continue/gomall/app/user --pass "-use github.com/777continue/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/user.proto

