// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/curiosity"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/like"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/schema"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/view"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	curiosityFields := schema.Curiosity{}.Fields()
	_ = curiosityFields
	// curiosityDescTitle is the schema descriptor for title field.
	curiosityDescTitle := curiosityFields[1].Descriptor()
	// curiosity.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	curiosity.TitleValidator = curiosityDescTitle.Validators[0].(func(string) error)
	// curiosityDescContent is the schema descriptor for content field.
	curiosityDescContent := curiosityFields[2].Descriptor()
	// curiosity.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	curiosity.ContentValidator = curiosityDescContent.Validators[0].(func(string) error)
	// curiosityDescCreatedAt is the schema descriptor for createdAt field.
	curiosityDescCreatedAt := curiosityFields[3].Descriptor()
	// curiosity.DefaultCreatedAt holds the default value on creation for the createdAt field.
	curiosity.DefaultCreatedAt = curiosityDescCreatedAt.Default.(func() time.Time)
	// curiosityDescUpdatedAt is the schema descriptor for updatedAt field.
	curiosityDescUpdatedAt := curiosityFields[4].Descriptor()
	// curiosity.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	curiosity.DefaultUpdatedAt = curiosityDescUpdatedAt.Default.(func() time.Time)
	// curiosity.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	curiosity.UpdateDefaultUpdatedAt = curiosityDescUpdatedAt.UpdateDefault.(func() time.Time)
	// curiosityDescID is the schema descriptor for id field.
	curiosityDescID := curiosityFields[0].Descriptor()
	// curiosity.DefaultID holds the default value on creation for the id field.
	curiosity.DefaultID = curiosityDescID.Default.(func() uuid.UUID)
	likeFields := schema.Like{}.Fields()
	_ = likeFields
	// likeDescCreatedAt is the schema descriptor for createdAt field.
	likeDescCreatedAt := likeFields[1].Descriptor()
	// like.DefaultCreatedAt holds the default value on creation for the createdAt field.
	like.DefaultCreatedAt = likeDescCreatedAt.Default.(func() time.Time)
	// likeDescID is the schema descriptor for id field.
	likeDescID := likeFields[0].Descriptor()
	// like.DefaultID holds the default value on creation for the id field.
	like.DefaultID = likeDescID.Default.(func() uuid.UUID)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescBreed is the schema descriptor for breed field.
	petDescBreed := petFields[2].Descriptor()
	// pet.BreedValidator is a validator for the "breed" field. It is called by the builders before save.
	pet.BreedValidator = petDescBreed.Validators[0].(func(string) error)
	// petDescSearch is the schema descriptor for search field.
	petDescSearch := petFields[3].Descriptor()
	// pet.SearchValidator is a validator for the "search" field. It is called by the builders before save.
	pet.SearchValidator = petDescSearch.Validators[0].(func(string) error)
	// petDescCreatedAt is the schema descriptor for createdAt field.
	petDescCreatedAt := petFields[4].Descriptor()
	// pet.DefaultCreatedAt holds the default value on creation for the createdAt field.
	pet.DefaultCreatedAt = petDescCreatedAt.Default.(func() time.Time)
	// petDescUpdatedAt is the schema descriptor for updatedAt field.
	petDescUpdatedAt := petFields[5].Descriptor()
	// pet.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	pet.DefaultUpdatedAt = petDescUpdatedAt.Default.(func() time.Time)
	// pet.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	pet.UpdateDefaultUpdatedAt = petDescUpdatedAt.UpdateDefault.(func() time.Time)
	// petDescID is the schema descriptor for id field.
	petDescID := petFields[0].Descriptor()
	// pet.DefaultID holds the default value on creation for the id field.
	pet.DefaultID = petDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	viewFields := schema.View{}.Fields()
	_ = viewFields
	// viewDescCreatedAt is the schema descriptor for createdAt field.
	viewDescCreatedAt := viewFields[1].Descriptor()
	// view.DefaultCreatedAt holds the default value on creation for the createdAt field.
	view.DefaultCreatedAt = viewDescCreatedAt.Default.(func() time.Time)
	// viewDescID is the schema descriptor for id field.
	viewDescID := viewFields[0].Descriptor()
	// view.DefaultID holds the default value on creation for the id field.
	view.DefaultID = viewDescID.Default.(func() uuid.UUID)
}
