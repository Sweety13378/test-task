package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task/pkg/database"
)

func StartTransaction(c *gin.Context) {
	// Get client id
	id := c.Param("client_id")

	// Start transaction
	q := `BEGIN 
		SET TRANSACTION ISOLATED LEVEL SERIALIZABLE 
			SELECT (balance) 
			FROM clients_schema.COLUMNS
			WHERE id = ?
	COMMIT`
	_, err := database.Conn.Exec(c.Request.Context(), q, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sql command error"})
		return
	}

}
