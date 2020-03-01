package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/application"
)

type TaskController struct {
	Ta application.TaskApplication
}

func (tc *TaskController) Handler2(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World from TaskGo.\n")

	taskList := tc.Ta.All()
	buf, err := json.Marshal(taskList)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)

	// return c.String(http.StatusOK, string(buf))
}

/*
TESTTESTTESTTESTTESTTESTTEST
*/
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.\n")
}

/*
func (h *TodoHandler) Save() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := new(model.Todo)

		err := c.Bind(t)

		if err != nil {
			log.Fatal(err)
		}

		err = h.UC.Save(t)

		if err != nil {
			log.Fatal(err)
		}

		return c.String(http.StatusOK, "")
	}
}
*/
