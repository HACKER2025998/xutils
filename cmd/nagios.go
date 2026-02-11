package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var nagiosCmd = &cobra.Command{
	Use:   "nagios",
	Short: "Install and configure Nagios",
	Run: func(cmd *cobra.Command, args []string) {

		// 1️⃣ Vérifier root
		if os.Geteuid() != 0 {
			fmt.Println("❌ Please run as root (sudo)")
			os.Exit(1)
		}

		// 2️⃣ Vérifier OS
		if runtime.GOOS != "linux" {
			fmt.Println("❌ Only Linux supported for now")
			return
		}

		fmt.Println("🔍 Detecting distribution...")

		// 3️⃣ Installation pour Debian/Ubuntu
		fmt.Println("📦 Installing Nagios...")

		runCommand("apt update")
		runCommand("apt install -y nagios4 nagios-plugins")

		// 4️⃣ Activer service
		runCommand("systemctl enable nagios4")
		runCommand("systemctl restart nagios4")

		// 5️⃣ Vérification
		fmt.Println("🔎 Checking service status...")
		runCommand("systemctl status nagios4")

		fmt.Println("✅ Nagios installation completed!")
		fmt.Println("🌐 Access via: http://SERVER-IP/nagios4")

	},
}

func runCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ Error:", err)
	}
}

func init() {
	installCmd.AddCommand(nagiosCmd)
}
