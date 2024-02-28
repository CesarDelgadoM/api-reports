package branch

import (
	"errors"

	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
)

type IService interface {
	CreateBranch(userId uint, name string, branch *Branch) error
	FindBranch(userId uint, name, manager string) (*Branch, error)
	FindBranchs(userId uint, name string) (*[]Branch, error)
	UpdateBranch(userId uint, name, manager string, branch *Branch) error
	DeleteBranch(userId uint, name, manager string) error
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) IService {
	return &Service{
		repo: repo,
	}
}

func (service *Service) CreateBranch(userId uint, name string, branch *Branch) error {
	err := service.repo.Create(userId, name, branch)
	if err != nil {
		switch {
		case errors.Is(err, httperrors.BranchNotFound):
			return httperrors.BranchNotFound

		default:
			zap.Logger.Error(httperrors.BranchCreateFailed, err)
			return httperrors.BranchCreateFailed
		}
	}

	zap.Logger.Info("Branch created: ", branch.Name)
	return nil
}

func (service *Service) FindBranch(userId uint, name, manager string) (*Branch, error) {
	branch, err := service.repo.Find(userId, name, manager)
	if err != nil {
		zap.Logger.Error(httperrors.ErrBranchNotFound, err)
		return nil, httperrors.BranchNotFound
	}

	return branch, nil
}

func (service *Service) FindBranchs(userId uint, name string) (*[]Branch, error) {
	branchs, err := service.repo.FindAll(userId, name)
	if err != nil {
		zap.Logger.Error(httperrors.ErrBranchNotFound, err)
		return nil, httperrors.BranchNotFound
	}

	if len(*branchs) == 0 {
		return nil, httperrors.BranchArrayEmpty
	}

	return branchs, nil
}

func (service *Service) UpdateBranch(userId uint, name, manager string, branch *Branch) error {
	err := service.repo.Update(userId, name, manager, branch)
	if err != nil {
		switch {
		case errors.Is(err, httperrors.BranchNotFound):
			zap.Logger.Error(httperrors.ErrBranchNotFound, err)
			return httperrors.BranchNotFound

		default:
			zap.Logger.Error(httperrors.ErrBranchUpdate, err)
			return httperrors.BranchUpdateFailed
		}
	}

	return nil
}

func (service *Service) DeleteBranch(userId uint, name, manager string) error {
	err := service.repo.Delete(userId, name, manager)
	if err != nil {
		switch {
		case errors.Is(err, httperrors.BranchNotFound):
			zap.Logger.Error(httperrors.ErrBranchNotFound, err)
			return httperrors.BranchNotFound

		default:
			zap.Logger.Error(httperrors.ErrBranchDelete)
			return httperrors.BranchDeleteFailed
		}
	}

	return nil
}
