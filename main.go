// Package pentest_mac_tools -----------------------------
// @file      : main.go
// @auther    : Penguin
// @contact   : 1920732889@qq.com
// @time      : 2024/6/6 12:25
// -------------------------------------------
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lengzhao/font/autoload"
	"gopkg.in/yaml.v3"
	"image/color"
	"log"
	"pentest_mac_tools/model"
	"pentest_mac_tools/tools"
	"strings"
)

type customTheme struct {
	fyne.Theme
}

func (t customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameForeground:
		return color.Black // 设置文字颜色为纯黑以增强对比度和可读性
	case theme.ColorNameButton:
		return color.White
	case theme.ColorNameBackground:
		return color.RGBA{R: 173, G: 216, B: 230, A: 255} // 清新一点的颜色
	default:
		return t.Theme.Color(name, variant)
	}
}

//var resourceConfigYaml = &fyne.StaticResource{
//	StaticName: "config.yaml",
//	StaticContent: []byte(
//		"javapath:\n  Java8: ../../java8/bin/java -jar\n  Java11: ../../java11/bin/java -jar\n  open: open\nredis:\n  host: 127.0.0.1\n  port: 6379\nCategories:\n  信息收集:\n    Fofa_search:\n      path: resources/info/Fofa_search/fofa_search.jar\n      command: Java8\n    webfinder:\n      path: resources/info/webfinder/webfinder-next.jar\n      command: Java8\n    dirscan:\n      path: resources/info/dirscan_3.0/scandir-3.0.jar\n      command: Java8\n    mitan v1.11:\n      path: resources/info/mitan1.11/mitan-jar-with-dependencies.jar\n      command: Java8\n    xray图形化工具:\n      path: resources/info/xray/super-xray-1.7.jar\n      command: Java8\n    gorailgun v1.5.5:\n      path: resources/comprehensive/gorailgun-normal/gorailgun-normal-1.5.5.app\n      command: Open\n    GitHack:\n      path: resources/info/GitHack/.\n      command: Open\n    swagger-hack:\n      path: resources/info/swagger-hack/.\n      command: Open\n    heapdump:\n      path: resources/info/heapdump/.\n      command: Open\n    svnExploit:\n      path: resources/info/svnExploit/.\n      command: Open\n    ds_store_exp:\n      path: resources/info/ds_store_exp/.\n      command: Open\n  ##webshell\n\n  webshell管理工具:\n    哥斯拉 v4.0.1:\n      path: resources/webshell/Godzilla/godzilla.jar\n      command: Java11\n    中国蚁剑 v2.1.15:\n      path: resources/webshell/AntSword/AntSword.app\n      command: Open\n    冰蝎 v4.1:\n      path: resources/webshell/Behinder_v4.1.t00ls/Behinder.jar\n      command: Java11\n    冰蝎 v3.0.11:\n      path: resources/webshell/Behinder_v3.0.11/Behinder.jar\n      command: Java11\n    冰蝎魔改 v3.3.2:\n      path: resources/webshell/Behinder-Mode/Behinder-Mode.jar\n      command: Java11\n    天蝎 v1.0:\n      path: resources/webshell/TianXie/Tianxie.jar\n      command: Java11\n    Webshell生成 v1.2.4:\n      path: resources/webshell/Webshell_Generate/Webshell_Generate-1.2.4.jar\n      command: Java8\n\n#  \"框架利用\":    {\"shiro_attack v4.7\", \"Pyke-Shiro v0.3\", \"Shiro550NoCC\",\n#    \"Nacos漏洞利用工具\", \"JenKins漏洞利用工具\", \"Jboss漏洞利用\", \"Ruoyi综合利用\", \"ThinkphpGUI\", \"SpringBootExploit v1.3\",\n#    \"Struts2 v19.23\", \"WeblogicTool v1.3\", \"AttackTomcat利用工具\", \"druid利用工具\"},\n  Framework_utilization:\n    shiro_attack v4.7:\n      path: resources/framework/shiro/shiro_attack-4.7/shiro_attack-4.7.0-SNAPSHOT-all.jar\n      command: Java8\n    Pyke-Shiro:\n      path: resources/framework/shiro/shiro_attack2/ Pyke-Shiro_0.3.jar\n      command: Java8\n    Shiro550NoCC:\n      path: resources/framework/shiro/Shiro-550-with-NoCC.jar\n      command: Java8\n    Nacos漏洞利用工具:\n      path: resources/framework/Nacos/NacosExploitGUI_v4.0.jar\n      command: Java8\n    JenKins漏洞利用工具:\n      path: resources/framework/JenKins/JenkinsExploit-GUI-1.3-SNAPSHOT.jar\n      command: Java8\n    Jboss漏洞利用:\n      path: resources/framework/jboss/JavaJboss.jar\n      command: Java8\n    Ruoyi综合利用:\n      path: resources/framework/Ruoyi/ruoyiVuln.jar\n      command: Java8\n    ThinkphpGUI:\n      path: resources/framework/thinkphp/ThinkphpGUI.jar\n      command: Java8\n    SpringBootExploit v1.3:\n      path: resources/framework/Spring/SpringBootExploit-1.3-SNAPSHOT-all/SpringBootExploit-1.3-SNAPSHOT-all.jar\n      command: Java8\n    Struts2 v19.23:\n      path: resources/framework/struts2/struts2_19.jar\n      command: Java8\n    WeblogicTool v1.3:\n      path: resources/framework/weblogic/WeblogicTool_1.3.jar\n      command: Java8\n    AttackTomcat利用工具:\n      path: resources/framework/tomcat/AttackTomcat.jar\n      command: Java8\n    druid利用工具:\n      path: resources/framework/druid/druid_sessions-1.0-SNAPSHOT.jar\n      command: Java8\n  Comprehensive_utilization_tools:\n    Thelostworld OA v1.1:\n      path: resources/comprehensive/Thelostworld_OA/Thelostworld_OA.jar\n      command: Java8\n    LiqunKit 工具箱 v1.6.2:\n      path: resources/comprehensive/LiqunKit/LiqunKit_1.6.2.jar\n      command: Java8\n    2023 0day利用:\n      path: resources/comprehensive/day2023/day2023.jar\n      command: Java8\n    Hvv综合利用:\n      path: resources/comprehensive/HVVExploitApply/HVVExploitApply.jar\n      command: Java8\n    APT Tools魔改:\n      path: resources/comprehensive/apt/IWannaGetAll.jar\n      command: Java8\n    0x7eTeamTools:\n      path: resources/comprehensive/0x7eTeamTools/0x7e.jar\n      command: Java8\n    MyExploit:\n      path: resources/comprehensive/MyExploit/MYExploit.jar\n      command: Java8\n    poc2jar综合利用工具:\n      path: resources/comprehensive/poc2jar-MACOS/poc2jar.jar\n      command: Java8\n    hyacinth综合利用工具:\n      path: resources/comprehensive/hyacinth/hyacinth.jar\n      command: Java8\n  Database_utilization_tools:\n    OracleShell v0.1:\n      path: resources/databases/OracleShell/oracleShell.jar\n      command: Java8\n    postgersql利用工具:\n      path: resources/databases/postgre/postgreUtil-1.0-SNAPSHOT-jar-with-dependencies.jar\n      command: Java8\n    MDUT v2.1.1:\n      path: resources/databases/MDUT/MDUT.jar\n      command: Java8\n  Proxy_tools:\n    Erfrp:\n      path: resources/proxy/Erfrp/.\n      command: Open\n    Neo-reGeorg-5:\n      path: resources/proxy/Neo-reGeorg-5.0.1/.\n      command: Open\n    Venom:\n      path: resources/proxy/Venom/.\n      command: Open\n#    Neo-reGeorg-5:\n#      path: resources/proxy/Neo-reGeorg-5.0.1/.\n#      command: Open\n\n  Other_tools:\n    蓝队分析工具箱 V0.92:\n        path: resources/other/blue/BlueTeamTools.jar\n        command: Java8\n    oss-browser:\n      path:  rresources/other/oss-browser-darwin-x64/oss-browser.app\n      command: Open\n    API-Tool v1.2:\n      path: resources/comprehensive/MyExploit/MYExploit.jar\n      command: Java8\n    PotatoTool v1.1:\n      path: resources/other/PotatoTool/PotatoTool-1.1-jdk11.jar\n      command: Java8\n    JMG内存马生成 v1.0.6:\n      path: resources/other/jMG/jMG-1.0.6-GUI.jar\n      command: Java8\n    RequestTemplate:\n      path: resources/other/RequestTemplate/RequestTemplate.jar\n      command: Java8\n    内网横向工具:\n      path: resources/other/gogogo-jar-with-dependencies/gogogo-jar-with-dependencies.jar\n      command: Java8\n    DecryptTools加解密工具:\n      path: resources/other/DecryptTools/DecryptTools.jar\n      command: Java8\n    JD-GUI 1.6.6:\n      path: resources/other/JD-GUI1.6.6/jd-gui-1.6.6.jar\n      command: Java8\n    jadx-gui:\n      path: resources/other/jadx-1.5.0/bin/jadx-gui\n      command: Open\n    apktool:\n      path: resources/other/apktool/.\n      command: Open\n    recode:\n      path: resources/other/recode/.\n      command: Open"),
//}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(customTheme{theme.LightTheme()})
	myWindow := myApp.NewWindow("Penguin_Penetration_Testing_Tools")

	background := canvas.NewRectangle(color.RGBA{R: 173, G: 216, B: 230, A: 255}) // 设置清新背景颜色

	outputLabel := widget.NewLabel("Output will be shown here")
	outputLabel.Wrapping = fyne.TextWrapBreak

	var categories model.Categories
	//file, err := ioutil.ReadFile("resources/config/test.yaml")
	yamldata, err := tools.ReadYAMLFile()
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(yamldata, &categories)

	var canvasObjects []fyne.CanvasObject
	var allButtons []*widget.Button
	var categoryContainers []fyne.CanvasObject

	for _, category := range categories.Category {
		label := widget.NewLabel(category.Name)
		labelContainer := container.NewMax(label)
		var buttons []fyne.CanvasObject
		for _, toolbase := range category.Task {
			toolCopy := toolbase.Name // 捕获工具名称以用于回调函数
			btn := widget.NewButton(toolbase.Name, func(tool string) func() {
				return func() {
					err := tools.ExecuteCommand(yamldata, toolbase.Name, toolbase.PATH, toolbase.Optional, toolbase.VALUE, toolbase.COMMAND)
					if err != nil {
						outputLabel.SetText("Error: " + err.Error())
					} else {
						outputLabel.SetText("Running: " + tool)
					}
				}
			}(toolCopy))
			buttons = append(buttons, btn)
			allButtons = append(allButtons, btn)
		}
		catContainer := container.NewVBox(labelContainer)
		catContainer.Add(container.NewGridWrap(fyne.NewSize(160, 30), buttons...))
		canvasObjects = append(canvasObjects, catContainer)
		categoryContainers = append(categoryContainers, catContainer)

	}
	mainContent := container.NewVBox(canvasObjects...)
	scrollableContent := container.NewScroll(mainContent)
	scrollableContent.SetMinSize(fyne.NewSize(850, 650))

	backgroundWithContent := container.NewMax(background, scrollableContent)
	// 创建搜索框
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("搜索工具...")

	// 搜索框内容变化时过滤显示的按钮
	searchEntry.OnChanged = func(s string) {
		if s == "" {
			scrollableContent.Content = container.NewVBox(categoryContainers...)
		} else {
			filteredObjects := []fyne.CanvasObject{}
			s = strings.ToLower(s)
			for _, btn := range allButtons {
				if strings.Contains(strings.ToLower(btn.Text), s) {
					filteredObjects = append(filteredObjects, btn)
				}
			}
			scrollableContent.Content = container.NewVBox(filteredObjects...)
		}
		scrollableContent.Refresh()
	}

	// 创建取消搜索按钮
	clearButton := widget.NewButton("取消搜索", func() {
		searchEntry.SetText("")
		scrollableContent.Content = container.NewVBox(categoryContainers...)
		scrollableContent.Refresh()
	})

	searchContainer := container.NewBorder(nil, nil, nil, clearButton, searchEntry)

	content := container.NewBorder(searchContainer, outputLabel, nil, nil, backgroundWithContent)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}

