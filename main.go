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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

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
			AverageRating     string        `json:"averageRating"`
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
			RatingRank     int         `json:"ratingRank"`
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

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err = gorm.Open("mysql", DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	db.AutoMigrate(&Anime{})

	// Connect to PlanetScale database using DSN environment variable.
	// db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to connect to PlanetScale: %v", err)
	// }
	baseUrl := os.Getenv("API_URL")

	fmt.Println(baseUrl)

	for offset := 0; offset < 18000; offset = offset + 20 {
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
			id := anime.Data[i].ID
			synopsis := anime.Data[i].Attributes.Synopsis
			discription := anime.Data[i].Attributes.Description
			titleEn := anime.Data[i].Attributes.Titles.En
			titleEnJp := anime.Data[i].Attributes.Titles.EnJp
			avaregeRating := anime.Data[i].Attributes.AverageRating
			popularityRank := anime.Data[i].Attributes.PopularityRank
			ratingRank := anime.Data[i].Attributes.RatingRank
			ageRating := anime.Data[i].Attributes.AgeRating
			ageRatingGuide := anime.Data[i].Attributes.AgeRatingGuide
			status := anime.Data[i].Attributes.Status
			posterImage := anime.Data[i].Attributes.PosterImage.Original
			episodeCount := anime.Data[i].Attributes.EpisodeCount
			subType := anime.Data[i].Attributes.Subtype

			animeData := &Anime{
				ID:             id,
				Synopsis:       synopsis,
				Description:    discription,
				TitleEn:        titleEn,
				TitleEnJp:      titleEnJp,
				AverageRating:  avaregeRating,
				PopularityRank: popularityRank,
				RatingRank:     ratingRank,
				AgeRating:      ageRating,
				AgeRatingGuide: ageRatingGuide,
				Status:         status,
				PosterImage:    posterImage,
				EpisodeCount:   episodeCount,
				Subtype:        subType,
			}

			db.Table("animes").Create(&animeData)

		}

		time.Sleep(time.Millisecond * 200)
	}

}

type Anime struct {
	ID             string `gorm:"primary_key" column:"id"`
	Synopsis       string `column:"synopsis"`
	Description    string `column:"description"`
	TitleEn        string `column:"title_en"`
	TitleEnJp      string `column:"title_en_jp"`
	AverageRating  string `column:"average_rating"`
	PopularityRank int    `column:"popularity_rank"`
	RatingRank     int    `column:"rating_rank"`
	AgeRating      string `column:"age_rating"`
	AgeRatingGuide string `column:"age_rating_guide"`
	Status         string `column:"status"`
	PosterImage    string `column:"poster_image"`
	EpisodeCount   int    `column:"episode_count"`
	Subtype        string `column:"subtype"`
}
