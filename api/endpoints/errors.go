package endpoints

import (
  "../response"
  "net/http"
)

func Error404 (w http.ResponseWriter, r *http.Request) {
  response.JSON(w, 200, response.NotFoundError)
}

func Error405 (w http.ResponseWriter, r *http.Request) {
  response.JSON(w, 200, response.NotAllowedError)
}
