package main

import (
	"3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql"
	"3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	domainPerson "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	repositoryPerson "3-Estudos-Sqlx/Estudos-Implementacao/src/repository/person"
	repositoryUnitOfWork "3-Estudos-Sqlx/Estudos-Implementacao/src/repository/unitofwork"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var schema = `
DROP TABLE IF EXISTS person;
DROP TABLE IF EXISTS place;

CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
);
`

func main() {
	connectionStringConfig := pkg_sql_connections_config.NewSqlConnectionConfig("localhost", 5432, "admin", "Password", "estudos_sqlx")
	sqlService := pkg_sql.NewSqlService(connectionStringConfig)
	err := sqlService.OpenConnection()
	if err != nil {
		log.Fatalln(err)
	}
	defer func(sqlService pkg_sql.SqlService) {
		err := sqlService.CloseConnection()
		if err != nil {
			log.Fatalln(err)
		}
	}(sqlService)

	db, err := sqlService.GetDb()
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)

	ctx := context.TODO()
	unitOfWork := repositoryUnitOfWork.NewUnitOfWork(sqlService)
	parserPerson := repositoryPerson.NewParserModelPerson()
	personRepository := repositoryPerson.NewRepositoryPerson(sqlService, parserPerson)
	_ = personRepository.Create(domainPerson.NewPerson("P1", "L1", "P1@P1.com"), &ctx, nil)
	_ = personRepository.Create(domainPerson.NewPerson("P2", "L2", "P2@P2.com"), &ctx, nil)
	_ = personRepository.Create(domainPerson.NewPerson("P3", "L3", "P3@P3.com"), &ctx, nil)
	_ = personRepository.Create(domainPerson.NewPerson("P3_1", "L3", "P3@P3.com"), &ctx, nil)
	_ = personRepository.Create(domainPerson.NewPerson("P3_2", "L3", "P3@P3.com"), &ctx, nil)
	err = personRepository.Create(domainPerson.NewPerson("P4", "L4", "P4@P4.com"), &ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = personRepository.Update(domainPerson.NewPerson("P1ALTERED", "L1ALTERED", "P1@P1.com"), &ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = personRepository.DeleteByEmail("P4@P4.com", &ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	name, err := personRepository.GetFirstOrDefaultByLastName("L3", &ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)

	names, err := personRepository.ListByLastName("L3", &ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(names)

	_, err = unitOfWork.Execute(func(tx *sqlx.Tx, ctx *context.Context) (interface{}, error) {
		_ = personRepository.Create(domainPerson.NewPerson("UnitOfWorkP1", "UnitOfWorkL1", "UnitOfWorkP1@P1.com"), ctx, tx)
		_ = personRepository.Create(domainPerson.NewPerson("UnitOfWorkP2", "UnitOfWorkL2", "UnitOfWorkP2@P2.com"), ctx, tx)
		_ = personRepository.Create(domainPerson.NewPerson("UnitOfWorkP3", "UnitOfWorkL3", "UnitOfWorkP3@P3.com"), ctx, tx)
		_ = personRepository.Create(domainPerson.NewPerson("UnitOfWorkP3_1", "UnitOfWorkL3", "UnitOfWorkP3_1@P3_1.com"), ctx, tx)
		_ = personRepository.Create(domainPerson.NewPerson("UnitOfWorkP3_2", "UnitOfWorkL3", "UnitOfWorkP3_2@P3_2.com"), ctx, tx)
		_ = personRepository.Create(domainPerson.NewPerson("UnitOfWorkP4", "UnitOfWorkL4", "UnitOfWorkP4@P4.com"), ctx, tx)
		if err != nil {
			log.Fatalln(err)
		}

		err = personRepository.Update(domainPerson.NewPerson("UnitOfWorkP1ALTERED", "UnitOfWorkL1Altered", "P1@P1.com"), ctx, tx)
		if err != nil {
			log.Fatalln(err)
		}

		err = personRepository.DeleteByEmail("UnitOfWorkP4@P4.com", ctx, tx)
		if err != nil {
			log.Fatalln(err)
		}

		name, err := personRepository.GetFirstOrDefaultByLastName("UnitOfWorkL3", ctx, tx)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(name)

		names, err := personRepository.ListByLastName("UnitOfWorkL3", ctx, tx)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(names)

		return nil, nil //errors.New("FAIO")
	}, &ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
