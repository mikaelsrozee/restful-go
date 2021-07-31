package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var QueryCreateTables = `
    CREATE TABLE Questionnaires (
      id CHAR(36) PRIMARY KEY,
      name TEXT NOT NULL,
      visibility BOOLEAN NOT NULL,
      questions TEXT NOT NULL
    );


    CREATE TABLE Users (
      id CHAR(21) PRIMARY KEY
    );


    CREATE TABLE UsersQuestionnaires (
      user_id CHAR(21),
      questionnaire_id CHAR(36),
      FOREIGN KEY (user_id) REFERENCES Users(id),
      FOREIGN KEY (questionnaire_id) REFERENCES Questionnaires(id)
    );


    CREATE TABLE Responses (
        response TEXT NOT NULL,
        questionnaire_id CHAR(36),
        FOREIGN KEY (questionnaire_id) REFERENCES Questionnaires(id)
    );


    INSERT INTO Users (id) VALUES
    ('_______________PUBLIC');
`

func (a* App) createTables() {
    _, err := a.DB.Exec(QueryCreateTables)

    if err != nil {
        panic(err)
    }

    log.Println("DB tables created")
}

func (a* App) createDatabaseFile(dbName string) {
    log.Println("Creating DB file")

    dbFile := "./" + dbName + ".db"

    // Delete database file (for testing)
    os.Remove(dbFile)

    var err error
    a.DB, err = sql.Open("sqlite3", dbFile)

    if err != nil {
        panic(err)
    }

    log.Println("DB file created")
}

func (a* App) InitDatabase(dbName string) {
    log.Println("Initializing database")
    a.createDatabaseFile(dbName)
    a.createTables()

    defer a.DB.Close()
    log.Println("Database initialized")
}
