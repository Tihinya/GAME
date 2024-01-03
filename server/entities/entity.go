package entities

type Entity struct {
	Id int
}

type EntityManager struct {
	entities []*Entity
	Id       int
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: make([]*Entity, 0),
		Id:       1,
	}
}

func (em *EntityManager) CreateEntity() *Entity {
	entity := &Entity{Id: em.Id}
	em.entities = append(em.entities, entity)
	em.Id++
	return entity
}
