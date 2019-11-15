package main

import (
    "github.com/go-flutter-desktop/go-flutter"
    "io/ioutil"
    //"os"
    //"path/filepath"
    "strconv"
)

func getPixelRatio() float64 {
    //s, _ := filepath.Abs(os.Args[0])
    //dir := filepath.Dir(s)
    //_ = ioutil.WriteFile("/tmp/config.txt", []byte(dir+"/ratio.txt"), os.ModePerm)
    bytes, e := ioutil.ReadFile("ratio.txt")
    if e != nil {
        return float64(2)
    }
    f, e := strconv.ParseFloat(string(bytes), 64)
    if e != nil {
        return float64(2)
    }
    return f
}

var options = []flutter.Option{
    flutter.WindowInitialDimensions(1920, 1080),
    flutter.ForcePixelRatio(getPixelRatio()), // retina set to 2
}
