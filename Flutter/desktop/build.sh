flutter pub get
flutter config --enable-macos-desktop
go get -u github.com/go-flutter-desktop/hover
$GOPATH/bin/hover build darwin-bundle
ls go/build/outputs/darwin