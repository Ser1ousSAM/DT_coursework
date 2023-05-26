package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"time"
)

type Wine_sort struct {
	Fixed_acidity        float64
	Volatile_acidity     float64
	Citric_acid          float64
	Residual_sugar       float64
	Chlorides            float64
	Free_sulfur_dioxide  float64
	Total_sulfur_dioxide float64
	Density              float64
	PH                   float64
	Sulphates            float64
	Alcohol              float64
	Quality              float64
}

type Distance struct {
	Dist float64
	Wine Wine_sort
}

func GetWineList(s string) []Wine_sort {
	path := "wine_dataset/processed_data/" + s
	data := parseCsv(path)
	var wineList []Wine_sort
	for i := 1; i < len(data); i++ {
		var wine Wine_sort
		for j := 0; j < len(data[0]); j++ {
			value, err := strconv.ParseFloat(data[i][j], 64)
			if err != nil {
				fmt.Println("Error during conversion")
			}
			switch j {
			case 0:
				wine.Fixed_acidity = value
			case 1:
				wine.Volatile_acidity = value
			case 2:
				wine.Citric_acid = value
			case 3:
				wine.Residual_sugar = value
			case 4:
				wine.Chlorides = value
			case 5:
				wine.Free_sulfur_dioxide = value
			case 6:
				wine.Total_sulfur_dioxide = value
			case 7:
				wine.Density = value
			case 8:
				wine.PH = value
			case 9:
				wine.Sulphates = value
			case 10:
				wine.Alcohol = value
			case 11:
				wine.Quality = value
			default:
				fmt.Printf("out of range \n", j)
				return nil
			}
		}
		wineList = append(wineList, wine)
	}
	return wineList
}

func GetWeights(path string) Wine_sort {
	data := parseCsv(path)
	var weights Wine_sort
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data[0]); j++ {
			value, err := strconv.ParseFloat(data[i][j], 64)
			if value > 0 {
				value *= 2
			} else {
				value *= -1
			}
			if err != nil {
				fmt.Println("Error during conversion")
			}
			switch i {
			case 0:
				weights.Alcohol = value
			case 1:
				weights.Sulphates = value
			case 2:
				weights.Citric_acid = value
			case 3:
				weights.Fixed_acidity = value
			case 4:
				weights.Residual_sugar = value
			case 5:
				weights.Free_sulfur_dioxide = value
			case 6:
				weights.PH = value
			case 7:
				weights.Density = value
			case 8:
				weights.Total_sulfur_dioxide = value
			case 9:
				weights.Chlorides = value
			case 10:
				weights.Volatile_acidity = value
			default:
				fmt.Printf("out of range \n", j)
			}
		}
	}
	return weights
}

func getRandomSort() Wine_sort {
	var rand_wine Wine_sort
	path := "wine_dataset/processed_data/proc_exam_materials.csv"
	data := parseCsv(path)
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := len(data)
	i := rand.Intn(max-min+1) + min

	for j := 0; j < len(data[0]); j++ {
		value, err := strconv.ParseFloat(data[i][j], 64)
		if err != nil {
			fmt.Println("Error during conversion")
		}
		switch j {
		case 0:
			rand_wine.Fixed_acidity = value
		case 1:
			rand_wine.Volatile_acidity = value
		case 2:
			rand_wine.Citric_acid = value
		case 3:
			rand_wine.Residual_sugar = value
		case 4:
			rand_wine.Chlorides = value
		case 5:
			rand_wine.Free_sulfur_dioxide = value
		case 6:
			rand_wine.Total_sulfur_dioxide = value
		case 7:
			rand_wine.Density = value
		case 8:
			rand_wine.PH = value
		case 9:
			rand_wine.Sulphates = value
		case 10:
			rand_wine.Alcohol = value
		case 11:
			rand_wine.Quality = value
		default:
			fmt.Printf("out of range \n", j)
		}
	}
	return rand_wine
}

func GetISort(sort_number int) Wine_sort {
	var wine Wine_sort
	path := "wine_dataset/processed_data/proc_exam_materials.csv"
	data := parseCsv(path)
	i := sort_number

	for j := 0; j < len(data[0]); j++ {
		value, err := strconv.ParseFloat(data[i][j], 64)
		if err != nil {
			fmt.Println("Error during conversion")
		}
		switch j {
		case 0:
			wine.Fixed_acidity = value
		case 1:
			wine.Volatile_acidity = value
		case 2:
			wine.Citric_acid = value
		case 3:
			wine.Residual_sugar = value
		case 4:
			wine.Chlorides = value
		case 5:
			wine.Free_sulfur_dioxide = value
		case 6:
			wine.Total_sulfur_dioxide = value
		case 7:
			wine.Density = value
		case 8:
			wine.PH = value
		case 9:
			wine.Sulphates = value
		case 10:
			wine.Alcohol = value
		case 11:
			wine.Quality = value
		default:
			fmt.Printf("out of range \n", j)
		}
	}
	return wine
}

