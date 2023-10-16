package utils

import "github.com/theritikchoure/shorten/pkg/repository"

func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable
	var credentials []string

	switch role {
	case repository.AdminRoleName:
		// Admin credentials (all access)
		credentials = []string{
			repository.BookCreateCredential, repository.BookUpdateCredential, repository.BookDeleteCredential,
		}
	case repository.ModeratorRoleName:
		credentials = []string{
			repository.BookCreateCredential, repository.BookUpdateCredential,
		}
	case repository.UserRoleName:
		credentials = []string{
			repository.BookCreateCredential,
		}
	}

	return credentials, nil
}
