package repository_person

const createPersonQuery string = `
	INSERT INTO person (first_name,last_name,email) VALUES (:first_name,:last_name,:email)
`

const updatePersonQuery string = `
	UPDATE person 
	SET 
	    first_name=:first_name,
	    last_name=:last_name	    
	WHERE email=:email
`

const deletePersonByEmailQuery string = `
	DELETE FROM person WHERE email=:email
`

const getByLastNameQuery = `
	SELECT * FROM person
	WHERE last_name=:last_name
`
