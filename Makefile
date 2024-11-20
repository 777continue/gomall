.PHONY: cd1
cd1:
	@cd app/frontend 

.PHONY: svr
svr:
	@cd app/frontend && air


.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server  --type HTTP --service frontend --module github.com/777continue/gomall/app/frontend --idl ../../idl/frontend/auth_page.proto -I ../../idl