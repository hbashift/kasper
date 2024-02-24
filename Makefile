LOCAL_DB_DSN = "postgres://postgres@localhost:5432/new_uir?sslmode=disable"

jet:
	@PATH=$(LOCAL_BIN):$(PATH) jet -dsn $(LOCAL_DB_DSN) -path=./internal/generated/new_kasper