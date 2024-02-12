package engine

import (
	"testing"
)

func TestEntityManager_CreateEntity(t *testing.T) {
	entityManager := NewEntityManager()

	entity1 := entityManager.CreateEntity("player")
	if entity1.Id != 1 {
		t.Errorf("Expected Id to be 1, got %d", entity1.Id)
	}

	entity2 := entityManager.CreateEntity("player")
	if entity2.Id != 2 {
		t.Errorf("Expected Id to be 2, got %d", entity2.Id)
	}

	entity3 := entityManager.CreateEntity("player")
	if entity3.Id != 3 {
		t.Errorf("Expected Id to be 3, got %d", entity3.Id)
	}

	if len(entityManager.entities) != 3 {
		t.Errorf("Expected length of entities slice to be 3, got %d", len(entityManager.entities))
	}
}
