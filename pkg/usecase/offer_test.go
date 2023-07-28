package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/sreenathsvrm/voix-me/pkg/domain"
	"github.com/sreenathsvrm/voix-me/pkg/repository/mockrepo"
	mockinmemmory_db "github.com/sreenathsvrm/voix-me/pkg/service/mockservice"
	"github.com/stretchr/testify/assert"
)

func TestOfferUseCase_FindAll(t *testing.T) {
	// Create a new controller to manage the mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock instances of OfferRepo and InMemoryDB
	mockOfferRepo := mockrepo.NewMockOfferRepository(ctrl)
	mockInMemoryDB := mockinmemmory_db.NewMockInMemoryStorage(ctrl)

	// Create the offerUseCase instance with the mock dependencies
	offerUC := &offerUseCase{
		offerRepo:  mockOfferRepo,
		inMemoryDB: mockInMemoryDB,
	}

	testData := []struct {
		name           string
		input          string
		buildStub      func(offerRepo *mockrepo.MockOfferRepository, inMemoryDB *mockinmemmory_db.MockInMemoryStorage, countryName string)
		expectedOutput []domain.OfferCompany
		expectedError  error
	}{
		{
			name:  "is not found in inmemmory",
			input: "BR",
			buildStub: func(offerRepo *mockrepo.MockOfferRepository, inMemoryDB *mockinmemmory_db.MockInMemoryStorage, countryName string) {
				// We expect the FindAll method to be called once with the provided countryName
				inMemoryDB.EXPECT().GetOfferCompanyByCountryName(gomock.Any(), "BR").Return(nil, false)
				offerRepo.EXPECT().FindAll(gomock.Any(), countryName).Times(1).
					Return(nil, nil)
			},
			expectedOutput: nil,
			expectedError:  nil,
		},
		{
			name:  "successfully get the offers",
			input: "BR",
			buildStub: func(offerRepo *mockrepo.MockOfferRepository, inMemoryDB *mockinmemmory_db.MockInMemoryStorage, countryName string) {
				// We expect the FindAll method to be called once with the provided countryName
				inMemoryDB.EXPECT().GetOfferCompanyByCountryName(gomock.Any(), "BR").Return([]domain.OfferCompany{
					domain.OfferCompany{
                        OfferID: 3,
						Country: "BR",
						OfferHome: 2,
					},
				}, true)
			},
			expectedOutput: []domain.OfferCompany{
				domain.OfferCompany{
					OfferID: 3,
					Country: "BR",
					OfferHome: 2,
				},
			},
			expectedError:  nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			// Build the stub behavior for the offerRepo and inMemoryDB
			tt.buildStub(mockOfferRepo, mockInMemoryDB, tt.input)

			// Call the FindAll function
			offerCompanies, err := offerUC.FindAll(context.Background(), tt.input)

			// Check the expected output or error
			assert.Equal(t, tt.expectedOutput, offerCompanies)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
