package attacks

import (
	"github.com/shosta/androSecTest/command"
	grep "github.com/shosta/androSecTest/command/grep"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/variables"
)

// Voici le commande à utiliser :
// ./ShostaSyncBox/Developpement/HackingTools/DecompilingAndroidApp/jadx/bin/jadx --deobf -d ~/android/security/com.orange.owtv/attacks/decodedPackage ~/android/security/com.orange.owtv/attacks/sourcePackage/com.orange.owtv.apk
func reverseApk(apkname string) {
	// TODO : Il faut changer le chemin absolu vers le binaire de jadx pour que cela soit rentré par l'utilisateur dans un fichier settings.
	cmd := "~/ShostaSyncBox/Developpement/HackingTools/DecompilingAndroidApp/jadx/bin/jadx --deobf -d " +
		DisassemblePackageDirPath(apkname) +
		DecompiledPackageDirPath(apkname) + "/" + apkname + ".apk"
	logging.PrintlnDebug("Cmd : " + cmd)

	logging.Println("Decompiling apk to " + logging.Bold(apkname+"/attacks/decodedPackage/"))
	logging.Println("In progress...")
	command.RunAlias(cmd)
	logging.Println("Done")
}

// DoReverse :
func DoReverse(pkgname string) {
	logging.Println(logging.Green("Reverse apk"))
	reverseApk(pkgname)
	logging.Println(logging.Bold("Done"))

	logging.Println(logging.Green("Check for leakage in codebase"))
	checkForLeakage(pkgname)
	logging.Println(logging.Bold("Done"))
}

func checkForLeakage(pkgname string) {
	decoPkgPath := DecompiledPackageDirPath(pkgname)
	createLeakageDir(pkgname)
	grep.Passwd(decoPkgPath, decoPkgPath+variables.LeakagesDir)
	grep.Admin(decoPkgPath, decoPkgPath+variables.LeakagesDir)
	grep.Key(decoPkgPath, decoPkgPath+variables.LeakagesDir)
}