//func main() {
//	//读取配置文件
//
//	yamldata := resourceConfigYaml.StaticContent
//	var categories model.Categories
//	yaml.Unmarshal(yamldata, &categories)
//
//	myApp := app.New()
//	myApp.Settings().SetTheme(customTheme{theme.LightTheme()})
//	myWindow := myApp.NewWindow("hss")
//
//	background := canvas.NewRectangle(color.RGBA{R: 173, G: 216, B: 230, A: 255}) // 设置清新背景颜色
//
//	outputLabel := widget.NewLabel("Output will be shown here")
//	outputLabel.Wrapping = fyne.TextWrapBreak
//
//	var canvasObjects []fyne.CanvasObject
//	var allButtons []*widget.Button
//	var categoryContainers []fyne.CanvasObject
//	fmt.Println(categories.Categories)
//	for categoryName, category := range categories.Categories {
//		fmt.Printf("Category: %s\n", categoryName)
//		label := widget.NewLabel(categoryName)
//		labelContainer := container.NewMax(label)
//		var buttons []fyne.CanvasObject
//		for toolName, toolbase := range category.Tools {
//			//fmt.Printf("  Tool: %s\n", toolName)
//			//fmt.Printf("    Path: %s\n", toolbase.Path)
//			//fmt.Printf("    Command: %s\n", toolbase.Command)
//			toolCopy := toolName // Capture the tool name for the lambda function
//			btn := widget.NewButton(toolName, func(tool string) func() {
//				return func() {
//					err := tools.ExecuteCommand(yamldata, toolName, toolbase.Path, toolbase.Command)
//					if err != nil {
//						outputLabel.SetText("Error: " + err.Error())
//					} else {
//						outputLabel.SetText("Running: " + tool)
//					}
//				}
//			}(toolCopy))
//			buttons = append(buttons, btn)
//			allButtons = append(allButtons, btn)
//		}
//		catContainer := container.NewVBox(labelContainer)
//		catContainer.Add(container.NewGridWrap(fyne.NewSize(160, 30), buttons...))
//		canvasObjects = append(canvasObjects, catContainer)
//		categoryContainers = append(categoryContainers, catContainer)
//
//	}
//	mainContent := container.NewVBox(canvasObjects...)
//	scrollableContent := container.NewScroll(mainContent)
//	scrollableContent.SetMinSize(fyne.NewSize(850, 650))
//
//	backgroundWithContent := container.NewMax(background, scrollableContent)
//	searchEntry := widget.NewEntry()
//	searchEntry.SetPlaceHolder("搜索工具...")
//
//	// 搜索框内容变化时过滤显示的按钮
//	searchEntry.OnChanged = func(s string) {
//		if s == "" {
//			scrollableContent.Content = container.NewVBox(categoryContainers...)
//		} else {
//			filteredObjects := []fyne.CanvasObject{}
//			s = strings.ToLower(s)
//			for _, btn := range allButtons {
//				if strings.Contains(strings.ToLower(btn.Text), s) {
//					filteredObjects = append(filteredObjects, btn)
//				}
//			}
//			scrollableContent.Content = container.NewVBox(filteredObjects...)
//		}
//		scrollableContent.Refresh()
//	}
//
//	// 创建取消搜索按钮
//	clearButton := widget.NewButton("取消搜索", func() {
//		searchEntry.SetText("")
//		scrollableContent.Content = container.NewVBox(categoryContainers...)
//		scrollableContent.Refresh()
//	})
//
//	searchContainer := container.NewBorder(nil, nil, nil, clearButton, searchEntry)
//
//	content := container.NewBorder(searchContainer, outputLabel, nil, nil, backgroundWithContent)
//
//	myWindow.SetContent(content)
//	myWindow.Resize(fyne.NewSize(800, 600))
//	myWindow.CenterOnScreen()
//	myWindow.ShowAndRun()
//}
