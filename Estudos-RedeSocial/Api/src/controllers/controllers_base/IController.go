package controllers_base

type IController interface {
	IControllerCriar
	IControllerListar
	IControllerListarPorId
	IControllerAtualizar
	IControllerExcluir
}
