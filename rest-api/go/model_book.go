/*
 * Books service example
 *
 * Everything about books
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Book struct {

	Id string `json:"id"`

	AuthorId string `json:"authorId"`

	CategorieId string `json:"categorieId"`

	Title string `json:"title"`
}
