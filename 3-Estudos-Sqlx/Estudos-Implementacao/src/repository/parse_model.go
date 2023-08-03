package repository

type ParserModel[TEntity interface{}, TModel interface{}] interface {
	ParseToModel(entity TEntity) TModel
	ParseToEntity(model TModel) TEntity
}
