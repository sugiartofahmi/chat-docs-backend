package seeder

import (
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, seederCommands []string) error {
	listSeeders := map[string]Seeder{}

	if len(seederCommands) > 0 {
		for _, name := range seederCommands {
			s, ok := listSeeders[name]
			if !ok {
				log.Printf("unknown seeder: %s", name)
				continue
			}
			if err := s.Handle(db); err != nil {
				return err
			}
		}
	} else {
		for _, s := range listSeeders {
			if err := s.Handle(db); err != nil {
				return err
			}
		}
	}

	log.Println("Seeding completed!")
	return nil
}
