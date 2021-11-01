package main

import (
	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	type Student struct {
		Name  string
		Phone string
	}

	var myStudentData []Student

	data_from_file, _ := ioutil.ReadFile("myFile_name.txt")
	//Ingoring the error thats why underscore
	json.Unmarshal(data_from_file, &myStudentData)
	a := app.New()
	w := a.NewWindow("Data_Recorder")
	w.Resize(fyne.NewSize(400, 400))
	e_name := widget.NewEntry()
	e_name.SetPlaceHolder("Enter Name")
	// entry widget for phone
	e_phone := widget.NewEntry()
	e_phone.SetPlaceHolder("Enter Phone Number")
	// submit button
	submit_btn := widget.NewButton("Submit", func() {
		// let create a struct for our data

		myData1 := &Student{
			Name:  e_name.Text, // data from name entry widget
			Phone: e_phone.Text,
		}
		// append / push data to our slice
		myStudentData = append(myStudentData, *myData1)
		// convert / parse data to json format
		// 1st : slicing of data
		// 2nd : we are using space to indent our data
		final_data, _ := json.MarshalIndent(myStudentData, "", " ")
		// writing data to file
		//file name .txt or .json or anything you want to use
		// data source, final_data in our case
		// writing/reading/execute permission
		ioutil.WriteFile("myFile_name.txt", final_data, 0644)

		// lets test
		// empty input field after data insertion
		e_name.Text = ""
		e_phone.Text = ""
		// refresh name & phone entry object
		e_name.Refresh()
		e_phone.Refresh()
	})
	// list widget
	list := widget.NewList(
		// first argument is item count
		// len() function to get myStudentData slice len
		func() int { return len(myStudentData) },
		// 2nd argument is for widget choice. I want to use label
		func() fyne.CanvasObject { return widget.NewLabel("") },
		// 3rd argument is to update widget with our new data
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myStudentData[lii].Name)
		},
	)

	w.SetContent(
		//creating Hsplit container cause I want the appt to be displayed in two columns

		container.NewHSplit(
			// first argument is list of data
			list, //Recalling list
			// vbox containersc
			container.NewVBox(e_name, e_phone, submit_btn),
		),
	)
	w.ShowAndRun()
}
