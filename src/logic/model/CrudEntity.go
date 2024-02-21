package model

import persistenceentites "mealwhile/data/persistenceentities"

type CrudEntity interface {
	GetId() string
	SetId(id string)

	Empty() CrudEntity
	ToPersistenceEntity() persistenceentites.CrudPersistenceEntity
	String() string
	EntityName() string
	FromArguments(map[string]string) CrudEntity
}
