package admin

import (
	"context"
	"testing"

	"uir_draft/internal/pkg/configs"
	"uir_draft/internal/pkg/repository"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func TestService_GenerateReportOne(t *testing.T) {
	err := initConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	db, err := configs.InitPostgresDB(ctx, configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		panic(err)
	}
	type fields struct {
		clientRepo ClientRepository
		marksRepo  MarksRepository
		userRepo   UsersRepository
		db         *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				clientRepo: repository.NewClientRepository(),
				marksRepo:  repository.NewMarksRepository(),
				userRepo:   repository.NewUsersRepository(),
				db:         db,
			},
			args:    args{ctx: ctx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				clientRepo: tt.fields.clientRepo,
				marksRepo:  tt.fields.marksRepo,
				userRepo:   tt.fields.userRepo,
				db:         tt.fields.db,
			}
			result, err := s.GenerateReportOne(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateReportOne() error = %v, wantErr %v", err, tt.wantErr)
			}

			t.Logf("RESULT: %v", result)
		})
	}
}

func initConfig() error {
	viper.AddConfigPath("/Users/pchen/GolandProjects/uir_draft/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
