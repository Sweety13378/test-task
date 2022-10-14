package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task/pkg/database"
)

func AddBalance(c *gin.Context) {
	// Get client id
	id := c.Param("id")

	// Start tx
	q := `BEGIN
		SET TRANSACTION ISOLATED LEVEL SERIALIZED
			UPDATE clients_schema.CLIENTS
			SET balance = balance + 1
			WHERE = id = ?
	COMMIT`
	_, err := database.Conn.Exec(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error adding balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
