package mysql

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestInit(t *testing.T) {
	klog.SetLevel(klog.LevelDebug)
	Init()

}
