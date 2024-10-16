package seed

import (
	"context"
	"fmt"

	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
)

var dogBreeds = []string{
	"Affenpinscher",
	"Afghan Hound",
	"Airedale Terrier",
	"Alaskan Malamute",
	"American Bully",
	"American Eskimo Dog",
	"American Pit Bull Terrier",
	"American Staffordshire Terrier",
	"Anatolian Shepherd Dog",
	"Australian Cattle Dog",
	"Australian Shepherd",
	"Basenji",
	"Basset Hound",
	"Beagle",
	"Bedlington Terrier",
	"Belgian Shepherd",
	"Bernese Mountain Dog",
	"Bichon Frisé",
	"Black Russian Terrier",
	"Border Collie",
	"Boston Terrier",
	"Boxer",
	"Brittany Spaniel",
	"Bull Terrier",
	"Bulldog",
	"Bullmastiff",
	"Cairn Terrier",
	"Cavalier King Charles Spaniel",
	"Chihuahua",
	"Dachshund",
	"Dobermann",
	"English Cocker Spaniel",
	"French Bulldog",
	"German Shepherd",
	"Goldendoodle",
	"Great Dane",
	"Havanese",
	"Maltese dog",
	"Maltipoo",
	"Newfoundland dog",
	"Pembroke Welsh Corgi",
	"Pomeranian",
	"Poodle",
	"Rottweiler",
	"Samoyed",
	"Sarabi dog",
	"Shiba Inu",
	"Shih Tzu",
	"Siberian Husky",
	"St. Bernard",
	"Yorkshire Terrier",
}

var catBreeds = []string{
	"Abyssinian",
	"American Bobtail",
	"American Curl",
	"American Shorthair",
	"American Wirehair",
	"Australian Mist",
	"Balinese cat",
	"Birman",
	"Bombay cat",
	"British Longhair",
	"British Shorthair",
	"Burmese cat",
	"Burmilla",
	"Chartreux",
	"Chausie",
	"Colorpoint Shorthair",
	"Cornish Rex",
	"Devon Rex",
	"Egyptian Mau",
	"Exotic Shorthair",
	"Havana Brown",
	"Himalayan cat",
	"Japanese Bobtail",
	"Javanese cat",
	"Khao Manee",
	"Korat",
	"LaPerm",
	"Lykoi",
	"Maine Coon",
	"Manx Cat",
	"Munchkin cat",
	"Norwegian Forest cat",
	"Ocicat",
	"Oriental Shorthair",
	"Persian cat",
	"Ragamuffin",
	"Ragdoll",
	"Russian Blue",
	"Scottish Fold",
	"Selkirk Rex",
	"Siamese cat",
	"Siberian cat",
	"Singapura cat",
	"Snowshoe cat",
	"Sokoke",
	"Somali cat",
	"Sphynx cat",
	"Tonkinese cat",
	"Toyger",
	"Turkish Angora",
	"Turkish Van",
}

func CreateDogs(
	ctx context.Context,
	dbClient *ent.Client,
) error {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return err
	}

	for _, dogBreed := range dogBreeds {
		_, err := tx.Pet.Create().
			SetSpecie(pet.SpecieDOG).
			SetBreed(dogBreed).
			Save(ctx)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf(
					"create dogs failed: %v, unable to rollback: %v",
					err,
					rollbackErr,
				)
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func CreateCats(
	ctx context.Context,
	dbClient *ent.Client,
) error {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return err
	}

	for _, catBreed := range catBreeds {
		_, err := tx.Pet.Create().
			SetSpecie(pet.SpecieCAT).
			SetBreed(catBreed).
			Save(ctx)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf(
					"create dogs failed: %v, unable to rollback: %v",
					err,
					rollbackErr,
				)
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
