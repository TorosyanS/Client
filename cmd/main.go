package main

import (
	"errors"
	"fmt"
	"os"

	"time"

	"test/internal/embedding/training"
	"test/internal/encapculation/bank"
	"test/internal/polymorphism/storage"
	"test/internal/polymorphism/storage/file_storage"
	"test/internal/polymorphism/storage/map_storage"
	"test/internal/polymorphism/storage/service_storage"
	"test/internal/polymorphism/storage/slice_storage"
)

//type WeatherResponse struct {
//	Coord struct {
//		Lon float64 `json:"lon"`
//		Lat float64 `json:"lat"`
//	} `json:"coord"`
//	Weather []struct {
//		Id          int    `json:"id"`
//		Main        string `json:"main"`
//		Description string `json:"description"`
//		Icon        string `json:"icon"`
//	} `json:"weather"`
//	Base string `json:"base"`
//	Main struct {
//		Temp      float64 `json:"temp"`
//		FeelsLike float64 `json:"feels_like"`
//		TempMin   float64 `json:"temp_min"`
//		TempMax   float64 `json:"temp_max"`
//		Pressure  int     `json:"pressure"`
//		Humidity  int     `json:"humidity"`
//	} `json:"main"`
//	Visibility int `json:"visibility"`
//	Wind       struct {
//		Speed float64 `json:"speed"`
//		Deg   int     `json:"deg"`
//	} `json:"wind"`
//	Clouds struct {
//		All int `json:"all"`
//	} `json:"clouds"`
//	Dt  int `json:"dt"`
//	Sys struct {
//		Type    int    `json:"type"`
//		Id      int    `json:"id"`
//		Country string `json:"country"`
//		Sunrise int64  `json:"sunrise"`
//		Sunset  int64  `json:"sunset"`
//	} `json:"sys"`
//	Timezone int    `json:"timezone"`
//	Id       int    `json:"id"`
//	Name     string `json:"name"`
//	Cod      int    `json:"cod"`
//}
//
//func main() {
//	url := "https://api.openweathermap.org/data/2.5/weather?q=Moscow&appid=337307c1c55cadf4514d845751c090ca"
//	method := "GET"
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, nil)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//
//	body, err := io.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	resp := WeatherResponse{}
//	err = json.Unmarshal(body, &resp)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	sunriseTime := time.Unix(resp.Sys.Sunrise, 0)
//	sunsetTime := time.Unix(resp.Sys.Sunset, 0)
//	fmt.Printf("Расвет %s в Москве начинается в %s\n", sunriseTime.Format("02.01.2006"), sunriseTime.Format(time.TimeOnly))
//	fmt.Printf("Закат %s в Москве начинается в %s\n", sunsetTime.Format("02.01.2006"), sunsetTime.Format(time.TimeOnly))
//}

//func main() {
//	// ПОЛИМОРФИЗМ
//	//polymorphism()
//}

func encapsulation() {
	transferBank := bank.NewBank()

	transferBank.CreateNewAccount("200")
	err := transferBank.TransferMoney("100", "200", 500)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	account1, err := transferBank.FindAccount("100")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	account2, err := transferBank.FindAccount("200")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Balance %s - %d\n", account1.Number(), account1.Balance())
	fmt.Printf("Balance %s - %d\n", account2.Number(), account2.Balance())
}

func embedding() {
	swimming := training.Swimming{
		Training: training.Training{
			TrainingType: "Плавание",
			Action:       2000,
			LenStep:      training.SwimmingLenStep,
			Duration:     90 * time.Minute,
			Weight:       85,
		},
		LengthPool: 50,
		CountPool:  5,
	}

	fmt.Println(training.ReadData(swimming))

	walking := training.Walking{
		Training: training.Training{
			TrainingType: "Ходьба",
			Action:       20000,
			LenStep:      training.LenStep,
			Duration:     3*time.Hour + 45*time.Minute,
			Weight:       85,
		},
		Height: 185,
	}

	fmt.Println(training.ReadData(walking))

	running := training.Running{
		Training: training.Training{
			TrainingType: "Бег",
			Action:       5000,
			LenStep:      training.LenStep,
			Duration:     30 * time.Minute,
			Weight:       85,
		},
	}

	fmt.Println(training.ReadData(running))
}

func polymorphism() {
	store, err := createStorage(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer store.Close()

	err = store.SavePair("hello1", "bye")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	value, err := store.GetValue("hello1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(value)

	value, err = store.GetValue("hello_unknown")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(value)
}

func createStorage(storageType string) (storage storage.Storage, err error) {
	switch storageType {
	case "map":
		storage = map_storage.NewStorage()
	case "slice":
		storage = slice_storage.NewStorage()
	case "file":
		storage, err = file_storage.NewStorage()
	case "service":
		storage = service_storage.NewStorage()
	default:
		err = errors.New("unknown storage type")
	}
	return
}
