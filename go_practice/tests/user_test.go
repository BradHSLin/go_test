package tests

import (
	"go_practice/db"
	"go_practice/models"
	"log"
	"os"
	"testing"
)

func init() {
	e := os.Remove("api.db")
	if e != nil {
		log.Fatal(e)
	}

	db.InitDB()
}

func TestGetAllUsers(t *testing.T) {

	user1 := models.User{Name: "brad", Email: "brad@test.com", Height: 175, Weight: 65, Group: 1}
	user1.Save()

	user2 := models.User{Name: "brad2", Email: "brad2@test.com", Height: 1752, Weight: 652, Group: 12}
	user2.Save()

	allUsers, err := models.GetAllUsers()
	if err != nil {
		t.Errorf("TestGetAllUsers error:%s", err)
	}

	if len(allUsers) != 2 {
		t.Errorf("TestGetAllUsers len:%d", len(allUsers))
	}

	models.EmptyUser()
}

func TestGetUserByID(t *testing.T) {
	user := models.User{Name: "brad", Email: "brad@test.com", Height: 175, Weight: 65, Group: 1}
	user.Save()

	userByID, err := models.GetUserByID(user.ID)
	if err != nil {
		t.Errorf("GetUserByID error:%s", err)
	}

	if user.Name != userByID.Name {
		t.Errorf("GetUserByID Name error: %s != %s", user.Name, userByID.Name)
	}

	if user.Email != userByID.Email {
		t.Errorf("GetUserByID Email error: %s != %s", user.Email, userByID.Email)
	}

	if user.Height != userByID.Height {
		t.Errorf("GetUserByID Email error: %f != %f", user.Height, userByID.Height)
	}

	if user.Weight != userByID.Weight {
		t.Errorf("GetUserByID Email error: %f != %f", user.Weight, userByID.Weight)
	}

	if user.Group != userByID.Group {
		t.Errorf("GetUserByID Email error: %d != %d", user.Group, userByID.Group)
	}

	models.EmptyUser()
}

func TestUpdate(t *testing.T) {
	user := models.User{Name: "brad", Email: "brad@test.com", Height: 175, Weight: 65, Group: 1}
	user.Save()

	user.Name = "updatedBrad"
	user.Email = "updatedBrad@test.com"
	user.Height = 176
	user.Weight = 75
	user.Group = 2

	user.Update()

	userByID, err := models.GetUserByID(user.ID)
	if err != nil {
		t.Errorf("Update error:%s", err)
	}

	if user.Name != userByID.Name {
		t.Errorf("Update Name error: %s != %s", user.Name, userByID.Name)
	}

	if user.Email != userByID.Email {
		t.Errorf("Update Email error: %s != %s", user.Email, userByID.Email)
	}

	if user.Height != userByID.Height {
		t.Errorf("Update Email error: %f != %f", user.Height, userByID.Height)
	}

	if user.Weight != userByID.Weight {
		t.Errorf("Update Email error: %f != %f", user.Weight, userByID.Weight)
	}

	if user.Group != userByID.Group {
		t.Errorf("Update Email error: %d != %d", user.Group, userByID.Group)
	}

	models.EmptyUser()
}

func TestDelete(t *testing.T) {
	user := models.User{Name: "brad", Email: "brad@test.com", Height: 175, Weight: 65, Group: 1}
	err := user.Save()

	if err != nil {
		t.Errorf("TestDelete save error:%s", err)
	}

	user.Delete()

	userByID, err := models.GetUserByID(user.ID)
	if err == nil {
		t.Errorf("TestDelete Delete error")
		t.Error(userByID)
	}
}
