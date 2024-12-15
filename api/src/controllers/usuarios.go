package controllers

import "net/http"

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario!"))
}

// BuscarUsuario busca todos os usuarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuarios!"))
}

// BuscarUsuario busca um usuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario!"))
}

// AtualizarUsuario altera as informações de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario!"))
}

// DeletarUsuario exclui as informações de um usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario!"))
}
