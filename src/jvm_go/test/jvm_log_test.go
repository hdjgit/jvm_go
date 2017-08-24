package test

import (
	"testing"
	"jvm_go/log"
)

func TestLog(t *testing.T)  {
	log.InitLog()
	log.Debugf("hello world! %s","haha")
}

