package trade

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler 处理交割单相关的 HTTP 请求。
type Handler struct {
	store *Store
}

// NewHandler 返回绑定 store 的 Handler。
func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

// Register 把交割单路由挂载到 gin 引擎。
func (h *Handler) Register(r *gin.Engine) {
	g := r.Group("/api/trades")
	g.POST("", h.create)
	g.GET("", h.list)
	g.GET("/stats", h.stats)
}

func (h *Handler) create(c *gin.Context) {
	var t Trade
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.store.Create(&t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, t)
}

func (h *Handler) list(c *gin.Context) {
	trades, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trades)
}

func (h *Handler) stats(c *gin.Context) {
	st, err := h.store.Stats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, st)
}
