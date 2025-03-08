package casbin

import (
	"sync"

	"github.com/casbin/casbin/v2"
)

var (
	EF         *casbin.Enforcer
	once       sync.Once
	modelPath  = "casbin/model.conf"
	policyPath = "casbin/policy.csv"
)

// InitCasbin 初始化Casbin enforcer
func InitCasbin() {
	var err error
	once.Do(func() {

		EF, err = casbin.NewEnforcer(modelPath, policyPath)
	})
	if err != nil {
		panic(err)
	}
}
