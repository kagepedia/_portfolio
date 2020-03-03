package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/application"
)

type TaskController struct {
	Ta     application.TaskApplication
	Logger application.Logger
}

func (tc *TaskController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World from TaskGo.\n")
	// tc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	taskList := tc.Ta.All()
	res, err := json.Marshal(taskList)

	if err != nil {
		fmt.Println("n = nil, Invalid Error")
		// log.Fatal("Error: %s", err)
		// tc.Logger.LogError("%s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

/*
TESTTESTTESTTESTTESTTESTTEST
*/
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go11111.\n")
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
