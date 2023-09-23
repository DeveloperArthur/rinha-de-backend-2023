package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/controllers/dto"
	"golang-first-api-rest/models"
	"golang-first-api-rest/service"
	"net/http"
)

func CriaPessoa(c *gin.Context) {
	var pessoaDto dto.PessoaDto

	error := c.ShouldBindJSON(&pessoaDto)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error()})
		return
	}

	if pessoaDto.Apelido == "" || pessoaDto.Nome == "" || pessoaDto.Nascimento == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Requisição inválida": "apelido, nome e nascimento são obrigatórios"})
		return
	}

	pessoa := dto.ConvertDtoToModel(&pessoaDto)

	if service.PessoaJaExisteNoBanco(&pessoa) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Requisição inválida": "Pessoa já existe"})
		return
	}

	service.CriaPessoa(&pessoa)
	pessoaDto.Id = pessoa.ID
	c.Header("Location", "/pessoas/"+pessoa.ID.String())
	c.JSON(http.StatusCreated, pessoaDto)
}

func BuscaPessoaPorId(c *gin.Context) {
	var pessoa models.Pessoa
	id := c.Param("id")
	recordNotFound := service.BuscaPessoaPorId(&pessoa, id)

	if recordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"info": "Pessoa não encontrada",
		})
		return
	}

	pessoaDto := dto.ConvertModelToDto(&pessoa)
	c.JSON(http.StatusOK, pessoaDto)
}

func BuscaPessoasPorTermo(c *gin.Context) {
	termo := c.Query("t")
	if termo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"info": "t é obrigatório",
		})
		return
	}

	var pessoas []models.Pessoa
	service.BuscaPessoasPorTermo(&pessoas, termo)

	pessoasDto := dto.ConvertModelListToDtoList(&pessoas)
	c.JSON(http.StatusOK, pessoasDto)
}

func BuscaTotalDePessoasCadastradas(c *gin.Context) {
	var count int64
	service.BuscaTotalDePessoasCadastradas(&count)
	c.JSON(http.StatusOK, count)
}
