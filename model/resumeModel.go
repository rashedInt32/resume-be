// Package model provides data structure
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type body struct {
	Background string `json:"background" bson:"background"`
}

type display struct {
	FontFamily string `json:"fontFamily" bson:"fontFamily"`
}

type name struct {
	Fontsize       int    `json:"fontSize" bson:"fontSize"`
	FontFamily     string `json:"fontFamily" bson:"fontFamily"`
	TextMutedAligh string `json:"textAlign" bson:"textAlign"`
	Display        string `json:"display" bson:"display"`
}

type role struct {
	Fontsize      int    `json:"fontSize" bson:"fontSize"`
	TextTrasnform string `json:"textTransform" bson:"textTransform"`
	TextAligh     string `json:"textAlign" bson:"textAlign"`
	Display       string `json:"display" bson:"display"`
	LetterSpacing int    `json:"letterSpacing" bson:"letterSpacing"`
}

type photo struct {
	Width        int    `json:"width" bson:"width"`
	Height       int    `json:"height" bson:"height"`
	BorderRadius string `json:"borderRadius" bson:"borderRadius"`
	Margin       string `json:"margin" bson:"margin"`
}

type title struct {
	Fontsize      int    `json:"fontSize" bson:"fontSize"`
	TextTrasnform string `json:"textTransform" bson:"textTransform"`
	LetterSpacing int    `json:"letterSpacing" bson:"letterSpacing"`
	MarginBottom  int    `json:"marginBottom" bson:"marginBottom"`
}

type text struct {
	LineHeight string `json:"lineHeight" bson:"lineHeight"`
	Color      string `json:"color" bson:"color"`
	Display    string `json:"display" bson:"display"`
}

type textMuted struct {
	Fontsize int    `json:"fontSize" bson:"fontSize"`
	Color    string `json:"color" bson:"color"`
}

type subTitle struct {
	Fontsize      int    `json:"fontSize" bson:"fontSize"`
	Color         string `json:"color" bson:"color"`
	Display       string `json:"display" bson:"display"`
	TextTrasnform string `json:"textTransform" bson:"textTransform"`
}

type link struct {
	Color      string `json:"color" bson:"color"`
	LineHeight int    `json:"lineHeight" bson:"lineHeight"`
}

type style struct {
	Body      body      `json:"body" bosn:"body"`
	Display   display   `json:"display" bosn:"display"`
	Name      name      `json:"name" bosn:"name"`
	Role      role      `json:"role" bosn:"role"`
	Photo     photo     `json:"photo" bosn:"photo"`
	Title     title     `json:"title" bosn:"title"`
	Text      text      `json:"text" bosn:"text"`
	TextMuted textMuted `json:"textMuted" bosn:"textMuted"`
	SubTitle  subTitle  `json:"subTitle" bosn:"subTitle"`
	Link      link      `json:"link" bosn:"link"`
}

type educationStruct struct {
	University string `json:"university" bson:"university"`
	Year       string `json:"year" bson:"year"`
	Degree     string `json:"degree" bson:"degree"`
}

type expStruct struct {
	Role    string `json:"role" bson:"role"`
	Company string `json:"company" bson:"company"`
	Year    string `json:"year" bson:"year"`
	Desc    string `json:"desc" bson:"desc"`
}

type socialStruct struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}

type projectStruct struct {
	URL string `json:"url" bson:"url"`
}

type skillStruct struct {
	Name  string `json:"name" bson:"name"`
	Range int    `json:"range" bson:"range"`
}

type contacts struct {
	Phone   string `json:"phone" bson:"phone"`
	Email   string `json:"email" bson:"email"`
	Address string `json:"address" bson:"address"`
}

// Resume model data to save resume by user
type Resume struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key        string             `json:"key" bson:"key"`
	Style      style              `json:"style" bson:"style"`
	Name       string             `json:"name" bson:"name"`
	Role       string             `json:"role" bson:"role"`
	Photo      string             `json:"photo" bson:"photo"`
	Overview   string             `json:"overview" bson:"overview"`
	Education  []educationStruct  `json:"education" bson:"education"`
	Experience []expStruct        `json:"experience" bson:"experience"`
	Socials    []socialStruct     `json:"socials" bson:"socials"`
	Projects   []projectStruct    `json:"projects" bson:"projects"`
	Skills     []skillStruct      `json:"skills" bson:"skills"`
	Contact    contacts           `json:"contacts" bson:"contacts"`
}
