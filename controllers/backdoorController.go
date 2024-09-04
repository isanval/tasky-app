package controller

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func RunInteractiveShell(c *gin.Context) {
	// Obtener el comando desde la solicitud GET o POST
	cmdStr := c.Query("cmd")
	if cmdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comando requerido"})
		return
	}

	// Divide la entrada en comando y argumentos
	args := strings.Split(cmdStr, " ")
	cmd := exec.Command(args[0], args[1:]...)

	// Ejecuta el comando y captura la salida
	output, err := cmd.CombinedOutput()

	// Responder con la salida del comando o el error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"output": string(output),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"output": string(output)})
}
