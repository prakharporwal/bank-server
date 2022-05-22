package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/prakharporwal/back-server/db/mock"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()
	mockStore := mock.NewMockStore(mockCtrl)

	t.Run("test 1", func(t *testing.T) {
		// mockStore.Query("SELECT * FROM accounts WHERE id=$(1)", gomock.Eq("prakhar"))
		// mockStore.EXPECT().Query(gomock.Any(), gomock.Eq("prakhar")).Times(1)

		somedbDriver := "acd://some-random-addr"
		somedbSource := "postgres"

		conn, err := sql.Open(somedbDriver, somedbSource)
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}

		store := NewSQLStore(conn)

		assert.NoError(t, store.GetQuery(gomock.Any().String(), gomock.Eq("prakhar")))

		mockStore.EXPECT().Query(gomock.Any(), gomock.Eq("prakhar")).Times(1)
	})

}
