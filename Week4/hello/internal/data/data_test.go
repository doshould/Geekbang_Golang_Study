package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	conf2 "hello/internal/conf"
	"hello/internal/data/ent"
	"hello/internal/data/ent/movie"
	"os"
	"testing"
	"time"
)

func TestNewData(t *testing.T) {
	conf := &conf2.Data{
		Database: &conf2.Data_Database{
			Driver: "mysql",
			Source: "root:using@tcp(127.0.0.1:3306)/yanl",
		},
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	defer time.Sleep(time.Second)
	data, _, err := NewData(conf, logger)
	log.Info("data = ", data)
	log.Info("err = ", err)
	if err != nil {
		log.Info("err = ", err)
		return
	}
	movies, err := data.db.Movie.Query().Where(movie.ID(1001)).First(context.Background())
	log.Info("err = ", err)
	log.Info("movies = ", movies)
	if err != nil {
		log.Info("err = ", err)
		return
	}
	log.Info("m = ", movies)
}

func Test_Open(t *testing.T) {
	client, err := ent.Open("mysql", "root:using@tcp(127.0.0.1:3306)/yanl?parseTime=True")
	if err != nil {
		log.Info("err = ", err)
		return
	}
	//m, err := client.Movie.Create().SetID(1001).SetTitle("test").SetGeners("unkonw").Save(context.Background())
	m, err := client.Movie.Query().Where(movie.ID(1001)).First(context.Background())
	log.Info("m = ", m)
	log.Info("err = ", err) // 创建并返回
	time.Sleep(time.Second)
}