func AllDistances(exam_wine, weights Wine_sort, train_wine_list []Wine_sort) []Distance {
	var weight_distances []Distance
	val_exam_wine := reflect.ValueOf(exam_wine)
	val_weights := reflect.ValueOf(weights)
	for _, wine := range train_wine_list {
		var sum_f float64
		val_train_wine := reflect.ValueOf(wine)
		for i := 0; i < val_exam_wine.NumField()-1; i++ {
			field_exam_wine := val_exam_wine.Field(i).Float()
			field_weights_wine := val_weights.Field(i).Float()
			field_train_wine := val_train_wine.Field(i).Float()
			sum_f += field_weights_wine * (field_train_wine - field_exam_wine) * (field_train_wine - field_exam_wine)
		}
		var d Distance
		d.Dist = math.Sqrt(sum_f)
		d.Wine = wine
		weight_distances = append(weight_distances, d)
	}
	sort.Slice(weight_distances, func(i, j int) bool {
		return weight_distances[i].Dist < weight_distances[j].Dist
	})
	return weight_distances
}

// k%6 = 1
func KNNClassify(k int, distances []Distance) float64 {
	var class float64
	var list_classes [6]int
	for i, dist := range distances {
		if k == i {
			break
		}
		list_classes[int(dist.Wine.Quality)-1] += 1
	}
	var max_i int
	var max_v int
	for i, val := range list_classes {
		if val > max_v {
			max_v = val
			max_i = i
		}
	}
	class = float64(max_i + 1)
	return class
}

func GetProcentQuality(wineTrainList []Wine_sort, weights Wine_sort) float64 {
	wineExamList := GetWineList("proc_exam_materials.csv")
	var countExam = float64(len(wineExamList))
	var countMatches float64
	for _, wine := range wineExamList {
		distances := AllDistances(wine, weights, wineTrainList)
		class := KNNClassify(7, distances)
		if class == wine.Quality {
			countMatches++
		}
	}
	return countMatches / countExam * 100
}

func GetStringExamNumbers(wineTrainList []Wine_sort) []string {
	wineExamList := GetWineList("proc_exam_materials.csv")
	var nums []string
	for i, _ := range wineExamList {
		nums = append(nums, string(i+1))
	}
	return nums
}

func GetProcentCategory(wineTrainList []Wine_sort, weights Wine_sort) float64 {
	wineExamList := GetWineList("proc_exam_materials.csv")
	var countExam = float64(len(wineExamList))
	var countMatches float64
	for _, wine := range wineExamList {
		var examQuality string
		switch wine.Quality {
		case 1, 2:
			examQuality = "bad"
		case 3, 4:
			examQuality = "normal"
		case 5, 6:
			examQuality = "good"
		default:
			fmt.Printf("out of range \n", examQuality)
		}
		distances := AllDistances(wine, weights, wineTrainList)
		class := KNNClassify(7, distances)
		var predictedQuality string
		switch class {
		case 1, 2:
			predictedQuality = "bad"
		case 3, 4:
			predictedQuality = "normal"
		case 5, 6:
			predictedQuality = "good"
		default:
			fmt.Printf("out of range \n", predictedQuality)
		}
		if examQuality == predictedQuality {
			countMatches++
		}
	}
	return countMatches / countExam * 100
}

func main() {
	wineTrainList := GetWineList("proc_train_materials (copy).csv")
	weights := GetWeights("wine_dataset/weights.csv")
	procent := fmt.Sprintf("%.2f", GetProcentCategory(wineTrainList, weights))
	var exam_material = 2
	fmt.Println(procent)
	wine := GetISort(exam_material)
	var examQuality string
	switch wine.Quality {
	case 1, 2:
		examQuality = "bad"
	case 3, 4:
		examQuality = "normal"
	case 5, 6:
		examQuality = "good"
	default:
		fmt.Printf("out of range \n", examQuality)
	}
	distances := AllDistances(wine, weights, wineTrainList)
	class := KNNClassify(7, distances)
	var predictedQuality string
	switch class {
	case 1, 2:
		predictedQuality = "bad"
	case 3, 4:
		predictedQuality = "normal"
	case 5, 6:
		predictedQuality = "good"
	default:
		fmt.Printf("out of range \n", predictedQuality)
	}
	fmt.Println("Exam material:", exam_material)
	fmt.Println("Predicted category:", predictedQuality)
	fmt.Println("Class", class)
	fmt.Println("Real category:", examQuality)
	fmt.Println("Real class", wine.Quality)
}
