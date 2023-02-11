package entity_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/tenant/entity"
	"testing"
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
