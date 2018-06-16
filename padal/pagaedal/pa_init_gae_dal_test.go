package pagaedal

import (
	"testing"
	"github.com/prizarena/prizarena-public/padal"
)

func TestRegisterDal(t *testing.T) {
	RegisterDal()
	if padal.DB == nil {
		t.Error("padal.DB == nil")
	}
	if padal.Tournament == nil {
		t.Error("padal.Tournament == nil")
	}
	if padal.HandleWithContext == nil {
		t.Error("padal.HandleWithContext == nil")
	}
}
