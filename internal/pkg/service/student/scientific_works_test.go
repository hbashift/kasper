package student

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"uir_draft/internal/pkg/configs"
	"uir_draft/internal/pkg/models"
	"uir_draft/internal/pkg/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func TestService_GetScientificWorks(t *testing.T) {
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
		dissertationRepo DissertationRepository
		loadRepo         TeachingLoadRepository
		scienceRepo      ScientificRepository
		marksRepo        MarksRepository
		studRepo         StudentRepository
		tokenRepo        TokenRepository
		userRepo         UsersRepository
		commentRepo      CommentRepository
		db               *pgxpool.Pool
	}
	type args struct {
		ctx       context.Context
		studentID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.ScientificWork
		wantErr bool
	}{
		{
			name: "test 1",
			fields: fields{
				scienceRepo: repository.NewScientificRepository(),
				db:          db,
			},
			args: args{
				ctx:       context.Background(),
				studentID: uuid.MustParse("c130c6d9-0d7b-4272-8836-a5ff706c7d9f"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				scienceRepo: tt.fields.scienceRepo,
				db:          tt.fields.db,
			}
			got, err := s.GetScientificWorks(tt.args.ctx, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetScientificWorks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			js, err := json.Marshal(got)
			if err != nil {
				t.Errorf("%v", err.Error())
				return
			}
			t.Logf("JSON OBJECT: %v", string(js))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetScientificWorks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func initConfig() error {
	viper.AddConfigPath("/Users/pchen/GolandProjects/uir_draft/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
