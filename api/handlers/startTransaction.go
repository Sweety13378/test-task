package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task/pkg/database"
)

func StartTransaction(c *gin.Context) {
	// Get client id
	id := c.Param("client_id")

	// Get cost
	cost := c.Param("cost")

	// Start tx
	q := `BEGIN 
		SET TRANSACTION ISOLATED LEVEL SERIALIZABLE 
			UPDATE clients_schema.CLIENTS
			SET balance = client.balance - ?
			FROM (
				SELECT balance FROM clients_schema.COLUMNS
			) as client
			WHERE id = ?
	COMMIT`
	_, err := database.Conn.Exec(c.Request.Context(), q, cost, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sql command error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "cost": cost})
}
