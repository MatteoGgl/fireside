/*
* main.go
* Copyright (C) <2021>  <Matteo Guglielmetti>
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as published
* by the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/matteoggl/fireside/internal/data"
	"github.com/petaki/inertia-go"
)

const version = "0.1.0"

type config struct {
	url  string
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config  config
	inertia *inertia.Inertia
	logger  *log.Logger
	models  data.Models
	session *sessions.Session
}

func main() {
	logger := log.New(os.Stdout, "fireside: ", log.Ldate|log.Ltime)

	flag.Parse()

	var cfg config

	initConfig(&cfg)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	rootTemplate := "./ui/app.tmpl"
	inertiaManager := inertia.New(cfg.url, rootTemplate, version)

	secret := os.Getenv("FIRESIDE_SECRET")
	if secret == "" {
		logger.Fatal("FIRESIDE_SECRET must not be empty")
	}
	session := sessions.New([]byte(secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := &application{
		config:  cfg,
		inertia: inertiaManager,
		logger:  logger,
		models:  data.NewModels(db),
		session: session,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func initConfig(cfg *config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("FIRESIDE_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	cfg.port = port

	cfg.url = os.Getenv("FIRESIDE_URL")
	if cfg.url == "" {
		cfg.url = fmt.Sprintf("http://localhost:%d", cfg.port)
	}

	cfg.env = os.Getenv("FIRESIDE_ENV")
	if cfg.env == "" {
		cfg.env = "development"
	}

	cfg.db.dsn = os.Getenv("FIRESIDE_DB_DSN")

	maxOpenConns, err := strconv.Atoi(os.Getenv("FIRESIDE_DB_MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatal(err)
	}
	cfg.db.maxOpenConns = maxOpenConns

	maxIdleConns, err := strconv.Atoi(os.Getenv("FIRESIDE_DB_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatal(err)
	}
	cfg.db.maxIdleConns = maxIdleConns

	cfg.db.maxIdleTime = os.Getenv("FIRESIDE_DB_MAX_IDLE_TIME")
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
