package entity_test

import (
	"testing"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/tenant/entity"
)

func TestNewUserSuccess(t *testing.T) {

	received, err := entity.Create()
	if err != nil {
		t.Errorf("Should not return an error: %s", err)
	}
	if received.GetCreatedAt().IsZero() {
		t.Errorf("CreatedAt must not be zero")
	}
}
