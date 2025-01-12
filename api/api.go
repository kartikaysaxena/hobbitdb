package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kartikayaxena/hobbitdb/hobbit"
	"github.com/labstack/echo/v4"
)

type Server struct {
	db *hobbit.Hobbit
}

func NewServer(db *hobbit.Hobbit) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) HandlePostInsert(c echo.Context) error {
	var (
		collname = c.Param("collname")
	)
	var data hobbit.Map
	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		return err
	}
	id, err := s.db.Coll(collname).Insert(data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, hobbit.Map{"id": id})
}

func (s *Server) HandleGetQuery(c echo.Context) error {
	var (
		collname  = c.Param("collname")
		filterMap = NewFilterMap()
	)
	for k, v := range c.QueryParams() {
		filterParts := strings.Split(k, ".")
		if len(filterParts) != 2 {
			return fmt.Errorf("mallformed query")
		}
		if len(v) == 0 {
			return fmt.Errorf("mallformed query")
		}
		if v[0] == "" {
			return fmt.Errorf("mallformed query")
		}
		var (
			filterType  = filterParts[0]
			filterKey   = filterParts[1]
			filterValue = v[0]
		)
		filterMap.Add(filterType, filterKey, filterValue)
	}
	records, err := s.db.Coll(collname).
		Eq(filterMap.Get(hobbit.FilterTypeEQ)).
		Find()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, records)
}
