package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
)

var (
	jwkWriteToDisk        = false
	jwtPrivateKeyFilePath = "jwk.key"
	jwtPublicKeyFilePath  = "jwk.pub"
)

var jwkCmd = &cobra.Command{
	Use:   "jwk",
	Short: "Command group for managing JWK logic",
}

var jwkGenCmd = &cobra.Command{
	Use:     "generate-keyset",
	Aliases: []string{"genkeys"},
	Short:   "Generates a public/private keyset",
	Run: func(_ *cobra.Command, _ []string) {
		c := getAndValidateConfig()
		log := getCmdLogger(c, "jwk").Named("genkeys")

		log.Info("generating jwk keySet")

		keySet, err := token.GenerateRsaKeySet()
		if err != nil {
			log.Fatalf("error generating keySet: %s", err)
		}

		if !jwkWriteToDisk {
			log.Infof("printing keys")
			log.Infof("private Key: \n\n%s\n", keySet.PrivateKey)
			log.Infof("public Key: \n\n%s\n", keySet.PublicKey)
			log.Infof("printed keys, to write to disk, use the --write or -w flag")

			return
		}

		log.Infof("Writing private key to %s", jwtPrivateKeyFilePath)

		err = os.WriteFile(jwtPrivateKeyFilePath, keySet.PrivateKey, 0600) //nolint:gocritic // This is fine
		if err != nil {
			log.Fatalf("error writing private key to disk: %s", err)
		}

		log.Infof("Writing public key to %s", jwtPublicKeyFilePath)

		err = os.WriteFile(jwtPublicKeyFilePath, keySet.PublicKey, 0644) //nolint:gocritic,gosec // This is fine
		if err != nil {
			log.Fatalf("error writing public key to disk: %s", err)
		}

		log.Infof("wrote keys to disk")
	},
}

//nolint:gochecknoinits // Cobra needs usage of init functions
func init() {
	jwkGenCmd.Flags().BoolVarP(&jwkWriteToDisk, "write", "w", false, "Write keys to disk, to jwk.key and jwk.pub")

	// Add sub commands to command
	jwkCmd.AddCommand(jwkGenCmd)

	// Add to root command
	rootCmd.AddCommand(jwkCmd)
}
