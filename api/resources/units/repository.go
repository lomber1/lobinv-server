package units

import (
	"lobinv-server/database"
	"log"
)

func GetAllUnits() []Unit {
	query, err := database.DB.Query("SELECT id, singular_name, created_at from units;")
	if err != nil {
		log.Println(err.Error())
		return []Unit{}
	}

	defer query.Close()

	results := []Unit{}

	for query.Next() {
		var unit Unit

		err = query.Scan(&unit.ID, &unit.SingularName, &unit.CreatedAt)

		if err != nil {
			log.Fatal(err.Error())
		}

		results = append(results, unit)
	}

	log.Println(results)

	return results
}

func CreateUnit(body UnitDTO) (*Unit, error) {
	result, err := database.DB.Exec("INSERT INTO units (singular_name) VALUES (?)", body.SingularName)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	var unit Unit

	err = database.DB.QueryRow("SELECT id, singular_name, created_at from units WHERE id = ?", id).Scan(&unit.ID, &unit.SingularName, &unit.CreatedAt)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &unit, nil
}

func UpdateUnit(id int, body UnitDTO) (*Unit, error) {
	_, err := database.DB.Exec("UPDATE units SET singular_name = ? WHERE id = ?", body.SingularName, id)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	var unit Unit

	err = database.DB.QueryRow("SELECT id, singular_name, created_at from units WHERE id = ?", id).Scan(&unit.ID, &unit.SingularName, &unit.CreatedAt)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &unit, nil
}

func RemoveUnit(id int) error {
	_, err := database.DB.Exec("DELETE FROM units WHERE id = ?", id)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
