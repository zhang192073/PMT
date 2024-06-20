rm -rf pentest_mac_tools.app
/Users/penguin/Data/Coder/Go/bin/darwin_amd64/fyne  package -os darwin -icon Icon1.png

cd pentest_mac_tools.app/Contents/Resources
mkdir resources

cp -r /Users/penguin/Data/Coder/Go_Coder/pentest_mac_tools/resources/* /Users/penguin/Data/Coder/Go_Coder/pentest_mac_tools/pentest_mac_tools.app/Contents/Resources/resources/
cd /Users/penguin/Data/Coder/Go_Coder/pentest_mac_tools/
open pentest_mac_tools.app

#/Users/penguin/Data/Coder/Go/bin/darwin_amd64/fyne  bundle resources/config/config.yaml >> config.go
