package seed

import (
	"log"

	"github.com/BlackBoyZoovie/fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Ilesanmi Similoluwa",
		Email:    "oladayoilesanmi@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Gbolahan Fakorede",
		Email:    "gbolahanfakorede@gmail.com",
		Password: "password",
	},
}

var post = []models.Post{
	models.Post{
		Title:   "My Profile",
		Content: "I am Similoluwa, I am about 5'10 or 5'11ft tall. I am a golang web developer, hoping to learn more.",
	},

	models.Post{
		Title:   "My Bio",
		Content: "I am Gbolahan, I am a short man, I mean very very short.",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		post[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&post[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
