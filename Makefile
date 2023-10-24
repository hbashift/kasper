LOCAL_DB_DSN = "postgres://postgres@localhost:5432/uir?sslmode=disable"

jet:
	@PATH=$(LOCAL_BIN):$(PATH) jet -dsn $(LOCAL_DB_DSN) -path=./internal/generated/kasper