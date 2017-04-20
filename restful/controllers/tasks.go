package controllers

import "net/http"

var Tasks = new(tasksController)

type tasksController struct {
}

func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request) {}
func (tc *tasksController) Get(w http.ResponseWriter, r *http.Request)    {}
func (tc *tasksController) Show(w http.ResponseWriter, r *http.Request)   {}
func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request) {}
func (tc *tasksController) Delete(w http.ResponseWriter, r *http.Request) {}
