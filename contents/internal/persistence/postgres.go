package persistence

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// connectTimeout bounds how long Init waits for the database to become
// reachable on startup. The platform reports the database resource Ready as
// soon as the claim reconciles, but the underlying role + database can take
// up to ~1-2 minutes to actually exist (Crossplane/provider-sql provisioning
// lag). We retry well past that so the app rides out the lag rather than
// crash-looping. (Even if it does exceed this, Kubernetes restarts the pod and
// it connects on the next attempt.)
const connectTimeout = 180 * time.Second

func Init(databaseURL string) error {
	deadline := time.Now().Add(connectTimeout)
	var lastErr error
	for attempt := 1; ; attempt++ {
		pool, err := pgxpool.New(context.Background(), databaseURL)
		if err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			err = pool.Ping(ctx)
			cancel()
			if err == nil {
				db = pool
				if attempt > 1 {
					slog.Info("persistence connected", "attempt", attempt)
				}
				return nil
			}
			pool.Close()
		}
		lastErr = err
		if time.Now().After(deadline) {
			return fmt.Errorf("persistence: not reachable after %s: %w", connectTimeout, lastErr)
		}
		slog.Warn("persistence not ready, retrying", "attempt", attempt, "error", err)
		time.Sleep(3 * time.Second)
	}
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func DB() *pgxpool.Pool {
	return db
}
