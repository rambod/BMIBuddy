package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func calculateBMI(weightKg, heightM float64) (float64, string, string) {
	bmi := weightKg / (heightM * heightM)
	fmt.Println(bmi)
	var category, tip string

	switch {
	case bmi < 18.5:
		category = "Underweight"
		tip = "Consider increasing your calorie intake with nutrient-rich foods \n and consult with a healthcare provider for a plan to reach a healthy weight."
	case bmi >= 18.5 && bmi <= 24.9:
		category = "Normal weight"
		tip = "Maintain your healthy lifestyle with a balanced diet and regular physical activity."
	case bmi >= 25 && bmi <= 29.9:
		category = "Overweight"
		tip = "Incorporate more physical activity into your \n routine and watch your portion sizes to manage your weight effectively."
	case bmi >= 30 && bmi <= 34.9:
		category = "Obesity (Class 1)"
		tip = "Focus on a healthier diet and regular exercise.\n It’s advisable to consult with a healthcare provider for personalized advice."
	case bmi >= 35 && bmi <= 39.9:
		category = "Obesity (Class 2)"
		tip = "Medical advice is recommended to develop a weight loss plan,\n as well as to monitor any related health conditions."
	default:
		category = "Severe Obesity (Class 3)"
		tip = "Seek professional medical assistance to address potential \n health risks and to develop a comprehensive weight management plan."
	}

	return bmi, category, tip
}

func main() {
	app := app.New()
	w := app.NewWindow("BMI Buddy")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(700, 500))

	weight := widget.NewEntry()
	height := widget.NewEntry()
	bmiText := widget.NewLabel("")
	categoryText := widget.NewLabel("")
	tipText := widget.NewLabel("")

	themeStyle := widget.NewRadioGroup([]string{"Light", "Dark"}, func(value string) {
		if value == "Light" {
			app.Settings().SetTheme(theme.LightTheme())
		} else {
			app.Settings().SetTheme(theme.DarkTheme())
		}
	})

	themeStyle.SetSelected("Light")

	w.SetContent(container.NewVBox(
		widget.NewLabel("Weight (kg):"),
		weight,
		widget.NewLabel("Height (m):"),
		height,

		widget.NewButton("Calculate", func() {
			weightValue, err1 := strconv.ParseFloat(weight.Text, 64)
			if err1 != nil {
				fmt.Println("Error converting weight:", err1)
				bmiText.SetText("Invalid weight input")
				fmt.Println(err1)
				return
			}
			heightValue, err := strconv.ParseFloat(height.Text, 64)
			if err != nil {
				fmt.Println("Error converting height:", err)
				bmiText.SetText("Invalid height input")
				fmt.Println(err)
				return
			}
			bmi, category, tip := calculateBMI(weightValue, heightValue)
			bmiText.SetText(fmt.Sprintf("Your BMI is: %.2f", bmi))
			categoryText.SetText("Category: " + category)
			tipText.SetText("Tip: " + tip)

			// Update the UI
			fmt.Println("Done!")

		}),
		bmiText,
		categoryText,
		tipText,
		themeStyle,
	))
	w.ShowAndRun()
}
