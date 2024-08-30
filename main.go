package main

import (
	"fmt"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/julien-noblet/noutilus/lib"
)

var MaxSize = 10

type operation struct {
	resultat  int
	operateur int
	/*
		+ = 0
		- = 1
		* = 2
		/ = 3
	*/
	label *widget.Label
	// choose    *widget.Select
	container *fyne.Container
}

func (o *operation) GetCalc() {
	if o.resultat > 0 {
		switch o.operateur {
		case 0:
			t, err := lib.RandIntMax(o.resultat)
			if err != nil {
				o.label.SetText("Une erreur c'est produite, le nombre est surement trop petit!")

				return
			}

			o.label.SetText(fmt.Sprintf("%v = %v + %v", o.resultat, o.resultat-t, t))

		case 1:
			t, err := lib.RandIntMin(o.resultat)
			if err != nil {
				o.label.SetText("Une erreur c'est produite, le nombre est surement trop grand!")

				return
			}

			o.label.SetText(fmt.Sprintf("%v = %v - %v", o.resultat, o.resultat+t, t))

		case 2: //nolint:gomnd // why checking on this switch :/
			t1, t2, err := lib.Find2Factors(o.resultat)
			if err != nil {
				o.label.SetText("Une erreur c'est produite, le nombre est surement premier!")

				return
			}

			o.label.SetText(fmt.Sprintf("%v = %v × %v", o.resultat, t1, t2))
		}
	}
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Noutilus")

	motMystere := ""
	scrambled, _ := lib.ShuffleWord(lib.AddNoise(lib.ReduceUniqueLetters(motMystere), 10))
	number := lib.ConcatInt(lib.Numerize(scrambled, motMystere))
	nbOperationInternal := 1
	title := container.New(layout.NewCenterLayout(), widget.NewLabel("Noutilus"))

	scrambledWord := widget.NewLabel(scrambled)
	scrambleTable := container.NewHBox(
		container.NewVBox(
			container.NewHBox(widget.NewLabel(""), scrambledWord),
			container.NewHBox(widget.NewLabel(""), widget.NewLabel("0123456789")),
		),
		container.New(layout.NewCenterLayout(), widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
			scrambled, _ = lib.ShuffleWord(lib.AddNoise(lib.ReduceUniqueLetters(motMystere), 10))
			scrambledWord.SetText(scrambled)
			number = lib.ConcatInt(lib.Numerize(scrambled, motMystere))
		})),
	)
	nbOperationsLabel := widget.NewLabel("Nombre d'operations : ")
	nbOperations := widget.NewLabel("1")
	nbOperationsSlider := widget.NewSlider(float64(1), float64(MaxSize))
	operations := container.NewVBox()

	nbOperationContainer := container.NewVBox(container.NewHBox(nbOperationsLabel, nbOperations), nbOperationsSlider)
	nbOperationsSlider.OnChanged = func(nb float64) {
		nbOperations.SetText(fmt.Sprintf("%.f", nb))
		nbOperationInternal = int(nb)
	}
	btnRefresh := widget.NewButton("Calcul", func() {
		operations := myApp.NewWindow("Calculs")
		ops := container.NewVBox()
		fmt.Println("On Changed!")
		reste := number
		for i := 0; i < nbOperationInternal; i++ {
			labelOp := widget.NewLabel("")
			containerOP := container.NewHBox(labelOp)

			operation := &operation{
				label:     labelOp,
				container: containerOP,
				operateur: 0,
			}
			chooseOp := widget.NewSelect([]string{
				"+",
				"-",
				"×",
			}, func(s string) {
				switch s {
				case "+":
					operation.operateur = 0
					operation.GetCalc()
				case "-":
					operation.operateur = 1
					operation.GetCalc()
				case "×":
					operation.operateur = 2
					operation.GetCalc()
				}
			})
			operation.container.Add(chooseOp)
			if i < nbOperationInternal-1 {
				operation.resultat, _ = lib.RandIntMax(reste)
				reste -= operation.resultat
			} else {
				operation.resultat = reste
			}

			operation.GetCalc()
			ops.Add(operation.container)
		}
		operations.SetContent(ops)
		operations.Show()
	})

	motMystereLabel := widget.NewLabel("Mot Mystère: ")
	motMystereInput := widget.NewEntry()
	motMystereInput.SetPlaceHolder("Mot Mystère")

	motMystereInputSize := widget.NewLabel(fmt.Sprintf("Taille du mot: %v", len(lib.ReduceUniqueLetters(motMystereInput.Text))))
	motMystereInput.OnChanged = func(s string) {
		motMystereInputSize.SetText(fmt.Sprintf("Taille du mot: %v", len(lib.ReduceUniqueLetters(s))))
		motMystere = s
		scrambled, _ = lib.ShuffleWord(lib.AddNoise(lib.ReduceUniqueLetters(motMystere), 10))
		scrambledWord.SetText(scrambled)
		number = lib.ConcatInt(lib.Numerize(scrambled, motMystere))
	}
	motMystereBox := container.NewVBox(container.NewVBox(motMystereLabel, motMystereInput), motMystereInputSize)
	container := container.NewVBox(title, motMystereBox, scrambleTable, nbOperationContainer, btnRefresh, operations)

	w.SetContent(container)
	w.ShowAndRun()
}
