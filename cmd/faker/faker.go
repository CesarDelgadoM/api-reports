package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/CesarDelgadoM/api-reports/config"
	"github.com/CesarDelgadoM/api-reports/internal/api/branch"
	"github.com/CesarDelgadoM/api-reports/internal/api/restaurant"
	"github.com/CesarDelgadoM/api-reports/pkg/database"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"github.com/go-faker/faker/v4"
)

// Paramaters to create a restaurant
const (
	userid            = 1
	name              = "LaMargarita"
	numberBranches    = 17
	numberOptionsMenu = 15
	numberEmployees   = 10
	minrange          = 1000000
	maxrange          = 1000000000
	maxage            = 60
	maxSales          = 100
	maxScore          = 5
)

// Create a restaurant in mongodb with fake data
func main() {
	// Config load
	loadcfg := config.LoadConfig("config-local.yml")

	cfg := config.ParseConfig(loadcfg)

	// Logger zap initialization
	zap.InitLogger(cfg)

	// Mongo connection
	mongoConn := database.ConnectMongoDB(cfg.Mongo)
	defer mongoConn.Disconnect()

	// Restaurant
	restaurantRepository := restaurant.NewRepository(mongoConn)
	restaurantService := restaurant.NewService(restaurantRepository)

	// Branch
	branchRepository := branch.NewRepository(mongoConn)
	branchService := branch.NewService(branchRepository)

	//McDonalds fake
	macdonalds := restaurant.Restaurant{
		UserId:      userid,
		Name:        name,
		Founder:     faker.Name(),
		Location:    "Chicago, Illinois",
		Country:     "EEUU",
		Fundation:   "15 de abril de 1955",
		Headquarter: "Chicago, Illinois",
		Branches:    []branch.Branch{},
	}

	zap.Logger.Info("Creating restaurant: ", macdonalds)
	restaurantService.CreateRestaurant(&macdonalds)

	// Fake branch creation
	for i := 0; i < numberBranches; i++ {
		// Menu
		menu := branch.Menu{
			EntreePlates: make([]string, numberOptionsMenu),
			MainCourse:   make([]string, numberOptionsMenu),
			Drinks:       make([]string, numberOptionsMenu),
			Desserts:     make([]string, numberOptionsMenu),
		}

		for m := 0; m < numberOptionsMenu; m++ {
			mstr := strconv.Itoa(m + 1)

			menu.EntreePlates[m] = "entree plates " + mstr
			menu.MainCourse[m] = "main course " + mstr
			menu.Drinks[m] = "drink " + mstr
			menu.Desserts[m] = "dessert " + mstr
		}

		// Financial
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		financial := branch.Financial{
			Sales:    uint32(r.Intn(maxrange-minrange) + minrange),
			Expenses: uint32(r.Intn(maxrange-minrange) + minrange),
		}

		// Employees
		employess := branch.Employees{
			Admins:       make([]branch.Employee, numberEmployees),
			Waiters:      make([]branch.Employee, numberEmployees),
			Chefs:        make([]branch.Employee, numberEmployees),
			TotalAdmins:  numberEmployees,
			TotalWaiters: numberEmployees,
			TotalChefs:   numberEmployees,
		}

		for e := 0; e < numberEmployees; e++ {

			employess.Admins[e] = branch.Employee{
				Name:  faker.Name(),
				Years: uint8(r.Intn(maxage)),
				Sales: uint8(r.Intn(maxSales)),
			}

			employess.Waiters[e] = branch.Employee{
				Name:  faker.Name(),
				Years: uint8(r.Intn(maxage)),
				Sales: uint8(r.Intn(maxSales)),
			}

			employess.Chefs[e] = branch.Employee{
				Name:  faker.Name(),
				Years: uint8(r.Intn(maxage)),
				Sales: uint8(r.Intn(maxSales)),
			}
		}

		// Branch
		branch := branch.Branch{
			Name:      faker.Name(),
			Manager:   faker.Name(),
			City:      faker.Name(),
			Address:   "cll/cra 000 # 000 - 000",
			Phone:     faker.Phonenumber(),
			Score:     uint8(r.Intn(maxScore)),
			Employees: employess,
			Financial: financial,
			Menu:      menu,
		}

		zap.Logger.Info("Creating branch: ", branch)
		if err := branchService.CreateBranch(userid, name, &branch); err != nil {
			zap.Logger.Fatal("error creating fake data: ", err)
		}
	}
}
