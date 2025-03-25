// should to do
// gorm things like

// types for each tables
// for each type make get things
// get by id or part of name
// GORM things work only with gorm things, and u should use Raw sql requests

//  actor
//  actor_info
//  address
//  category
//  city
//  country
//  customer
//  customer_list
//  film
//  film_actor
//  film_category
//  film_list
//  language
//  nicer_but_slower_film_list
//  payment
//  rental
//  sales_by_film_category
//  sales_by_store
//  staff
//  staff_list
//  store

package main

import (
	"fmt"

	"github.com/TheEphir/LearningProject1/API/models"
)

func main() {
	fmt.Println(models.GetActors())
}
