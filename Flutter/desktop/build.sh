flutter pub get
flutter config --enable-macos-desktop
go get -u github.com/go-flutter-desktop/hover
~/go/bin/hover build darwin
ls go/build/outputs/darwin
cd go/build/outputs
tar cvzf darwin.tar.gz darwin