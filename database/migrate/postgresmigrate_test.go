package migrate_test

import (
	"Edwardz43/godevframework/database/migrate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	var isSuccess bool

	migration := migrate.PostgresMigration{}

	isSuccess, err := migration.Migrate()

	assert.True(t, isSuccess)
	assert.Nil(t, err)
}
