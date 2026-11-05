// Lyrinox machID wrapper - Only allows eMachID generation
// This wrapper restricts users to generating ephemeral machine IDs only.
// reMachID generation is reserved for system administrators and internal use.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LyrinoxTechnologies/machID"
)

const (
	defaultSalt = "lyrinox-cloud-platform-2024"
)

func main() {
	// Parse command-line flags
	salt := flag.String("salt", defaultSalt, "Salt for ID generation (optional)")
	showHelp := flag.Bool("help", false, "Show help message")
	flag.BoolVar(showHelp, "h", false, "Show help message")
	
	flag.Parse()

	if *showHelp {
		printHelp()
		os.Exit(0)
	}

	// Check for root privileges
	if os.Geteuid() != 0 {
		fmt.Fprintln(os.Stderr, "Error: This program must be run with root privileges")
		fmt.Fprintln(os.Stderr, "Usage: sudo lyrinox-machid")
		os.Exit(1)
	}

	// Generate eMachID only
	emachid, err := machid.GenerateEMachID(*salt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating machine ID: %v\n", err)
		os.Exit(1)
	}

	// Output the ID
	fmt.Println(emachid)
}

func printHelp() {
	fmt.Println("Lyrinox Machine ID Generator")
	fmt.Println("=============================")
	fmt.Println()
	fmt.Println("Generates a unique, ephemeral machine identifier (eMachID) using")
	fmt.Println("kernel-level hardware interrogation.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  sudo lyrinox-machid [OPTIONS]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -salt string")
	fmt.Println("        Salt for ID generation (default: lyrinox-cloud-platform-2024)")
	fmt.Println("  -h, -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  Prints a unique machine ID to stdout")
	fmt.Println()
	fmt.Println("Security:")
	fmt.Println("  • Requires root privileges (setuid binary)")
	fmt.Println("  • Each invocation generates a UNIQUE ID")
	fmt.Println("  • ID cannot be regenerated or reconstructed")
	fmt.Println("  • Uses hardware-specific kernel data")
	fmt.Println()
	fmt.Println("Note:")
	fmt.Println("  This tool only generates ephemeral IDs (eMachID).")
	fmt.Println("  Reconstructable IDs (reMachID) are reserved for system use.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  sudo lyrinox-machid")
	fmt.Println("  sudo lyrinox-machid -salt \"my-custom-salt\"")
	fmt.Println()
}