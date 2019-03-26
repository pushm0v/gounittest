package cmd

import (
	"github.com/spf13/cobra"

	t "github.com/pushm0v/gounittest/toko"
)

// tokoCmd represents the toko command
var tokoCmd = &cobra.Command{
	Use:   "toko",
	Short: "Start toko app",
	Long:  `Toko app is our simple app for selling goods`,
	Run: func(cmd *cobra.Command, args []string) {
		toko := t.NewToko()
		toko.Buka()
		minuman := t.Minuman{
			Nama:  "Sprite",
			Harga: 10000,
		}
		toko.TambahStokMinuman(minuman, 10)
		toko.JualMinuman(minuman, 10)
	},
}

func init() {
	rootCmd.AddCommand(tokoCmd)
}
