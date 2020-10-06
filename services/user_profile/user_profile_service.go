package user_profile

import "DulceDayServer/database/models"

type Service interface {
	GetProfileByUsername(username string) *models.UserProfile
	CreateNewProfile(username string) *models.UserProfile
}

type ServiceImpl struct {
	store UserProfileStore
}

func NewUserProfileServiceImpl(store UserProfileStore) *ServiceImpl {
	return &ServiceImpl{store: store}
}

func (s ServiceImpl) GetProfileByUsername(username string) *models.UserProfile {
	return s.store.getProfileByUsername(username)
}

func (s ServiceImpl) CreateNewProfile(username string) *models.UserProfile {
	return s.store.createNewProfile(username)
}