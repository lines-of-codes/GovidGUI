package main

import (
	"fmt"
	"github.com/tadvi/winc"
	"github.com/YeffyCodeGit/Govid/govid"
)

func onWindowClose(arg *winc.Event) {
	winc.Exit()
}

func main() {
	LogInfo("Initializing GUIs...")
	window := winc.NewForm(nil)
	window.SetSize(350, 315)
	window.SetText("GovidGUI")

	countrySelection := winc.NewComboBox(window)

	data, err := govid.GetAllCountriesData()

	if err != nil {
		panic(err)
	}

	countryNames := make([]govid.CountryData, 0)
	for index, value := range data {
		countrySelection.InsertItem(index, value.Country)
		countryNames = append(countryNames, value)
	}
	countrySelection.SetSelectedItem(0)
	countrySelection.SetPos(75, 20)

	todayCases := winc.NewLabel(window)
	todayCases.SetText(fmt.Sprintf("Today cases: %d", countryNames[0].TodayCases))
	todayCases.SetPos(30, 60)
	todayCases.SetSize(200, 25)

	totalCases := winc.NewLabel(window)
	totalCases.SetText(fmt.Sprintf("Total cases: %d", countryNames[0].Cases))
	totalCases.SetPos(30, 90)
	totalCases.SetSize(200, 25)

	activeCases := winc.NewLabel(window)
	activeCases.SetText(fmt.Sprintf("Active cases: %d", countryNames[0].Active))
	activeCases.SetPos(30, 120)
	activeCases.SetSize(200, 25)

	todayDeath := winc.NewLabel(window)
	todayDeath.SetText(fmt.Sprintf("Today deaths: %d", countryNames[0].TodayDeaths))
	todayDeath.SetPos(30, 150)
	todayDeath.SetSize(200, 25)

	totalDeath := winc.NewLabel(window)
	totalDeath.SetText(fmt.Sprintf("Total deaths: %d", countryNames[0].Deaths))
	totalDeath.SetPos(30, 180)
	totalDeath.SetSize(200, 25)

	totalRecovered := winc.NewLabel(window)
	totalRecovered.SetText(fmt.Sprintf("Total recovered: %d", countryNames[0].Recovered))
	totalRecovered.SetPos(30, 210)
	totalRecovered.SetSize(200, 25)

	countrySelection.OnSelectedChange().Bind(func(arg *winc.Event) {
		countryData := countryNames[countrySelection.SelectedItem()]
		todayCases.SetText(fmt.Sprintf("Today cases: %d", countryData.TodayCases))
		totalCases.SetText(fmt.Sprintf("Total cases: %d", countryData.Cases))
		activeCases.SetText(fmt.Sprintf("Active cases: %d", countryData.Active))
		todayDeath.SetText(fmt.Sprintf("Today deaths: %d", countryData.TodayDeaths))
		totalDeath.SetText(fmt.Sprintf("Total deaths: %d", countryData.Deaths))
		totalRecovered.SetText(fmt.Sprintf("Total recovered: %d", countryData.Recovered))
	})

	fetchAllButton := winc.NewPushButton(window)
	fetchAllButton.SetText("Fetch data (all)")
	fetchAllButton.OnClick().Bind(func(arg *winc.Event) {
		data, err := govid.GetAllCountriesData()

		if err != nil {
			panic(err)
		}

		for index, value := range data {
			countryNames[index] = value
		}
	})
	fetchAllButton.SetPos(20, 240)

	fetchThisButton := winc.NewPushButton(window)
	fetchThisButton.SetText("Fetch data (this)")
	fetchThisButton.OnClick().Bind(func(arg *winc.Event) {
		selection := countrySelection.SelectedItem()
		newData, err := govid.GetCountryData(countryNames[selection].Country)

		if err != nil {
			panic(err)
		}

		countryNames[selection] = *newData
	})
	fetchThisButton.SetPos(125, 240)

	LogInfo("UI created.")

	window.Center()
	window.Show()
	window.OnClose().Bind(onWindowClose)

	winc.RunMainLoop()
}
