# rm -rf go/build/outputs

~/go/bin/hover build darwin
# hover build darwin

echo "build to next files:"
ls go/build/outputs/darwin

echo "copy tempate"
cp -r go/template/darwin/JsonToDart.app go/build/outputs

echo "move binary files to app"
mv go/build/outputs/darwin/* go/build/outputs/JsonToDart.app/Contents/MacOS

cd go/build/outputs
tar cvzf JsonToDart-darwin.tar.gz JsonToDart.app