package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// "log"
	"net/http"
	"time"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// var db *gorm.DB

type AnimeModel struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			CreatedAt           time.Time `json:"createdAt"`
			UpdatedAt           time.Time `json:"updatedAt"`
			Slug                string    `json:"slug"`
			Synopsis            string    `json:"synopsis"`
			Description         string    `json:"description"`
			CoverImageTopOffset int       `json:"coverImageTopOffset"`
			Titles              struct {
				En   string `json:"en"`
				EnJp string `json:"en_jp"`
				JaJp string `json:"ja_jp"`
			} `json:"titles"`
			CanonicalTitle    string        `json:"canonicalTitle"`
			AbbreviatedTitles []interface{} `json:"abbreviatedTitles"`
			AverageRating     interface{}   `json:"averageRating"`
			RatingFrequencies struct {
				Num4  string `json:"4"`
				Num9  string `json:"9"`
				Num10 string `json:"10"`
			} `json:"ratingFrequencies"`
			UserCount      int         `json:"userCount"`
			FavoritesCount int         `json:"favoritesCount"`
			StartDate      string      `json:"startDate"`
			EndDate        string      `json:"endDate"`
			NextRelease    interface{} `json:"nextRelease"`
			PopularityRank int         `json:"popularityRank"`
			RatingRank     interface{} `json:"ratingRank"`
			AgeRating      string      `json:"ageRating"`
			AgeRatingGuide string      `json:"ageRatingGuide"`
			Subtype        string      `json:"subtype"`
			Status         string      `json:"status"`
			Tba            interface{} `json:"tba"`
			PosterImage    struct {
				Tiny     string `json:"tiny"`
				Large    string `json:"large"`
				Small    string `json:"small"`
				Medium   string `json:"medium"`
				Original string `json:"original"`
				Meta     struct {
					Dimensions struct {
						Tiny struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"tiny"`
						Large struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"large"`
						Small struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"small"`
						Medium struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"medium"`
					} `json:"dimensions"`
				} `json:"meta"`
			} `json:"posterImage"`
			CoverImage     interface{} `json:"coverImage"`
			EpisodeCount   int         `json:"episodeCount"`
			EpisodeLength  int         `json:"episodeLength"`
			TotalLength    int         `json:"totalLength"`
			YoutubeVideoID interface{} `json:"youtubeVideoId"`
			ShowType       string      `json:"showType"`
			Nsfw           bool        `json:"nsfw"`
		} `json:"attributes"`
		Relationships struct {
			Genres struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"genres"`
			Categories struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"categories"`
			Castings struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"castings"`
			Installments struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"installments"`
			Mappings struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mappings"`
			Reviews struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"reviews"`
			MediaRelationships struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mediaRelationships"`
			Characters struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"characters"`
			Staff struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"staff"`
			Productions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"productions"`
			Quotes struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"quotes"`
			Episodes struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"episodes"`
			StreamingLinks struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"streamingLinks"`
			AnimeProductions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"animeProductions"`
			AnimeCharacters struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"animeCharacters"`
			AnimeStaff struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"animeStaff"`
		} `json:"relationships"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

func main() {

	//Load environment variables from file.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Dbdriver := os.Getenv("DB_DRIVER")
	// DbHost := os.Getenv("DB_HOST")
	// DbUser := os.Getenv("DB_USER")
	// DbPassword := os.Getenv("DB_PASSWORD")
	// DbName := os.Getenv("DB_NAME")
	// DbPort := os.Getenv("DB_PORT")

	// DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	// db, err = gorm.Open("mysql", DBURL)

	// if err != nil {
	// 	fmt.Printf("Cannot connect to %s database", Dbdriver)
	// 	log.Fatal("This is the error:", err)
	// } else {
	// 	fmt.Printf("We are connected to the %s database", Dbdriver)
	// }

	// Connect to PlanetScale database using DSN environment variable.
	// db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to connect to PlanetScale: %v", err)
	// }
	baseUrl := os.Getenv("API_URL")

	fmt.Println(baseUrl)

	for offset := 0; offset < 6; offset = offset + 2 {
		url := fmt.Sprintf(`%s%d`, baseUrl, offset)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var anime AnimeModel

		err = json.Unmarshal([]byte(body), &anime)

		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(anime.Data); i++ {
			fmt.Print(anime.Data[i].Attributes.Titles.En)
		}
		fmt.Print("\n \n next \n \n")

		time.Sleep(time.Millisecond * 100)
	}

}
