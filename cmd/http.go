package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/JaneliaSciComp/idservice/id"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var (
	port    string
	workdir string // path to working directory
)

// serveCmd represents the serve command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start HTTP server for unique ids",
	Long: `Start HTTP server for unique ids.

idserver returns unique, monotonically increasing uint64 integers via HTTP.

Example:
% idserver http -p :8000 -w /path/to/workdir

To use, the client sends a POST request with an optional count:

POST /v1/id

	Returns application/json:
	{"id":1}

POST /v1/id?count=10

	Returns application/json:
	{"ids":[1,10]}
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := id.LoadID(workdir); err != nil {
			log.Fatalf("unable to initialize: %v\n", err)
		}
		app := fiber.New()

		app.Post("/v1/id", func(c *fiber.Ctx) error {
			countStr := c.Query("count")
			if countStr == "" {
				id, err := id.GenerateID()
				if err != nil {
					return err
				}
				return c.JSON(fiber.Map{"id": id})
			}
			count, err := strconv.ParseUint(countStr, 10, 64)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"error": fmt.Sprintf("invalid count: %v", err),
				})
			}
			ids, err := id.GenerateIDs(count)
			if err != nil {
				return err
			}
			return c.JSON(fiber.Map{"ids": ids})
		})
		app.Listen(port)
	},
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %v", err)
	}

	rootCmd.AddCommand(httpCmd)

	httpCmd.Flags().StringVarP(&port, "port", "p", ":8000", `port for server`)
	httpCmd.Flags().StringVarP(&workdir, "workdir", "w", pwd, `working directory`)
}
