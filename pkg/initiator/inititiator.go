package initiator

import (
	"github.com/rs/zerolog"

	"github.com/quickfixgo/quickfix"

	"sylr.dev/fix/pkg/utils"
)

func Initiate(app quickfix.Application, settings *quickfix.Settings, logger *zerolog.Logger) (*quickfix.Initiator, error) {
	var msgStoreFactory quickfix.MessageStoreFactory

	if settings.GlobalSettings().HasSetting("SQLStoreDriver") {
		if driver, err := settings.GlobalSettings().Setting("SQLStoreDriver"); err == nil && driver == "sqlite3" {
			msgStoreFactory = quickfix.NewSQLStoreFactory(settings)
		}
	}

	if msgStoreFactory == nil {
		msgStoreFactory = quickfix.NewMemoryStoreFactory()
	}

	return quickfix.NewInitiator(app, msgStoreFactory, settings, utils.NewQuickFixLogFactory(logger))
}
