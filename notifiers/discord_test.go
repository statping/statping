package notifiers

import (
	"github.com/hunterlong/statup/core/notifier"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	DISCORD_URL    = os.Getenv("DISCORD_URL")
	discordMessage = `{"content": "The Discord notifier on Statup has been tested!"}`
)

func init() {
	DISCORD_URL = os.Getenv("DISCORD_URL")
	discorder.Host = DISCORD_URL
}

func TestDiscordNotifier(t *testing.T) {
	t.Parallel()
	if DISCORD_URL == "" {
		t.Log("Discord notifier testing skipped, missing DISCORD_URL environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Discord", func(t *testing.T) {
		discorder.Host = DISCORD_URL
		discorder.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifier(discorder)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", discorder.Author)
		assert.Equal(t, DISCORD_URL, discorder.Host)
	})

	t.Run("Load Discord Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("Discord Within Limits", func(t *testing.T) {
		ok, err := discorder.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Discord Send", func(t *testing.T) {
		err := discorder.Send([]byte(discordMessage))
		assert.Nil(t, err)
	})

	t.Run("Discord Queue", func(t *testing.T) {
		go notifier.Queue(discorder)
		time.Sleep(1 * time.Second)
		assert.Equal(t, DISCORD_URL, discorder.Host)
		assert.Equal(t, 0, len(discorder.Queue))
	})

}
