LOCAL_DB_DSN = "postgres://postgres@localhost:5432/uir_draft?sslmode=disable"

jet:
	@PATH=$(LOCAL_BIN):$(PATH) jet -dsn $(LOCAL_DB_DSN) -path=./internal/generated/kasper