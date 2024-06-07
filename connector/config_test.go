package connector

import (
	"testing"

	"github.com/joshmedeski/sesh/dir"
	"github.com/joshmedeski/sesh/home"
	"github.com/joshmedeski/sesh/lister"
	"github.com/joshmedeski/sesh/model"
	"github.com/joshmedeski/sesh/tmux"
	"github.com/joshmedeski/sesh/zoxide"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestConfigStrategy(t *testing.T) {
	mockDir := new(dir.MockDir)
	mockHome := new(home.MockHome)
	mockLister := new(lister.MockLister)
	mockTmux := new(tmux.MockTmux)
	mockZoxide := new(zoxide.MockZoxide)
	c := &RealConnector{
		mockDir,
		mockHome,
		mockLister,
		mockTmux,
		mockZoxide,
		model.Config{},
	}
	mockTmux.On("AttachSession", mock.Anything).Return("attaching", nil)
	mockZoxide.On("Add", mock.Anything).Return(nil)

	t.Run("should create and attach to config session", func(t *testing.T) {
		mockTmux.On("IsAttached").Return(false)
		mockLister.On("FindConfigSession", "tmux config").Return(model.SeshSession{
			Name: "tmux config",
			Path: "/Users/joshmedeski/c/dotfiles/.config/tmux",
		}, true)
		connection, err := configStrategy(c, "tmux config")
		assert.Nil(t, err)
		assert.Equal(t, "tmux config", connection.Session.Name)
	})
}
