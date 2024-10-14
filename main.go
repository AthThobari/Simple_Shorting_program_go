package main

import (
  "fmt"
  "strings"
)

func selectionSort(arr []string) {

  for i := 0; i < len(arr)-1; i++ {
    minIndex := i
    for j := i + 1; j < len(arr); j++ {
      if arr[minIndex] > arr[j] {
        minIndex = j
      }
    }
    tmp := arr[i]
    arr[i] = arr[minIndex]
    arr[minIndex] = tmp
  }
}

func main() {
  data := ",Alexa,Azis,Bayu,Chiko,Dewi,Dhanisa,Dian Ningrum,Fikri,Ibnu,Naswa,Salsabila,Shirly,Suci,Syifa,Vita"
    // data := ",Alexa(libur),Azis(libur),Bayu(libur),Chiko(libur),Dewi(izin),Dhanisa(libur),Dian Ningrum(hadir),Fikri(?),Ibnu(hadir),Naswa(izin),Salsabila(hadir),Shirly(hadir),Suci(hadir),Syifa(izin),Vita(izin)"
  
  // Split the string into an array of names
  arr := strings.Split(data, ",")

  // Remove the first empty element if needed
  if arr[0] == "" {
    arr = arr[1:]
  }

  // Sort the array
  selectionSort(arr)

  // Print the sorted names
//   for i, name := range arr {
	// fmt.Printf("%d.%s \n", i +1, name)
//   }

  for _, name := range arr {
	fmt.Printf("%s \n",  name)
  }
}
