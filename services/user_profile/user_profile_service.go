package user_profile

import (
	"DulceDayServer/database/models"
)

type Service interface {
	CreateNewProfileByUserIdentifier(userIdentifier string) *models.UserProfile
	UpdateProfileByUserIdentifier(userIdentifier string, new *models.UserProfile)

	GetFullUserByUserIdentifier(userIdentifier string) *FullUser
	GetFullUserByUsername(username string) *FullUser
}

type ServiceImpl struct {
	store UserProfileStore
}

func NewUserProfileServiceImpl(store UserProfileStore) *ServiceImpl {
	return &ServiceImpl{store: store}
}

func (s ServiceImpl) CreateNewProfileByUserIdentifier(userIdentifier string) *models.UserProfile {
	return s.store.createNewProfile(userIdentifier)
}

func (s ServiceImpl) UpdateProfileByUserIdentifier(userIdentifier string, new *models.UserProfile) {
	orig := s.store.findUserProfileByUserIdentifier(userIdentifier)
	s.store.updateUserProfile(orig, new)
}

func (s ServiceImpl) GetFullUserByUserIdentifier(userIdentifier string) *FullUser {
	return s.store.findFullUserByUserIdentifier(userIdentifier)
}

func (s ServiceImpl) GetFullUserByUsername(username string) *FullUser {
	return s.store.findFullUserByUsername(username)
}
