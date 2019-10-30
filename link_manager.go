package main

import (
	"database/sql"
)

const (

	INSERT_LINK = "INSERT INTO links (short_url, target_url) VALUES (?, ?)"
	DELETE_LINK = "DELETE FROM links WHERE short_url = ?"
	QUERY_LINKS = "SELECT * FROM links"
)

type Link struct {
	ShortUrl string
	TargetUrl string
}


func AddLink(db *sql.DB, shortUrl, targetUrl string) string {

	_, err := db.Exec(INSERT_LINK, shortUrl, targetUrl)
	if err != nil {
		return "Error while adding link: " + err.Error()
	}

	return "link added"
}

func DeleteLink(db *sql.DB, id string) string {

	_, err := db.Exec(DELETE_LINK, id)
	if err != nil {
		return "Error while deleting link: " + err.Error()
	}

	return "link deleted"
}

func GetLinks(db *sql.DB) ([]*Link, error) {

	var links []*Link

	rows, err := db.Query(QUERY_LINKS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		link := &Link{}
		err = rows.Scan(&(link.ShortUrl), &(link.TargetUrl))
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func ShowLinks(db *sql.DB) string {

	list := ""

	links, err := GetLinks(db)
	if err != nil {
		return "Error while showing links: " + err.Error()
	}

	for _, link := range links {

		list = list + link.ShortUrl + "\t" + link.TargetUrl + "\n"
	}

	return list
}