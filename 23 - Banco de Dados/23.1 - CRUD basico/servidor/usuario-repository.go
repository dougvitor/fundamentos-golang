package servidor

import (
	"crud/banco"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint64 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario processa a criação de um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if handleError(w, erro, "Falha ao ler o request body!", http.StatusBadRequest) {
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(body, &usuario); handleError(w, erro, "Erro ao preencher struct de usuário!", http.StatusBadRequest) {
		return
	}

	db, erro := banco.Conectar()
	if handleError(w, erro, "Erro ao conectar no banco de dados!", http.StatusInternalServerError) {
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios(nome, email) values (?, ?)")
	if handleError(w, erro, "Erro ao criar o statement!", http.StatusInternalServerError) {
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if handleError(w, erro, "Erro ao executar o statement!", http.StatusInternalServerError) {
		return
	}

	id, erro := insercao.LastInsertId()
	if handleError(w, erro, "Erro ao obter o ID inserido!", http.StatusInternalServerError) {
		return
	}

	writeResponse(w, http.StatusCreated, fmt.Sprintf("Usuário inserido com sucesso! ID: %d", id))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if handleError(w, erro, "Erro ao conectar no banco de dados!", http.StatusInternalServerError) {
		return
	}
	defer db.Close()

	registros, erro := db.Query("SELECT * FROM usuarios")
	if handleError(w, erro, "Erro ao buscar usuarios", http.StatusInternalServerError) {
		return
	}
	defer registros.Close()

	var usuarios []usuario
	for registros.Next() {
		var usuario usuario

		if erro = registros.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); handleError(w, erro, "Erro ao scanear o usuário", http.StatusInternalServerError) {
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(usuarios); handleError(w, erro, "Erro ao parsear os usuários para JSON", http.StatusInternalServerError) {
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if handleError(w, erro, "Erro ao converter o parâmetro para um Long", http.StatusInternalServerError) {
		return
	}

	db, erro := banco.Conectar()
	if handleError(w, erro, "Erro ao conectar no banco de dados!", http.StatusInternalServerError) {
		return
	}
	defer db.Close()

	registro, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", id)
	if handleError(w, erro, "Erro ao buscar o usuário!", http.StatusInternalServerError) {
		return
	}
	defer registro.Close()

	var usuario usuario
	if !registro.Next() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if erro = registro.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); handleError(w, erro, "Erro ao scanear o usuário!", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(usuario); handleError(w, erro, "Erro ao parsear o usuário para JSON!", http.StatusInternalServerError) {
		return
	}

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if handleError(w, erro, "Erro ao converter parâmetro para Long!", http.StatusInternalServerError) {
		return
	}

	requestBody, erro := io.ReadAll(r.Body)
	if handleError(w, erro, "Erro ao ler corpo da requisição", http.StatusInternalServerError) {
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(requestBody, &usuario); handleError(w, erro, "Erro ao converter o usuário para struct!", http.StatusInternalServerError) {
		return
	}

	db, erro := banco.Conectar()
	if handleError(w, erro, "Erro ao conectar no banco de dados", http.StatusInternalServerError) {
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if handleError(w, erro, "Erro ao criar o statemet!", http.StatusInternalServerError) {
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, id); handleError(w, erro, "Erro ao atualizar o usuário!", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func RemoverUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if handleError(w, erro, "Erro ao converter parâmetro para Long!", http.StatusInternalServerError) {
		return
	}

	db, erro := banco.Conectar()
	if handleError(w, erro, "Erro ao conectar no banco de dados", http.StatusInternalServerError) {
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if handleError(w, erro, "Erro ao criar o statemet!", http.StatusInternalServerError) {
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); handleError(w, erro, "Erro ao deletar o usuário!", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Função auxiliar para tratar erros e responder ao cliente
func handleError(w http.ResponseWriter, erro error, message string, status int) bool {
	if erro != nil {
		writeResponse(w, status, message)
		return true
	}
	return false
}

// Conectar cria e retorna uma nova conexão ao banco de dados.
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", "usuario:senha@/nome_do_banco")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Função auxiliar para escrever respostas HTTP
func writeResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
